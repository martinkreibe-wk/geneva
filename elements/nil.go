package elements

const (
	// NilLiteral defines the nil string representation.
	NilLiteral = "nil"
)

// initNil will add the element factory to the collection of factories
func initNil() error {
	return AddElementTypeFactory(NilType, func(input interface{}) (elem Element, err error) {
		if input == nil {
			elem, err = NewNilElement()
		} else {
			err = ErrInvalidInput
		}
		return elem, err
	})
}

// NewNilElement returns the nil element or an error.
func NewNilElement() (elem Element, err error) {
	elem, err = makeBaseElement(nil, NilType, func(value interface{}) (string, error) {
		return NilLiteral, nil
	})
	return elem, err
}
