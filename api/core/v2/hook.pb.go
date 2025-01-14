// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/sensu/sensu-go/api/core/v2/hook.proto

package v2

import (
	bytes "bytes"
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// HookConfig is the specification of a hook
type HookConfig struct {
	// Metadata contains the name, namespace, labels and annotations of the hook
	ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3,embedded=metadata" json:"metadata,omitempty"`
	// Command is the command to be executed
	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	// Timeout is the timeout, in seconds, at which the hook has to run
	Timeout uint32 `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout"`
	// Stdin indicates if hook requests have stdin enabled
	Stdin bool `protobuf:"varint,4,opt,name=stdin,proto3" json:"stdin"`
	// RuntimeAssets are a list of assets required to execute hook.
	RuntimeAssets        []string `protobuf:"bytes,5,rep,name=runtime_assets,json=runtimeAssets,proto3" json:"runtime_assets"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookConfig) Reset()         { *m = HookConfig{} }
func (m *HookConfig) String() string { return proto.CompactTextString(m) }
func (*HookConfig) ProtoMessage()    {}
func (*HookConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_de6598a27214377c, []int{0}
}
func (m *HookConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HookConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HookConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HookConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookConfig.Merge(m, src)
}
func (m *HookConfig) XXX_Size() int {
	return m.Size()
}
func (m *HookConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HookConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HookConfig proto.InternalMessageInfo

// A Hook is a hook specification and optionally the results of the hook's
// execution.
type Hook struct {
	// Config is the specification of a hook
	HookConfig `protobuf:"bytes,1,opt,name=config,proto3,embedded=config" json:""`
	// Duration of execution
	Duration float64 `protobuf:"fixed64,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// Executed describes the time in which the hook request was executed
	Executed int64 `protobuf:"varint,3,opt,name=executed,proto3" json:"executed"`
	// Issued describes the time in which the hook request was issued
	Issued int64 `protobuf:"varint,4,opt,name=issued,proto3" json:"issued"`
	// Output from the execution of Command
	Output string `protobuf:"bytes,5,opt,name=output,proto3" json:"output,omitempty"`
	// Status is the exit status code produced by the hook
	Status               int32    `protobuf:"varint,6,opt,name=status,proto3" json:"status"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hook) Reset()         { *m = Hook{} }
func (m *Hook) String() string { return proto.CompactTextString(m) }
func (*Hook) ProtoMessage()    {}
func (*Hook) Descriptor() ([]byte, []int) {
	return fileDescriptor_de6598a27214377c, []int{1}
}
func (m *Hook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Hook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Hook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Hook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hook.Merge(m, src)
}
func (m *Hook) XXX_Size() int {
	return m.Size()
}
func (m *Hook) XXX_DiscardUnknown() {
	xxx_messageInfo_Hook.DiscardUnknown(m)
}

var xxx_messageInfo_Hook proto.InternalMessageInfo

func (m *Hook) GetDuration() float64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Hook) GetExecuted() int64 {
	if m != nil {
		return m.Executed
	}
	return 0
}

func (m *Hook) GetIssued() int64 {
	if m != nil {
		return m.Issued
	}
	return 0
}

func (m *Hook) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *Hook) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type HookList struct {
	// Hooks is the list of hooks for the check hook
	Hooks []string `protobuf:"bytes,1,rep,name=hooks,proto3" json:"hooks"`
	// Type indicates the type or response code for the check hook
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookList) Reset()         { *m = HookList{} }
func (m *HookList) String() string { return proto.CompactTextString(m) }
func (*HookList) ProtoMessage()    {}
func (*HookList) Descriptor() ([]byte, []int) {
	return fileDescriptor_de6598a27214377c, []int{2}
}
func (m *HookList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HookList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HookList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HookList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookList.Merge(m, src)
}
func (m *HookList) XXX_Size() int {
	return m.Size()
}
func (m *HookList) XXX_DiscardUnknown() {
	xxx_messageInfo_HookList.DiscardUnknown(m)
}

var xxx_messageInfo_HookList proto.InternalMessageInfo

func (m *HookList) GetHooks() []string {
	if m != nil {
		return m.Hooks
	}
	return nil
}

func (m *HookList) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*HookConfig)(nil), "sensu.core.v2.HookConfig")
	proto.RegisterType((*Hook)(nil), "sensu.core.v2.Hook")
	proto.RegisterType((*HookList)(nil), "sensu.core.v2.HookList")
}

func init() {
	proto.RegisterFile("github.com/sensu/sensu-go/api/core/v2/hook.proto", fileDescriptor_de6598a27214377c)
}

var fileDescriptor_de6598a27214377c = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xbd, 0x6e, 0xd4, 0x40,
	0x10, 0xbe, 0xbd, 0x1f, 0xc7, 0xb7, 0xc9, 0x51, 0x6c, 0x81, 0xcc, 0x15, 0x5e, 0xeb, 0x24, 0x24,
	0x17, 0x60, 0x93, 0x0b, 0x0d, 0x34, 0x01, 0xd3, 0x50, 0x80, 0x90, 0x56, 0xa2, 0xa1, 0x41, 0x3e,
	0x7b, 0xe3, 0x2c, 0x91, 0xbd, 0xd6, 0xed, 0xec, 0x89, 0xbc, 0x01, 0xe2, 0x09, 0x28, 0x53, 0xe6,
	0x11, 0x78, 0x84, 0x94, 0x79, 0x02, 0x0b, 0x4c, 0x67, 0x89, 0x9e, 0x12, 0x79, 0xed, 0x3b, 0x02,
	0x12, 0x52, 0x1a, 0xcf, 0x7c, 0xdf, 0xcc, 0x37, 0xf6, 0x7c, 0x1e, 0xfc, 0x28, 0x13, 0x70, 0xaa,
	0x57, 0x41, 0x22, 0xf3, 0x50, 0xf1, 0x42, 0xe9, 0xee, 0xf9, 0x30, 0x93, 0x61, 0x5c, 0x8a, 0x30,
	0x91, 0x6b, 0x1e, 0x6e, 0x96, 0xe1, 0xa9, 0x94, 0x67, 0x41, 0xb9, 0x96, 0x20, 0xc9, 0xcc, 0x34,
	0x04, 0x6d, 0x25, 0xd8, 0x2c, 0xe7, 0x8f, 0x6f, 0x0c, 0xc8, 0x64, 0x26, 0x43, 0xd3, 0xb5, 0xd2,
	0x27, 0xcf, 0x36, 0x87, 0xc1, 0x51, 0x70, 0x68, 0x48, 0xc3, 0x99, 0xac, 0x1b, 0x32, 0xbf, 0xe5,
	0x6b, 0x73, 0x0e, 0x71, 0xa7, 0x58, 0x7c, 0x1e, 0x62, 0xfc, 0x52, 0xca, 0xb3, 0x17, 0xb2, 0x38,
	0x11, 0x19, 0x79, 0x8b, 0xed, 0xb6, 0x98, 0xc6, 0x10, 0x3b, 0xc8, 0x43, 0xfe, 0xfe, 0xf2, 0x5e,
	0xf0, 0xd7, 0x87, 0x05, 0x6f, 0x56, 0x1f, 0x78, 0x02, 0xaf, 0x39, 0xc4, 0x91, 0x7b, 0x55, 0xd1,
	0xc1, 0x75, 0x45, 0x51, 0x53, 0x51, 0xb2, 0x95, 0x3d, 0x90, 0xb9, 0x00, 0x9e, 0x97, 0x70, 0xce,
	0x76, 0xa3, 0x88, 0x83, 0xf7, 0x12, 0x99, 0xe7, 0x71, 0x91, 0x3a, 0x43, 0x0f, 0xf9, 0x53, 0xb6,
	0x85, 0xe4, 0x3e, 0xde, 0x03, 0x91, 0x73, 0xa9, 0xc1, 0x19, 0x79, 0xc8, 0x9f, 0x45, 0xfb, 0x4d,
	0x45, 0xb7, 0x14, 0xdb, 0x26, 0x84, 0xe2, 0x89, 0x82, 0x54, 0x14, 0xce, 0xd8, 0x43, 0xbe, 0x1d,
	0x4d, 0x9b, 0x8a, 0x76, 0x04, 0xeb, 0x02, 0x79, 0x82, 0xef, 0xac, 0x75, 0xd1, 0xb6, 0xbf, 0x8f,
	0x95, 0xe2, 0xa0, 0x9c, 0x89, 0x37, 0xf2, 0xa7, 0x11, 0x69, 0x2a, 0xfa, 0x4f, 0x85, 0xcd, 0x7a,
	0xfc, 0xdc, 0xc0, 0xa7, 0xf6, 0xa7, 0x0b, 0x3a, 0xb8, 0xbc, 0xa0, 0x68, 0xf1, 0x13, 0xe1, 0x71,
	0x6b, 0x06, 0x39, 0xc6, 0x56, 0x62, 0x0c, 0xf9, 0x8f, 0x09, 0x7f, 0x1c, 0x8b, 0x0e, 0x6e, 0x98,
	0x30, 0x60, 0xbd, 0x8c, 0xcc, 0xb1, 0x9d, 0xea, 0x75, 0x0c, 0x42, 0x16, 0x66, 0x63, 0xc4, 0x76,
	0x98, 0xf8, 0xd8, 0xe6, 0x1f, 0x79, 0xa2, 0x81, 0xa7, 0x66, 0xe7, 0x51, 0x74, 0xd0, 0x54, 0x74,
	0xc7, 0xb1, 0x5d, 0x46, 0x16, 0xd8, 0x12, 0x4a, 0x69, 0x9e, 0x9a, 0xb5, 0x47, 0x11, 0x6e, 0x2a,
	0xda, 0x33, 0xac, 0x8f, 0xe4, 0x2e, 0xb6, 0xa4, 0x86, 0x52, 0x83, 0x33, 0x31, 0xce, 0xf6, 0xa8,
	0xd5, 0x2a, 0x88, 0x41, 0x2b, 0xc7, 0xf2, 0x90, 0x3f, 0xe9, 0xb4, 0x1d, 0xc3, 0xfa, 0xb8, 0x38,
	0xc6, 0x76, 0xbb, 0xc9, 0x2b, 0xa1, 0x8c, 0xc3, 0xed, 0x35, 0x2a, 0x07, 0x19, 0xdf, 0x8c, 0xc3,
	0x86, 0x60, 0x5d, 0x20, 0x04, 0x8f, 0xe1, 0xbc, 0xe4, 0xfd, 0x0f, 0x34, 0x79, 0xe4, 0xfd, 0xfa,
	0xee, 0xa2, 0xcb, 0xda, 0x45, 0x5f, 0x6b, 0x17, 0x5d, 0xd5, 0x2e, 0xba, 0xae, 0x5d, 0xf4, 0xad,
	0x76, 0xd1, 0x97, 0x1f, 0xee, 0xe0, 0xdd, 0x70, 0xb3, 0x5c, 0x59, 0xe6, 0xcc, 0x8e, 0x7e, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x6e, 0xfa, 0xcd, 0x02, 0x11, 0x03, 0x00, 0x00,
}

func (this *HookConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HookConfig)
	if !ok {
		that2, ok := that.(HookConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ObjectMeta.Equal(&that1.ObjectMeta) {
		return false
	}
	if this.Command != that1.Command {
		return false
	}
	if this.Timeout != that1.Timeout {
		return false
	}
	if this.Stdin != that1.Stdin {
		return false
	}
	if len(this.RuntimeAssets) != len(that1.RuntimeAssets) {
		return false
	}
	for i := range this.RuntimeAssets {
		if this.RuntimeAssets[i] != that1.RuntimeAssets[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Hook) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Hook)
	if !ok {
		that2, ok := that.(Hook)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.HookConfig.Equal(&that1.HookConfig) {
		return false
	}
	if this.Duration != that1.Duration {
		return false
	}
	if this.Executed != that1.Executed {
		return false
	}
	if this.Issued != that1.Issued {
		return false
	}
	if this.Output != that1.Output {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HookList) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HookList)
	if !ok {
		that2, ok := that.(HookList)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Hooks) != len(that1.Hooks) {
		return false
	}
	for i := range this.Hooks {
		if this.Hooks[i] != that1.Hooks[i] {
			return false
		}
	}
	if this.Type != that1.Type {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type HookConfigFace interface {
	Proto() github_com_golang_protobuf_proto.Message
	GetObjectMeta() ObjectMeta
	GetCommand() string
	GetTimeout() uint32
	GetStdin() bool
	GetRuntimeAssets() []string
}

func (this *HookConfig) Proto() github_com_golang_protobuf_proto.Message {
	return this
}

func (this *HookConfig) TestProto() github_com_golang_protobuf_proto.Message {
	return NewHookConfigFromFace(this)
}

func (this *HookConfig) GetObjectMeta() ObjectMeta {
	return this.ObjectMeta
}

func (this *HookConfig) GetCommand() string {
	return this.Command
}

func (this *HookConfig) GetTimeout() uint32 {
	return this.Timeout
}

func (this *HookConfig) GetStdin() bool {
	return this.Stdin
}

func (this *HookConfig) GetRuntimeAssets() []string {
	return this.RuntimeAssets
}

func NewHookConfigFromFace(that HookConfigFace) *HookConfig {
	this := &HookConfig{}
	this.ObjectMeta = that.GetObjectMeta()
	this.Command = that.GetCommand()
	this.Timeout = that.GetTimeout()
	this.Stdin = that.GetStdin()
	this.RuntimeAssets = that.GetRuntimeAssets()
	return this
}

func (m *HookConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HookConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HookConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RuntimeAssets) > 0 {
		for iNdEx := len(m.RuntimeAssets) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RuntimeAssets[iNdEx])
			copy(dAtA[i:], m.RuntimeAssets[iNdEx])
			i = encodeVarintHook(dAtA, i, uint64(len(m.RuntimeAssets[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Stdin {
		i--
		if m.Stdin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Timeout != 0 {
		i = encodeVarintHook(dAtA, i, uint64(m.Timeout))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Command) > 0 {
		i -= len(m.Command)
		copy(dAtA[i:], m.Command)
		i = encodeVarintHook(dAtA, i, uint64(len(m.Command)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintHook(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Hook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Hook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Hook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Status != 0 {
		i = encodeVarintHook(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Output) > 0 {
		i -= len(m.Output)
		copy(dAtA[i:], m.Output)
		i = encodeVarintHook(dAtA, i, uint64(len(m.Output)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Issued != 0 {
		i = encodeVarintHook(dAtA, i, uint64(m.Issued))
		i--
		dAtA[i] = 0x20
	}
	if m.Executed != 0 {
		i = encodeVarintHook(dAtA, i, uint64(m.Executed))
		i--
		dAtA[i] = 0x18
	}
	if m.Duration != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Duration))))
		i--
		dAtA[i] = 0x11
	}
	{
		size, err := m.HookConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintHook(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *HookList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HookList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HookList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintHook(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Hooks) > 0 {
		for iNdEx := len(m.Hooks) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Hooks[iNdEx])
			copy(dAtA[i:], m.Hooks[iNdEx])
			i = encodeVarintHook(dAtA, i, uint64(len(m.Hooks[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintHook(dAtA []byte, offset int, v uint64) int {
	offset -= sovHook(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedHookConfig(r randyHook, easy bool) *HookConfig {
	this := &HookConfig{}
	v1 := NewPopulatedObjectMeta(r, easy)
	this.ObjectMeta = *v1
	this.Command = string(randStringHook(r))
	this.Timeout = uint32(r.Uint32())
	this.Stdin = bool(bool(r.Intn(2) == 0))
	v2 := r.Intn(10)
	this.RuntimeAssets = make([]string, v2)
	for i := 0; i < v2; i++ {
		this.RuntimeAssets[i] = string(randStringHook(r))
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedHook(r, 6)
	}
	return this
}

func NewPopulatedHook(r randyHook, easy bool) *Hook {
	this := &Hook{}
	v3 := NewPopulatedHookConfig(r, easy)
	this.HookConfig = *v3
	this.Duration = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Duration *= -1
	}
	this.Executed = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Executed *= -1
	}
	this.Issued = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Issued *= -1
	}
	this.Output = string(randStringHook(r))
	this.Status = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Status *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedHook(r, 7)
	}
	return this
}

func NewPopulatedHookList(r randyHook, easy bool) *HookList {
	this := &HookList{}
	v4 := r.Intn(10)
	this.Hooks = make([]string, v4)
	for i := 0; i < v4; i++ {
		this.Hooks[i] = string(randStringHook(r))
	}
	this.Type = string(randStringHook(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedHook(r, 3)
	}
	return this
}

type randyHook interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneHook(r randyHook) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringHook(r randyHook) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneHook(r)
	}
	return string(tmps)
}
func randUnrecognizedHook(r randyHook, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldHook(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldHook(dAtA []byte, r randyHook, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateHook(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateHook(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateHook(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateHook(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateHook(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateHook(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateHook(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *HookConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovHook(uint64(l))
	l = len(m.Command)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovHook(uint64(m.Timeout))
	}
	if m.Stdin {
		n += 2
	}
	if len(m.RuntimeAssets) > 0 {
		for _, s := range m.RuntimeAssets {
			l = len(s)
			n += 1 + l + sovHook(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Hook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.HookConfig.Size()
	n += 1 + l + sovHook(uint64(l))
	if m.Duration != 0 {
		n += 9
	}
	if m.Executed != 0 {
		n += 1 + sovHook(uint64(m.Executed))
	}
	if m.Issued != 0 {
		n += 1 + sovHook(uint64(m.Issued))
	}
	l = len(m.Output)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovHook(uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HookList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Hooks) > 0 {
		for _, s := range m.Hooks {
			l = len(s)
			n += 1 + l + sovHook(uint64(l))
		}
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovHook(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHook(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHook(x uint64) (n int) {
	return sovHook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HookConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHook
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HookConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HookConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Command", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Command = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stdin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Stdin = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuntimeAssets", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RuntimeAssets = append(m.RuntimeAssets, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHook
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Hook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHook
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Hook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Hook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HookConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HookConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Duration = float64(math.Float64frombits(v))
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Executed", wireType)
			}
			m.Executed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Executed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issued", wireType)
			}
			m.Issued = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Issued |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Output", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Output = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHook
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HookList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHook
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HookList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HookList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hooks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hooks = append(m.Hooks, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHook
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHook
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHook
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHook
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthHook
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHook
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHook
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHook        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHook          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHook = fmt.Errorf("proto: unexpected end of group")
)
