package elements

import "strconv"

// initBoolean will add the element factory to the collection of factories
func initBoolean() error {
	return AddElementTypeFactory(BooleanType, func(input interface{}) (elem Element, err error) {
		if v, ok := input.(bool); ok {
			elem, err = NewBooleanElement(v)
		} else {
			err = ErrInvalidInput
		}
		return elem, err
	})
}

// NewBooleanElement creates a new boolean element or an error.
func NewBooleanElement(value bool) (elem Element, err error) {

	elem, err = makeBaseElement(value, BooleanType, func(value interface{}) (out string, e error) {
		out = strconv.FormatBool(value.(bool))
		return out, e
	})

	return elem, err
}
