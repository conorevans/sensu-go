package opampd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-go/backend/messaging"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/open-telemetry/opamp-go/protobufs"
	"google.golang.org/protobuf/proto"
)

const (
	// emitMetricsEventInterval is the interval at which an OpAMP metrics event
	// is published.
	emitMetricsEventInterval = 30 * time.Second

	statusReportsReceivedLabel = "ot-status-reports-received-total"
	agentConfigsSentLabel      = "ot-agent-configs-sent-total"
	errorCountLabel            = "ot-error-count-total"
)

var (
	// statusReportsReceived is the total number of StatusReport messages
	// received from OpenTelemetry agents.
	statusReportReceived uint64

	// agentConfigsSent in the total number of AgentRemoteConfig messages sent
	// to OpenTelemetry agents.
	agentConfigsSent uint64

	// errorCount is the total number of errors encountered while exchanging
	// messages with OpenTelemetry agents.
	errorCount uint64
)

//
func emitMetricsEvent(bus messaging.MessageBus, entity *corev2.Entity) error {
	event := &corev2.Event{}
	event.Name = "opampd-metrics"
	event.Namespace = "default"
	event.Entity = entity
	event.Timestamp = time.Now().Unix()

	event.Metrics = &corev2.Metrics{
		Points: []*corev2.MetricPoint{
			{
				Name:      statusReportsReceivedLabel,
				Value:     float64(atomic.LoadUint64(&statusReportReceived)),
				Timestamp: time.Now().Unix(),
			},
			{
				Name:      agentConfigsSentLabel,
				Value:     float64(atomic.LoadUint64(&agentConfigsSent)),
				Timestamp: time.Now().Unix(),
			},
			{
				Name:      errorCountLabel,
				Value:     float64(atomic.LoadUint64(&errorCount)),
				Timestamp: time.Now().Unix(),
			},
		},
	}

	return bus.Publish(messaging.TopicEventRaw, event)
}

type Config struct {
	Host          string
	Port          int
	Path          string
	Handler       MessageHandler
	EventBus      messaging.MessageBus
	BackendEntity *corev2.Entity
}

type OpAMPD struct {
	host          string
	port          int
	path          string
	upgrader      *websocket.Upgrader
	connections   map[string]*websocket.Conn
	httpServer    *http.Server
	wg            *sync.WaitGroup
	errChan       chan error
	handler       MessageHandler
	eventBus      messaging.MessageBus
	backendEntity *corev2.Entity
}

// New creates and bind the OpAMP server to the specified port.
func New(config *Config) (*OpAMPD, error) {
	// Add validation here

	d := &OpAMPD{
		host: config.Host,
		port: config.Port,
		path: config.Path,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(_ *http.Request) bool {
				return true
			},
		},
		connections:   make(map[string]*websocket.Conn),
		wg:            &sync.WaitGroup{},
		errChan:       make(chan error, 1),
		handler:       config.Handler,
		eventBus:      config.EventBus,
		backendEntity: config.BackendEntity,
	}

	router := mux.NewRouter()
	router.HandleFunc(d.path, d.handleWS)

	d.httpServer = &http.Server{
		Addr:         net.JoinHostPort(d.host, strconv.Itoa(d.port)),
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		//TLSConfig:    tlsServerConfig,
		// Capture the log entries from agentd's HTTP server
		//ErrorLog: logger,
		ConnState: func(c net.Conn, cs http.ConnState) {
			if cs != http.StateClosed {
				var msg []byte
				if _, err := c.Read(msg); err != nil {
					logger.WithError(err).Error("websocket connection error")
				}
			}
		},
	}

	return d, nil
}

func (d *OpAMPD) Start() error {
	logger.Warn("starting opampd on address: ", d.httpServer.Addr)
	ln, err := net.Listen("tcp", d.httpServer.Addr)
	if err != nil {
		return fmt.Errorf("failed to start agentd: %s", err)
	}

	d.wg.Add(1)

	go func() {
		defer d.wg.Done()
		err := d.httpServer.Serve(ln)
		if err != nil && err != http.ErrServerClosed {
			d.errChan <- fmt.Errorf("opampd failed while serving: %s", err)
		}
	}()

	go func() {
		ticker := time.NewTicker(emitMetricsEventInterval)
		for range ticker.C {
			if err := emitMetricsEvent(d.eventBus, d.backendEntity); err != nil {
				d.errChan <- err
			}
		}
	}()

	return nil
}

