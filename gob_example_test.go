package sgob

import (
	"fmt"
	"strings"

	"github.com/gford1000-go/serialise"
)

func Example() {
	type myType struct {
		Data []string
	}

	data := &myType{Data: []string{"Hello", "World!"}}

	// Create the GOB serialisation Approach, including its own TypeRegistry
	testRegistry := NewTypeRegistry()
	testRegistry.AddTypeOf(data)
	approach := NewGOBApproach(func(o *TypeRegistryOptions) { o.Registry = testRegistry })

	// Register the Approach so that it can be retrieved later
	serialise.RegisterApproach(approach)

	// Serialise the data using the default version of MinData serialisation
	b, name, _ := serialise.ToBytes(data, serialise.WithSerialisationApproach(approach))

	// Retrieve the Approach used for serialisation, from the returned name
	approach, _ = serialise.GetApproach(name)

	// Deserialise
	v, _ := serialise.FromBytes(b, approach)

	fmt.Println(strings.Join(data.Data, " ") == strings.Join(v.(*myType).Data, " "))
	// Output: true

}
