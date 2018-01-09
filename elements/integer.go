package elements

import (
	"strconv"
)

// init will add the element factory to the collection of factories
func initInteger() error {
	return AddElementTypeFactory(IntegerType, func(input interface{}) (elem Element, err error) {
		if v, ok := input.(int64); ok {
			elem, err = NewIntegerElement(v)
		} else {
			err = ErrInvalidInput
		}
		return elem, err
	})
}

// NewIntegerElement creates a new integer element or an error.
func NewIntegerElement(value int64) (elem Element, err error) {
	elem, err = makeBaseElement(value, IntegerType, func(value interface{}) (out string, e error) {
		out = strconv.FormatInt(value.(int64), 10)
		return out, e
	})

	return elem, err
}
