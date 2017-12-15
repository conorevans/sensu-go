package schema

import (
	context "context"
	graphql "github.com/graphql-go/graphql"
	ast "github.com/jamesdphillips/graphql/language/ast"
	util "github.com/sensu/backend/graphql/generator/util"
)

// Code generated by graphql/generator package. DO NOT EDIT.

//
// FooResolver represents a collection of methods whose products represent the
// response values of an object type.
//
//  == Example SDL
//
//    """
//    Dog's are not hooman.
//    """
//    type Dog implements Pet {
//      "name of this fine beast."
//      name:  String!
//
//      "breed of this silly animal; probably shibe."
//      breed: [Breed]
//    }
//
//  == Example generated interface
//
//  // DateResolver ...
//  type DogResolver interface {
//    // Name implements response to request for name field.
//    Name(context.Context, interface{}, graphql.Params) interface{}
//    // Breed implements response to request for breed field.
//    Breed(context.Context, interface{}, graphql.Params) interface{}
//    // IsTypeOf is used to determine if a given value is associated with the Dog type
//    IsTypeOf(context.Context, graphql.IsTypeOfParams) bool
//  }
//
//  == Example implementation ...
//
//  // MyDogResolver implements DateResolver interface
//  type MyDogResolver struct {
//    logger logrus.LogEntry
//    store interface{
//      store.BreedStore
//      store.DogStore
//    }
//  }
//
//  // Name implements response to request for name field.
//  func (r *MyDogResolver) Name(ctx context.Context, r interface{}, p graphql.Params) interface{} {
//    // ... implementation details ...
//    dog := r.(DogGetter)
//    return dog.GetName()
//  }
//
//  // Breed implements response to request for breed field.
//  func (r *MyDogResolver) Name(ctx context.Context, r interface{}, p graphql.Params) interface{} {
//    // ... implementation details ...
//    dog := r.(DogGetter)
//    breed := r.store.GetBreed(dog.GetBreedName())
//    return breed
//  }
type FooResolver interface {
	// one implements response to request for 'One' field.
	One(context.Context, interface{}, graphql.Params) interface{}
	// two implements response to request for 'Two' field.
	Two(context.Context, interface{}, graphql.Params) interface{}
	// three implements response to request for 'Three' field.
	Three(context.Context, interface{}, graphql.Params) interface{}
	// four implements response to request for 'Four' field.
	Four(context.Context, interface{}, graphql.Params) interface{}
	// five implements response to request for 'Five' field.
	Five(context.Context, interface{}, graphql.Params) interface{}
	// six implements response to request for 'Six' field.
	Six(context.Context, interface{}, graphql.Params) interface{}
	// IsTypeOf is used to determine if a given value is associated with the Foo type
	IsTypeOf(context.Context, graphql.IsTypeOfParams) interface{}
}

// Foo is quite the type.
func Foo() graphql.ObjectConfig {
	return graphql.ObjectConfig{
		Description: "Foo is quite the type.",
		Fields: graphql.Fields{
			five: graphql.Field{
				Args:              graphql.FieldConfigArgument{argument: *graphql.ArgumentConfig{Type: util.InputType("List")}},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "five",
				Type:              util.OutputType("Named"),
			},
			four: graphql.Field{
				Args:              graphql.FieldConfigArgument{argument: *graphql.ArgumentConfig{Type: util.InputType("Named")}},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "four",
				Type:              util.OutputType("Named"),
			},
			one: graphql.Field{
				Args:              graphql.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "one is a number.",
				Name:              "one",
				Type:              util.OutputType("Named"),
			},
			six: graphql.Field{
				Args:              graphql.FieldConfigArgument{argument: *graphql.ArgumentConfig{Type: util.InputType("Named")}},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "six",
				Type:              util.OutputType("Named"),
			},
			three: graphql.Field{
				Args: graphql.FieldConfigArgument{
					argument: *graphql.ArgumentConfig{Type: util.InputType("Named")},
					other:    *graphql.ArgumentConfig{Type: util.InputType("Named")},
				},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "three",
				Type:              util.OutputType("Named"),
			},
			two: graphql.Field{
				Args:              graphql.FieldConfigArgument{argument: *graphql.ArgumentConfig{Type: util.InputType("NonNull")}},
				DeprecationReason: "",
				Description:       "I am told that two is also a number",
				Name:              "two",
				Type:              util.OutputType("Named"),
			},
		},
		Interfaces: []*graphql.Interface{{util.Interface(Bar)}},
		IsKindOf: func(_ ast.Value) {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see FooResolver.")
		},
		Name: "Foo",
	}
}

