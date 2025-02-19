package sgob

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/gford1000-go/serialise"
)

// GOBVersion describes a version of a GOB serialisation implementation
// All breaking changes to serialisation will trigger an increment, to ensure
// backwards compatibility to any consumers of existing versions.
type GOBVersion int8

const (
	UnknownVersion GOBVersion = iota
	V1
	OutOfRange
)

var defaultVersion GOBVersion = V1

// NewGOBApproach creates an instance of the
// current default version of the GOB serialisation
func NewGOBApproach(opts ...func(opt *TypeRegistryOptions)) serialise.Approach {
	return NewGOBApproachWithVersion(defaultVersion, opts...)
}

// NewGOBApproachWithVersion creates an Approach instance
// of the specified version, that uses gob serialisation.
func NewGOBApproachWithVersion(version GOBVersion, opts ...func(opt *TypeRegistryOptions)) serialise.Approach {

	var o TypeRegistryOptions
	for _, opt := range opts {
		opt(&o)
	}

	switch version {
	case V1:
		md := serialise.NewMinDataApproachWithVersion(serialise.V1)
		return &gobApproachV1{
			name: fmt.Sprintf("GOB%d", V1),
			o:    &o,
			md:   md,
		}
	default:
		panic(fmt.Sprintf("Illegal GOBVersion passed to NewGOBApproachWithVersion (%d)", version))
	}
}

type gobApproachV1 struct {
	name string
	o    *TypeRegistryOptions
	md   serialise.Approach
}

// Name of the approach
func (g *gobApproachV1) Name() string {
	return g.name
}

// IsSerialisable returns true if an instance of the specified type
// can be serialised
func (g *gobApproachV1) IsSerialisable(v any) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()

	_, err := g.Pack(v)
	return err == nil
}

// Pack serialises the instance to a byte slice
func (g *gobApproachV1) Pack(data any) ([]byte, error) {
	gd, err := g.toGobDataBytes(data)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(gd); err != nil {
		return nil, err
	} else {
		return buf.Bytes(), nil
	}
}

// Unpack deserialises an instance from the byte slice
func (g *gobApproachV1) Unpack(data []byte) (any, error) {
	var buf = bytes.NewBuffer(data)

	decoder := gob.NewDecoder(buf)

	var gd gobData

	err := decoder.Decode(&gd)
	if err != nil {
		return nil, err
	}

	v, err := g.fromGobDataBytes(&gd)
	if err != nil {
		return nil, err
	}
	return v, err

}

type gobTypeID int8

const (
	unknownType gobTypeID = iota
	minDataV1Type
	timeType
	ptimeType
	gobType
	nilType
)

type gobData struct {
	DataType gobTypeID
	TypeName string
	Data     []byte
}

// toGobDataBytes serialises data types to []byte using gob encoding
func (g *gobApproachV1) toGobDataBytes(data any) (*gobData, error) {
	if data == nil {
		return &gobData{DataType: nilType, Data: []byte{}}, nil
	}

	var buf bytes.Buffer

	switch v := data.(type) {
	case []byte, int8, *int8, []int8, int16, *int16, []int16, int32, *int32, []int32, int64, *int64, []int64,
		uint8, *uint8, uint16, *uint16, []uint16, uint32, *uint32, []uint32, uint64, *uint64, []uint64,
		float32, *float32, []float32, float64, *float64, []float64, bool, *bool, []bool, string, *string,
		time.Duration, *time.Duration, []time.Duration, []string, [][]byte:
		b, err := g.md.Pack(data)
		if err != nil {
			return nil, err
		}
		return &gobData{DataType: minDataV1Type, Data: b}, nil
	case time.Time:
		b, err := v.GobEncode()
		return &gobData{DataType: timeType, Data: b}, err
	case *time.Time:
		b, err := v.GobEncode()
		return &gobData{DataType: ptimeType, Data: b}, err
	default:
		encoder := gob.NewEncoder(&buf)
		err := encoder.Encode(data)
		return &gobData{DataType: gobType, TypeName: fmt.Sprintf("%T", data), Data: buf.Bytes()}, err
	}
}

// ErrNoGobData raised when GOB serialisation approach has no data to deserialise
var ErrNoGobData = errors.New("no data provided to deserialise")

// ErrNoDeserialisableData raised when GOB serialisation approach has value data to deserialise
var ErrNoDeserialisableData = errors.New("no data found to deserialise")

func (g *gobApproachV1) fromGobDataBytes(data *gobData) (any, error) {

	if data == nil {
		return nil, ErrNoGobData
	}

	if len(data.Data) == 0 {
		switch data.DataType {
		case nilType:
			return nil, nil

		default:
			return nil, ErrNoDeserialisableData
		}
	}

	switch data.DataType {
	case minDataV1Type:
		return g.md.Unpack(data.Data)
	case timeType:
		var buf = bytes.NewBuffer(data.Data)

		decoder := gob.NewDecoder(buf)

		var v time.Time
		err := decoder.Decode(&v)
		if err != nil {
			return nil, err
		}

		return v, nil
	case ptimeType:
		var buf = bytes.NewBuffer(data.Data)

		decoder := gob.NewDecoder(buf)

		var v = new(time.Time)
		err := decoder.Decode(v)
		if err != nil {
			return nil, err
		}

		return v, nil
	case gobType:

		var buf = bytes.NewBuffer(data.Data)

		decoder := gob.NewDecoder(buf)

		v, err := CreateInstancePtr(data.TypeName, func(o *TypeRegistryOptions) { o.replace(g.o) })
		if err != nil {
			return nil, err
		}

		err = decoder.Decode(v)
		if err != nil {
			return nil, err
		}
		return reflect.ValueOf(v).Elem().Interface(), nil

	default:
		panic(fmt.Sprintf("Ouch! [%s]", data.TypeName))
	}
}
