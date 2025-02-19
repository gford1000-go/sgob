package sgob

import (
	"fmt"
	"testing"
	"time"
)

func TestGetType(t *testing.T) {

	type testData struct {
		V        any
		TypeName string
	}

	var i int = 42
	var i64 int64 = 42
	var s string = "Hello World"
	var td testData
	var tdp *testData = &testData{}
	var tm = time.Now()

	tests := []testData{
		{
			i,
			"int",
		},
		{
			&i,
			"*int",
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
			s,
			"string",
		},
		{
			&s,
			"*string",
		},
		{
			tm,
			"time.Time",
		},
		{
			&tm,
			"*time.Time",
		},
		{
			td,
			"sgob.testData",
		},
		{
			&td,
			"*sgob.testData",
		},
		{
			tdp,
			"*sgob.testData",
		},
		{
			&tdp,
			"**sgob.testData",
		},
	}

	for _, test := range tests {
		name := getTypeName(test.V)
		if name != test.TypeName {
			t.Fatalf("Unexpected mismatch in type name: expected: %s, got: %s", test.TypeName, name)
		}
	}

	testRegistry := NewTypeRegistry()

	f := func(o *TypeRegistryOptions) {
		o.Registry = testRegistry
	}

	for _, test := range tests {
		RegisterType(test.V, f)
	}

	for _, test := range tests {
		ty, err := GetRegisteredType(test.TypeName, f)
		if err != nil {
			t.Fatalf("Unexpected error: %v for type: %s", err, test.TypeName)
		}
		if fmt.Sprintf("%v", ty) != test.TypeName {
			t.Fatalf("Mismatch in type: expected: %s, got: %s", test.TypeName, fmt.Sprintf("%v", ty))
		}
	}

	for _, test := range tests {
		v, err := CreateInstance(test.TypeName, f)
		if err != nil {
			t.Fatalf("Unexpected error: %v for type: %s", err, test.TypeName)
		}
		if getTypeName(v) != test.TypeName {
			t.Fatalf("Mismatch in type: expected: %s, got: %s", test.TypeName, fmt.Sprintf("%T", v))
		}
	}
}
