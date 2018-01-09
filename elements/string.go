package elements

import (
	"strconv"
)

// init will add the element factory to the collection of factories
func initString() error {
	return AddElementTypeFactory(StringType, func(input interface{}) (elem Element, err error) {
		if v, ok := input.(string); ok {
			elem, err = NewStringElement(v)
		} else {
			err = ErrInvalidInput
		}
		return elem, err
	})
}

// NewStringElement creates a new string element or an error.
func NewStringElement(value string) (elem Element, err error) {

	elem, err = makeBaseElement(value, StringType, func(value interface{}) (out string, e error) {
		out = strconv.Quote(value.(string))
		return out, e
	})

	return elem, err
}
