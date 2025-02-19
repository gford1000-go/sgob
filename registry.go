package sgob

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

// TypeRegistry manages a map of types to their type name
type TypeRegistry struct {
	registry map[string]reflect.Type
	lck      sync.RWMutex
}

// AddTypeOf adds the type of the value to the registry
func (r *TypeRegistry) AddTypeOf(v any) {
	r.lck.Lock()
	defer r.lck.Unlock()
	r.registry[getTypeName(v)] = reflect.TypeOf(v)
}

// ErrUnknownTypeName is raised if the requested type is not found in the registry
var ErrUnknownTypeName = errors.New("requested type has not been registered")

// GetType returns the registered type for the specified name
func (r *TypeRegistry) GetType(name string) (reflect.Type, error) {
	r.lck.RLock()
	defer r.lck.RUnlock()

	if t, ok := r.registry[name]; ok {
		return t, nil
	}

	return nil, ErrUnknownTypeName

}

// NewTypeRegistry creates an instance of TypeRegistry
func NewTypeRegistry() *TypeRegistry {
	return &TypeRegistry{
		registry: map[string]reflect.Type{},
	}
}

// TypeRegistryOptions allows overrides to the type registration behaviour
type TypeRegistryOptions struct {
	Registry *TypeRegistry
}

func (t *TypeRegistryOptions) replace(overrides *TypeRegistryOptions) {
	t.Registry = overrides.Registry
}

func getTypeName(v any) string {
	return fmt.Sprintf("%T", v)
}

// defaultTypeRegistry is the common registry of types used by default
var defaultTypeRegistry = NewTypeRegistry()

// ErrCannotRegisterNilType raised if nil passed to RegisterType
var ErrCannotRegisterNilType = errors.New("variable must not be nil in call to RegisterType")

// ErrNoRegistry raised when there is no registry provided to operate on
var ErrNoRegistry = errors.New("no registry provided into which to register type")

// RegisterType allows registry of the type specified by the supplied value
func RegisterType(v any, opts ...func(*TypeRegistryOptions)) error {
	if v == nil {
		return ErrCannotRegisterNilType
	}

	o := TypeRegistryOptions{}
	for _, opt := range opts {
		opt(&o)
	}

	if o.Registry == nil {
		o.Registry = defaultTypeRegistry
	}

	o.Registry.AddTypeOf(v)
	return nil
}

// GetRegisteredType returns an instance of the type specified by the name
func GetRegisteredType(name string, opts ...func(*TypeRegistryOptions)) (reflect.Type, error) {

	o := TypeRegistryOptions{Registry: defaultTypeRegistry}
	for _, opt := range opts {
		opt(&o)
	}

	if o.Registry == nil {
		return nil, ErrNoRegistry
	}

	return o.Registry.GetType(name)
}

// CreateInstance returns an instance of the type specified by the name
func CreateInstance(name string, opts ...func(*TypeRegistryOptions)) (any, error) {

	t, err := GetRegisteredType(name, opts...)
	if err != nil {
		return nil, err
	}
	return reflect.New(t).Elem().Interface(), nil
}

// CreateInstancePtr returns a pointer to an instance of the type specified by the name
func CreateInstancePtr(name string, opts ...func(*TypeRegistryOptions)) (any, error) {

	t, err := GetRegisteredType(name, opts...)
	if err != nil {
		return nil, err
	}
	return reflect.New(t).Interface(), nil
}
