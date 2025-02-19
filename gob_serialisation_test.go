package sgob

import (
	"fmt"
	"testing"
	"time"

	"github.com/gford1000-go/serialise"
)

func testCompareValue[T comparable](a, b any, name string, t *testing.T) {
	switch v := b.(type) {
	case T:
		if aa, ok := a.(T); ok {
			if aa != v {
				t.Fatalf("Data mismatch: expected %v, got: %v", v, aa)
			}
		} else {
			t.Fatalf("Type mismatch: expected: %s, got: %s", name, fmt.Sprintf("%T", a))
		}
	default:
		t.Fatalf("Unexpected error: b was the wrong type: %s", fmt.Sprintf("%T", b))
	}
}

func testCompareSliceValue[T comparable](a, b any, name string, t *testing.T) {
	switch v := b.(type) {
	case []T:
		if aa, ok := a.([]T); ok {
			if len(aa) != len(v) {
				t.Fatalf("Data size mismatch: expected %v, got: %v", len(v), len(aa))
			}
			for i, vv := range aa {
				if v[i] != vv {
					t.Fatalf("Data mismatch at %d: expected %v, got: %v", i, vv, aa)
				}
			}
		} else {
			t.Fatalf("Type mismatch: expected: %s, got: %s", name, fmt.Sprintf("%T", a))
		}
	default:
		t.Fatalf("Unexpected error: b was the wrong type: %s", fmt.Sprintf("%T", b))
	}
}

func testComparePtrValue[T comparable](a, b any, name string, t *testing.T) {
	switch v := b.(type) {
	case *T:
		if aa, ok := a.(*T); ok {
			if (aa == nil && v != nil) || (aa != nil && v == nil) {
				t.Fatalf("Pointer mismatch: expected %v, got: %v", v, aa)
			}
			if *aa != *v {
				t.Fatalf("Data mismatch: expected %v, got: %v", *v, *aa)
			}
		} else {
			t.Fatalf("Type mismatch: expected: %s, got: %s", name, fmt.Sprintf("%T", a))
		}
	default:
		t.Fatalf("Unexpected error: b was the wrong type: %s", fmt.Sprintf("%T", b))
	}
}