//
// AnnotatedObjectResolver represents a collection of methods whose products represent the
// response values of an object type.
//
//  == Example SDL
//
//    """
//    Dog's are not hooman.
//    """
//    type Dog implements Pet {
//      "name of this fine beast."
//      name:  String!
//
//      "breed of this silly animal; probably shibe."
//      breed: [Breed]
//    }
//
//  == Example generated interface
//
//  // DateResolver ...
//  type DogResolver interface {
//    // Name implements response to request for name field.
//    Name(context.Context, interface{}, graphql.Params) interface{}
//    // Breed implements response to request for breed field.
//    Breed(context.Context, interface{}, graphql.Params) interface{}
//    // IsTypeOf is used to determine if a given value is associated with the Dog type
//    IsTypeOf(context.Context, graphql.IsTypeOfParams) bool
//  }
//
//  == Example implementation ...
//
//  // MyDogResolver implements DateResolver interface
//  type MyDogResolver struct {
//    logger logrus.LogEntry
//    store interface{
//      store.BreedStore
//      store.DogStore
//    }
//  }
//
//  // Name implements response to request for name field.
//  func (r *MyDogResolver) Name(ctx context.Context, r interface{}, p graphql.Params) interface{} {
//    // ... implementation details ...
//    dog := r.(DogGetter)
//    return dog.GetName()
//  }
//
//  // Breed implements response to request for breed field.
//  func (r *MyDogResolver) Name(ctx context.Context, r interface{}, p graphql.Params) interface{} {
//    // ... implementation details ...
//    dog := r.(DogGetter)
//    breed := r.store.GetBreed(dog.GetBreedName())
//    return breed
//  }
type AnnotatedObjectResolver interface {
	// annotatedField implements response to request for 'AnnotatedField' field.
	AnnotatedField(context.Context, interface{}, graphql.Params) interface{}
	// IsTypeOf is used to determine if a given value is associated with the AnnotatedObject type
	IsTypeOf(context.Context, graphql.IsTypeOfParams) interface{}
}

// AnnotatedObject self descriptive
func AnnotatedObject() graphql.ObjectConfig {
	return graphql.ObjectConfig{
		Description: "self descriptive",
		Fields: graphql.Fields{annotatedField: graphql.Field{
			Args:              graphql.FieldConfigArgument{arg: *graphql.ArgumentConfig{Type: util.InputType("Named")}},
			DeprecationReason: "",
			Description:       "self descriptive",
			Name:              "annotatedField",
			Type:              util.OutputType("Named"),
		}},
		Interfaces: []*graphql.Interface{{}},
		IsKindOf: func(_ ast.Value) {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see AnnotatedObjectResolver.")
		},
		Name: "AnnotatedObject",
	}
}

//
// FeedResolver represents a collection of methods whose products represent the
// response values of a union type.
//
//  == Example generated interface
//
//  // FeedResolver ...
//  type FeedResolver interface {
//    // ResolveType ... TODO
//    ResolveType(graphql.ResolveTypeParams) *graphql.Object
//  }
//
//  // Example implementation ...
//
//  // MyFeedResolver implements FeedResolver interface
//  type MyFeedResolver struct {
//    logger    logrus.LogEntry
//  }
//
//  // ResolveType ... TODO
//  func (r *MyFeedResolver) ResolveType(p graphql.ResolveTypeParams) *graphql.Object {
//    // ... implementation details ...
//  }
type FeedResolver interface {
	// ResolveType ...
	ResolveType(graphql.ResolveTypeParams) *graphql.Object
}

// Feed includes all stuff and things.
func Feed() graphql.UnionConfig {
	return graphql.UnionConfig{
		Description: "Feed includes all stuff and things.",
		Name:        "Feed",
		ResolveType: func(_ graphql.ResolveTypeParams) *graphql.Object {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see FeedResolver.")
		},
		Types: []*graphql.Object{
			&graphql.Object{PrivateName: "Story"},
			&graphql.Object{PrivateName: "Article"},
			&graphql.Object{PrivateName: "Advert"}},
	}
}

