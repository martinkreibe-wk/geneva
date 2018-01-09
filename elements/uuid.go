package elements

import (
	"github.com/mattrobenolt/gocql/uuid"
)

const (

	// UUIDElementTag defines the uuid tag value.
	UUIDElementTag = "uuid"
)

// init will add the element factory to the collection of factories
func initUUID() error {
	return AddElementTypeFactory(UUIDType, func(input interface{}) (elem Element, err error) {
		if v, ok := input.(uuid.UUID); ok {
			elem, err = NewUUIDElement(v)
		} else {
			err = ErrInvalidInput
		}
		return elem, err
	})
}

// NewInstantElement creates a new instant element or an error.
func NewUUIDElement(value uuid.UUID) (elem Element, err error) {
	if elem, err = makeBaseElement(value, UUIDType, func(value interface{}) (out string, e error) {
		out = value.(uuid.UUID).String()
		return out, e
	}); err == nil {
		elem.SetTag(UUIDElementTag)
	}

	return elem, err
}