func TestToBytes(t *testing.T) {

	type testData struct {
		V        any
		TypeName string
	}

	var i8 int8 = 42
	var i16 int16 = 42
	var i32 int32 = 42
	var i64 int64 = 42
	var u8 uint8 = 42
	var u16 uint16 = 42
	var u32 uint32 = 42
	var u64 uint64 = 42
	var f32 float32 = 42.42
	var f64 float64 = -42.42
	var bl bool = true
	var td time.Duration = 1234
	var bs []byte = []byte("Hello World")
	var ss []string = []string{"Hello", "World"}
	var is8 []int8 = []int8{1, 2, 3, 4}
	var is16 []int16 = []int16{1, 2, 3, 4}
	var is32 []int32 = []int32{1, 2, 3, 4}
	var is64 []int64 = []int64{1, 2, 3, 4}
	var uis16 []uint16 = []uint16{1, 2, 3, 4}
	var uis32 []uint32 = []uint32{1, 2, 3, 4}
	var uis64 []uint64 = []uint64{1, 2, 3, 4}
	var fs32 []float32 = []float32{1, 2, 3, 4}
	var fs64 []float64 = []float64{1, 2, 3, 4}
	var bbs []bool = []bool{false, true, true, false}
	var tds []time.Duration = []time.Duration{1, 2, 3, 4}

	compareValue := func(a, b any, name string) {
		if b == nil {
			if a != nil {
				t.Fatalf("Mismatch in <nil>")
			}
			return
		}

		switch b.(type) {
		case []byte:
			testCompareSliceValue[byte](a, b, name, t)
		case int8:
			testCompareValue[int8](a, b, name, t)
		case *int8:
			testComparePtrValue[int8](a, b, name, t)
		case []int8:
			testCompareSliceValue[int8](a, b, name, t)
		case int16:
			testCompareValue[int16](a, b, name, t)
		case *int16:
			testComparePtrValue[int16](a, b, name, t)
		case []int16:
			testCompareSliceValue[int16](a, b, name, t)
		case int32:
			testCompareValue[int32](a, b, name, t)
		case *int32:
			testComparePtrValue[int32](a, b, name, t)
		case []int32:
			testCompareSliceValue[int32](a, b, name, t)
		case int64:
			testCompareValue[int64](a, b, name, t)
		case *int64:
			testComparePtrValue[int64](a, b, name, t)
		case []int64:
			testCompareSliceValue[int64](a, b, name, t)
		case uint8:
			testCompareValue[uint8](a, b, name, t)
		case *uint8:
			testComparePtrValue[uint8](a, b, name, t)
		case uint16:
			testCompareValue[uint16](a, b, name, t)
		case *uint16:
			testComparePtrValue[uint16](a, b, name, t)
		case []uint16:
			testCompareSliceValue[uint16](a, b, name, t)
		case uint32:
			testCompareValue[uint32](a, b, name, t)
		case *uint32:
			testComparePtrValue[uint32](a, b, name, t)
		case []uint32:
			testCompareSliceValue[uint32](a, b, name, t)
		case uint64:
			testCompareValue[uint64](a, b, name, t)
		case *uint64:
			testComparePtrValue[uint64](a, b, name, t)
		case []uint64:
			testCompareSliceValue[uint64](a, b, name, t)
		case float32:
			testCompareValue[float32](a, b, name, t)
		case *float32:
			testComparePtrValue[float32](a, b, name, t)
		case []float32:
			testCompareSliceValue[float32](a, b, name, t)
		case float64:
			testCompareValue[float64](a, b, name, t)
		case *float64:
			testComparePtrValue[float64](a, b, name, t)
		case []float64:
			testCompareSliceValue[float64](a, b, name, t)
		case bool:
			testCompareValue[bool](a, b, name, t)
		case *bool:
			testComparePtrValue[bool](a, b, name, t)
		case []bool:
			testCompareSliceValue[bool](a, b, name, t)
		case time.Duration:
			testCompareValue[time.Duration](a, b, name, t)
		case *time.Duration:
			testComparePtrValue[time.Duration](a, b, name, t)
		case []time.Duration:
			testCompareSliceValue[time.Duration](a, b, name, t)
		case []string:
			testCompareSliceValue[string](a, b, name, t)
		default:
			t.Fatalf("No test available for type: %s (%s)", fmt.Sprintf("%T", b), name)
		}

	}

	tests := []testData{
		{
			nil,
			"blah",
		},
		{
			i8,
			"int8",
		},
		{
			&i8,
			"*int8",
		},
		{
			i16,
			"int16",
		},
		{
			&i16,
			"*int16",
		},
		{
			i32,
			"int32",
		},
		{
			&i32,
			"*int32",
		},
		{
			i64,
			"int64",
		},
		{
			&i64,
			"*int64",
		},
		{
			u8,
			"uint8",
		},
		{
			&u8,
			"*uint8",
		},
		{
			u16,
			"uint16",
		},
		{
			&u16,
			"*uint16",
		},
		{
			u32,
			"uint32",
		},
		{
			&u32,
			"*uint32",
		},
		{
			u64,
			"uint64",
		},
		{
			&u64,
			"*uint64",
		},
		{
			f32,
			"float32",
		},
		{
			&f32,
			"*float32",
		},
		{
			f64,
			"float64",
		},
		{
			&f64,
			"*float64",
		},
		{
			bl,
			"bool",
		},
		{
			&bl,
			"*bool",
		},
		{
			td,
			"time.Duration",
		},
		{
			&td,
			"*time.Duration",
		},
		{
			bs,
			"[]byte",
		},
		{
			is8,
			"[]int8",
		},
		{
			is16,
			"[]int16",
		},
		{
			is32,
			"[]int32",
		},
		{
			is64,
			"[]int64",
		},
		{
			uis16,
			"[]uint16",
		},
		{
			uis32,
			"[]uint32",
		},
		{
			uis64,
			"[]uint64",
		},
		{
			fs32,
			"[]float32",
		},
		{
			fs64,
			"[]float64",
		},
		{
			bbs,
			"[]bool",
		},
		{
			tds,
			"[]time.Duration",
		},
		{
			ss,
			"[]string",
		},
	}

	testRegistry := NewTypeRegistry()
	testRegistry.AddTypeOf(ss)
	testRegistry.AddTypeOf(is8)

	f := func(o *TypeRegistryOptions) {
		o.Registry = testRegistry
	}

	approach := NewGOBApproach(f)

	for _, test := range tests {

		b, _, err := serialise.ToBytes(test.V, serialise.WithSerialisationApproach(approach))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		v, err := serialise.FromBytes(b, approach)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		compareValue(v, test.V, test.TypeName)
	}
}