//
// AnnotatedUnionResolver represents a collection of methods whose products represent the
// response values of a union type.
//
//  == Example generated interface
//
//  // FeedResolver ...
//  type FeedResolver interface {
//    // ResolveType ... TODO
//    ResolveType(graphql.ResolveTypeParams) *graphql.Object
//  }
//
//  // Example implementation ...
//
//  // MyFeedResolver implements FeedResolver interface
//  type MyFeedResolver struct {
//    logger    logrus.LogEntry
//  }
//
//  // ResolveType ... TODO
//  func (r *MyFeedResolver) ResolveType(p graphql.ResolveTypeParams) *graphql.Object {
//    // ... implementation details ...
//  }
type AnnotatedUnionResolver interface {
	// ResolveType ...
	ResolveType(graphql.ResolveTypeParams) *graphql.Object
}

// AnnotatedUnion i dont care
func AnnotatedUnion() graphql.UnionConfig {
	return graphql.UnionConfig{
		Description: "AnnotatedUnion i dont care",
		Name:        "AnnotatedUnion",
		ResolveType: func(_ graphql.ResolveTypeParams) *graphql.Object {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see AnnotatedUnionResolver.")
		},
		Types: []*graphql.Object{
			&graphql.Object{PrivateName: "A"},
			&graphql.Object{PrivateName: "B"}},
	}
}

//
// CustomScalarResolver represents a collection of methods whose products represent the input and
// response values of a scalar type.
//
//  == Example generated interface
//
//  // DateResolver ...
//  type DateResolver interface {
//    // Serialize an internal value to include in a response.
//    Serialize(interface{}) interface{}
//    // ParseValue parses an externally provided value to use as an input.
//    ParseValue(interface{}) interface{}
//    // ParseLiteral parses an externally provided literal value to use as an input.
//    ParseLiteral(ast.Value) interface{}
//  }
//
//  // Example implementation ...
//
//  // MyDateResolver implements DateResolver interface
//  type MyDateResolver struct {
//    defaultTZ *time.Location
//    logger    logrus.LogEntry
//  }
//
//  // Serialize serializes given date into RFC 943 compatible string.
//  func (r *MyDateResolver) Serialize(val interface{}) interface{} {
//    // ... implementation details ...
//  }
//
//  // ParseValue takes given value and coerces it into an instance of Time.
//  func (r *MyDateResolver) ParseValue(val interface{}) interface{} {
//    // ... implementation details ...
//    // eg. if val is an int use time.At(), if string time.Parse(), etc.
//  }
//
//  // ParseValue takes given value and coerces it into an instance of Time.
//  func (r *MyDateResolver) ParseValue(val ast.Value) interface{} {
//    // ... implementation details ...
//    //
//    // eg.
//    //
//    // if string value return value,
//    // if IntValue Atoi and return value,
//    // etc.
//  }
type CustomScalarResolver interface {
	// Serialize an internal value to include in a response.
	Serialize(interface{}) interface{}
	// ParseValue parses an externally provided value to use as an input.
	ParseValue(interface{}) interface{}
	// ParseLiteral parses an externally provided literal value to use as an input.
	ParseLiteral(ast.Value) interface{}
}

// CustomScalar self descriptive
func CustomScalar() *graphql.Scalar {
	return graphql.NewScalar(graphql.ScalarConfig{
		Description: "self descriptive",
		Name:        "CustomScalar",
		ParseLiteral: func(_ ast.Value) {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see CustomScalarResolver.")
		},
		ParseValue: func(_ interface{}) {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see CustomScalarResolver.")
		},
		Serialize: func(_ interface{}) {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see CustomScalarResolver.")
		},
	})
}

