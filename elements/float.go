package elements

import "strconv"

// init will add the element factory to the collection of factories
func initFloat() error {
	return AddElementTypeFactory(FloatType, func(input interface{}) (elem Element, err error) {
		if v, ok := input.(float64); ok {
			elem, err = NewFloatElement(v)
		} else {
			err = ErrInvalidInput
		}
		return elem, err
	})
}

// NewFloatElement creates a new float point element or an error.
func NewFloatElement(value float64) (elem Element, err error) {
	elem, err = makeBaseElement(value, FloatType, func(value interface{}) (out string, e error) {
		out = strconv.FormatFloat(value.(float64), 'E', -1, 64)
		return out, e
	})

	return elem, err
}