func (d *OpAMPD) Stop() error {
	if err := d.httpServer.Shutdown(context.Background()); err != nil {
		// failure/timeout shutting down the server gracefully
		logger.Error("failed to shutdown http server gracefully - forcing shutdown")
		if closeErr := d.httpServer.Close(); closeErr != nil {
			logger.Error("failed to shutdown http server forcefully")
		}
	}
	d.wg.Wait()

	return nil
}

func (d *OpAMPD) Err() <-chan error {
	return d.errChan
}

func (d *OpAMPD) Name() string {
	return "opampd"
}

// handleWS upgrades an incoming http message into a WebSocket connection.
func (d *OpAMPD) handleWS(response http.ResponseWriter, request *http.Request) {
	remoteAddr := request.RemoteAddr
	connection, err := d.upgrader.Upgrade(response, request, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			logger.Error(err)
		}
		return
	}
	logger.Infof("WebSocket client connected from %s\n", remoteAddr)
	d.connections[remoteAddr] = connection
	go d.messageReader(connection)
}

// messageReader  is a goroutine that reads messages from a websocket
///connection. There is one goroutine per client connection. The received
// messages are published to the inMessages channel to notify the listeners
func (d *OpAMPD) messageReader(connection *websocket.Conn) {
	for {
		_, message, err := connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Errorf("error: %v", err.Error())
			}
			break
		}
		go d.handleMessage(connection, message)
	}
}

// handleMessage parses a protobuf message received from the agent and calls the
// appropriate handler.
func (d *OpAMPD) handleMessage(connection *websocket.Conn, message []byte) {
	a2s := protobufs.AgentToServer{}
	err := proto.Unmarshal(message, &a2s)
	if err != nil {
		logger.Errorf("error parsing protobuf message: %v", err)
		return
	}

	logger.Infof("OpAMP message received from %s\n", a2s.InstanceUid)

	var s2a *protobufs.ServerToAgent
	if a2s.StatusReport != nil {
		logger.Infof("received status report from %s", a2s.InstanceUid)
		atomic.AddUint64(&statusReportReceived, 1)
		s2a, err = d.handler.OnStatusReport(a2s.InstanceUid, a2s.StatusReport)
	} else if a2s.AddonStatuses != nil {
		logger.Infof("received addon statuses from %s", a2s.InstanceUid)
		s2a, err = d.handler.OnAddonStatuses(a2s.InstanceUid, a2s.AddonStatuses)
	} else if a2s.AgentInstallStatus != nil {
		logger.Infof("received agent install status from %s", a2s.InstanceUid)
		s2a, err = d.handler.OnAgentInstallStatus(a2s.InstanceUid, a2s.AgentInstallStatus)
	} else if a2s.AgentDisconnect != nil {
		logger.Infof("received agent disconnect %s", a2s.InstanceUid)
		s2a, err = d.handler.OnAgentDisconnect(s2a.InstanceUid, a2s.AgentDisconnect)
	} else {
		// invalid message
		logger.Errorf("invalid message from %s", a2s.InstanceUid)
		atomic.AddUint64(&errorCount, 1)
	}

	if err != nil {
		logger.Errorf("error processing message from agent %s: %v", a2s.InstanceUid, err)
		atomic.AddUint64(&errorCount, 1)
		return
	}

	binary, err := proto.Marshal(s2a)
	if err != nil {
		logger.Errorf("error marshaling ServerToAgent message for agent %s: %v", a2s.InstanceUid, err)
		atomic.AddUint64(&errorCount, 1)
	}

	err = connection.WriteMessage(websocket.BinaryMessage, binary)
	if err != nil {
		logger.Errorf("error writing response back to agent %s: %v", a2s.InstanceUid, err)
		atomic.AddUint64(&errorCount, 1)
	}
}