//
// AnnotatedScalarResolver represents a collection of methods whose products represent the input and
// response values of a scalar type.
//
//  == Example generated interface
//
//  // DateResolver ...
//  type DateResolver interface {
//    // Serialize an internal value to include in a response.
//    Serialize(interface{}) interface{}
//    // ParseValue parses an externally provided value to use as an input.
//    ParseValue(interface{}) interface{}
//    // ParseLiteral parses an externally provided literal value to use as an input.
//    ParseLiteral(ast.Value) interface{}
//  }
//
//  // Example implementation ...
//
//  // MyDateResolver implements DateResolver interface
//  type MyDateResolver struct {
//    defaultTZ *time.Location
//    logger    logrus.LogEntry
//  }
//
//  // Serialize serializes given date into RFC 943 compatible string.
//  func (r *MyDateResolver) Serialize(val interface{}) interface{} {
//    // ... implementation details ...
//  }
//
//  // ParseValue takes given value and coerces it into an instance of Time.
//  func (r *MyDateResolver) ParseValue(val interface{}) interface{} {
//    // ... implementation details ...
//    // eg. if val is an int use time.At(), if string time.Parse(), etc.
//  }
//
//  // ParseValue takes given value and coerces it into an instance of Time.
//  func (r *MyDateResolver) ParseValue(val ast.Value) interface{} {
//    // ... implementation details ...
//    //
//    // eg.
//    //
//    // if string value return value,
//    // if IntValue Atoi and return value,
//    // etc.
//  }
type AnnotatedScalarResolver interface {
	// Serialize an internal value to include in a response.
	Serialize(interface{}) interface{}
	// ParseValue parses an externally provided value to use as an input.
	ParseValue(interface{}) interface{}
	// ParseLiteral parses an externally provided literal value to use as an input.
	ParseLiteral(ast.Value) interface{}
}

// AnnotatedScalar self descriptive
func AnnotatedScalar() *graphql.Scalar {
	return graphql.NewScalar(graphql.ScalarConfig{
		Description: "self descriptive",
		Name:        "AnnotatedScalar",
		ParseLiteral: func(_ ast.Value) {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see AnnotatedScalarResolver.")
		},
		ParseValue: func(_ interface{}) {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see AnnotatedScalarResolver.")
		},
		Serialize: func(_ interface{}) {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see AnnotatedScalarResolver.")
		},
	})
}

//
// NoFieldsResolver represents a collection of methods whose products represent the
// response values of an object type.
//
//  == Example SDL
//
//    """
//    Dog's are not hooman.
//    """
//    type Dog implements Pet {
//      "name of this fine beast."
//      name:  String!
//
//      "breed of this silly animal; probably shibe."
//      breed: [Breed]
//    }
//
//  == Example generated interface
//
//  // DateResolver ...
//  type DogResolver interface {
//    // Name implements response to request for name field.
//    Name(context.Context, interface{}, graphql.Params) interface{}
//    // Breed implements response to request for breed field.
//    Breed(context.Context, interface{}, graphql.Params) interface{}
//    // IsTypeOf is used to determine if a given value is associated with the Dog type
//    IsTypeOf(context.Context, graphql.IsTypeOfParams) bool
//  }
//
//  == Example implementation ...
//
//  // MyDogResolver implements DateResolver interface
//  type MyDogResolver struct {
//    logger logrus.LogEntry
//    store interface{
//      store.BreedStore
//      store.DogStore
//    }
//  }
//
//  // Name implements response to request for name field.
//  func (r *MyDogResolver) Name(ctx context.Context, r interface{}, p graphql.Params) interface{} {
//    // ... implementation details ...
//    dog := r.(DogGetter)
//    return dog.GetName()
//  }
//
//  // Breed implements response to request for breed field.
//  func (r *MyDogResolver) Name(ctx context.Context, r interface{}, p graphql.Params) interface{} {
//    // ... implementation details ...
//    dog := r.(DogGetter)
//    breed := r.store.GetBreed(dog.GetBreedName())
//    return breed
//  }
type NoFieldsResolver interface {
	// IsTypeOf is used to determine if a given value is associated with the NoFields type
	IsTypeOf(context.Context, graphql.IsTypeOfParams) interface{}
}

// NoFields self descriptive
func NoFields() graphql.ObjectConfig {
	return graphql.ObjectConfig{
		Description: "self descriptive",
		Fields:      graphql.Fields{},
		Interfaces:  []*graphql.Interface{{}},
		IsKindOf: func(_ ast.Value) {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see NoFieldsResolver.")
		},
		Name: "NoFields",
	}
}