func TestToBytes_1(t *testing.T) {

	type testData struct {
		V        any
		TypeName string
	}

	var i8 int8 = 42
	var ss []string = []string{"This", "is", "not", "encrypted"}
	var ss2 []string = []string{"01234567890123456789012345678901234567890123456789"}

	compareValue := func(a, b any, name string) {
		if b == nil {
			if a != nil {
				t.Fatalf("Mismatch in <nil>")
			}
			return
		}

		switch b.(type) {
		case int8:
			testCompareValue[int8](a, b, name, t)
		case []string:
			testCompareSliceValue[string](a, b, name, t)
		default:
			t.Fatalf("No test available for type: %s (%s)", fmt.Sprintf("%T", b), name)
		}

	}

	tests := []testData{
		{
			ss,
			"[]string",
		},
		{
			ss2,
			"[]string",
		},
		{
			i8,
			"int8",
		},
	}

	testRegistry := NewTypeRegistry()
	testRegistry.AddTypeOf(ss)

	f := func(o *TypeRegistryOptions) {
		o.Registry = testRegistry
	}

	approach := NewGOBApproach(f)

	key := []byte("01234567890123456789012345678912")

	for _, test := range tests {

		b, _, err := serialise.ToBytes(test.V, serialise.WithSerialisationApproach(approach), serialise.WithAESGCMEncryption(key))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		v, err := serialise.FromBytes(b, approach, serialise.WithAESGCMEncryption(key))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		compareValue(v, test.V, test.TypeName)
	}
}

func TestToBytes_2(t *testing.T) {

	type testData struct {
		V        int8
		TypeName string
	}

	compareValue := func(a, b any, name string) {
		if b == nil {
			if a != nil {
				t.Fatalf("Mismatch in <nil>")
			}
			return
		}

		switch b.(type) {
		case testData:
			testCompareValue[testData](a, b, name, t)
		case *testData:
			testComparePtrValue[testData](a, b, name, t)
		default:
			t.Fatalf("No test available for type: %s (%s)", fmt.Sprintf("%T", b), name)
		}

	}

	tests := []any{
		testData{
			42,
			"int8",
		},
		&testData{
			42,
			"int8",
		},
	}

	testRegistry := NewTypeRegistry()
	testRegistry.AddTypeOf(tests[0])
	testRegistry.AddTypeOf(tests[1])

	f := func(o *TypeRegistryOptions) {
		o.Registry = testRegistry
	}

	approach := NewGOBApproach(f)

	for _, test := range tests {
		b, _, err := serialise.ToBytes(test, serialise.WithSerialisationApproach(approach))
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		v, err := serialise.FromBytes(b, approach)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		compareValue(v, test, fmt.Sprintf("%T", test))
	}

}
