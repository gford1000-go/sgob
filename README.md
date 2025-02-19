[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://en.wikipedia.org/wiki/MIT_License)
[![Documentation](https://img.shields.io/badge/Documentation-GoDoc-green.svg)](https://godoc.org/github.com/gford1000-go/sgob)

# SGOB | GOB based serialisation

Extends MinData serialisation (from `github.com/gford1000-go/serialise`) to support pre-registered types using `gob`.

This is an example extension of the `Approach` interface.

Supports optional encryption of the byte slice using `aes-gcm`.

```go
func main() {
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
}
```
