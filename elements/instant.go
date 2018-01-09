package elements

import (
	"time"
)

const (

	// InstantElementTag defines the instant tag value.
	InstantElementTag = "inst"
)

// init will add the element factory to the collection of factories
func initInstant() error {
	return AddElementTypeFactory(InstantType, func(input interface{}) (elem Element, err error) {
		if v, ok := input.(time.Time); ok {
			elem, err = NewInstantElement(v)
		} else {
			err = ErrInvalidInput
		}
		return elem, err
	})
}

// NewInstantElement creates a new instant element or an error.
func NewInstantElement(value time.Time) (elem Element, err error) {

	if elem, err = makeBaseElement(value, InstantType, func(value interface{}) (out string, e error) {
		out = value.(time.Time).Format(time.RFC3339)
		return out, e
	}); err == nil {
		elem.SetTag(InstantElementTag)
	}

	return elem, err
}
