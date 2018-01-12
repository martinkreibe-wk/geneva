package elements

// Serializer defines the interface for converting the entity into a serialized edn value.
type Serializer interface {

	// Serialize the element into a string or return the appropriate error.
	Serialize() (composition string, err error)
}
