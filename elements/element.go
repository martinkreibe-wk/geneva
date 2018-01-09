package elements

import (
	"reflect"
	"time"

	"github.com/mattrobenolt/gocql/uuid"
)

const (

	// InvalidElement defines an invalid element was encountered.
	ErrInvalidElement = Error("Invalid Element")

	// TagPrefix defines the prefix for tags.
	TagPrefix = "#"
)

// Element defines the interface for EDN elements.
type Element interface {

	// ElementType returns the current type of this element.
	ElementType() ElementType

	// Value of the element
	Value() interface{}

	// Serialize the element into a string or return the appropriate error.
	Serialize() (composition string, err error)

	// HasTag returns true if the element has a tag prefix
	HasTag() bool

	// Tag returns the prefixed tag if it exists.
	Tag() string

	// SetTag sets the tag to the incoming value. If the value is an empty string then the tag is unset.
	SetTag(string) (err error)

	// Equals checks if the input element is equal to this element.
	Equals(e Element) (result bool)
}

// NewElement creates a new element from the inputs. I f the first parameter is a ElementType, then that will stereotype
// the rest of the values. For more then 2 inputs, a collection will be assumed and if they set a non collection
// ElementType is a non collection, then this will error.
func NewElement(value ...interface{}) (elem Element, err error) {

	var stereotype ElementType

	switch {
	case len(value) == 0:
		elem, err = NewNilElement()
	case len(value) == 1:
		if e, ok := value[0].(Element); ok {
			elem = e
		}
	case len(value) <= 2:
		switch v := value[0].(type) {
		case ElementType:
			if (v.IsCollection() && len(value) < 2) || (!v.IsCollection() && len(value) == 2) {
				stereotype = v
				value = value[1:]
			} else {
				err = ErrInvalidElement
			}
		}
	}

	if elem == nil && err == nil {
		switch {
		case stereotype.IsCollection():

		case len(value) == 1:
			val := value[0]

			if stereotype == UnknownType {
				switch v := val.(type) {
				case int32:
					stereotype = IntegerType
					val = int64(v)
				case int64:
					stereotype = IntegerType
				case float32:
					stereotype = FloatType
					val = float64(v)
				case float64:
					stereotype = FloatType
				case string:
					if v == NilLiteral {
						stereotype = NilType
						val = nil
					} else {
						stereotype = StringType
						if v[0] == ':' {
							stereotype = KeywordType
						}
					}
				case time.Time:
					stereotype = InstantType
				case uuid.UUID:
					stereotype = UUIDType
				}
			}

			if factory, has := typeFactories[stereotype]; has {
				elem, err = factory(val)
			} else {
				err = AppendError(ErrInvalidElement, NewError("Unknown type", stereotype.Name()))
			}

		default:
			err = ErrInvalidElement
		}
	}

	return elem, err
}

// elemStringer defines the mechanism to stringify the element.
type elemStringer func(interface{}) (string, error)

// elemEqualityChecker defines the mechanism testing equality
type elemEqualityChecker func(left, right Element) bool

// baseElement defines the base element features.
type baseElemImpl struct {

	// elemType is the type this element houses.
	elemType ElementType

	// stringer is the mechanism to serialize this element.
	stringer elemStringer

	// equality is the tester for equality
	equality elemEqualityChecker

	// tag of this element.
	tag string

	// value of this element.
	value interface{}
}

// makeBaseElement creates the base element.
func makeBaseElement(value interface{}, elementType ElementType, stringer elemStringer) (elem *baseElemImpl, err error) {

	if stringer != nil {
		elem = &baseElemImpl{
			elemType: elementType,
			stringer: stringer,
			value:    value,
			equality: func(left, right Element) (result bool) {
				return reflect.DeepEqual(left.Value(), right.Value())
			},
		}
	} else {
		err = ErrInvalidElement
	}

	return elem, err
}

// Equals checks if the input element is equal to this element.
func (elem *baseElemImpl) Equals(e Element) (result bool) {
	if elem.ElementType() == e.ElementType() {
		if elem.Tag() == e.Tag() {
			result = elem.equality(elem, e)
		}
	}
	return result
}

// ElementType returns the current type of this element.
func (elem *baseElemImpl) ElementType() ElementType {
	return elem.elemType
}

// Serialize the element into a string or return the appropriate error.
func (elem *baseElemImpl) Serialize() (composition string, err error) {

	// If the tag exists then prefix the value with the tag.
	if elem.HasTag() {
		composition = TagPrefix + elem.Tag() + " "
	}

	var comp string
	if comp, err = elem.stringer(elem.Value()); err == nil {
		composition += comp
	}
	return composition, err
}

// HasTag returns true if the element has a tag prefix
func (elem *baseElemImpl) HasTag() bool {
	return len(elem.tag) != 0
}

// Tag returns the prefixed tag if it exists.
func (elem *baseElemImpl) Tag() string {
	return elem.tag
}

// tagged elements
//
// edn supports extensibility through a simple mechanism. # followed immediately by a symbol starting with an alphabetic
// character indicates that that symbol is a tag. A tag indicates the semantic interpretation of the following element.
// It is envisioned that a reader implementation will allow clients to register handlers for specific tags. Upon
// encountering a tag, the reader will first read the next element (which may itself be or comprise other tagged
// elements), then pass the result to the corresponding handler for further interpretation, and the result of the
// handler will be the data value yielded by the tag + tagged element, i.e. reading a tag and tagged element yields one
// value. This value is the value to be returned to the program and is not further interpreted as edn data by the
// reader.
//
// This process will bottom out on elements either understood or built-in.
//
// Thus you can build new distinct readable elements out of (and only out of) other readable elements, keeping extenders
// and extension consumers out of the text business.
//
// The semantics of a tag, and the type and interpretation of the tagged element are defined by the steward of the tag.
//
// #myapp/Person {:first "Fred" :last "Mertz"}
//
// If a reader encounters a tag for which no handler is registered, the implementation can either report an error, call
// a designated 'unknown element' handler, or create a well-known generic representation that contains both the tag and
// the tagged element, as it sees fit. Note that the non-error strategies allow for readers which are capable of reading
// any and all edn, in spite of being unaware of the details of any extensions present.

// rules for tags
//
// Tag symbols without a prefix are reserved by edn for built-ins defined using the tag system.
// User tags must contain a prefix component, which must be owned by the user (e.g. trademark or domain) or known unique
// in the communication context.
// A tag may specify more than one format for the tagged element, e.g. both a string and a vector representation.
// Tags themselves are not elements. It is an error to have a tag without a corresponding tagged element.

// SetTag sets the tag to the incoming value. If the value is an empty string then the tag is unset.
func (elem *baseElemImpl) SetTag(value string) (err error) {

	// TODO: check if the first is # and remove
	// TODO: make sure tag is not ws or ,
	// TODO: Follow the rules above

	elem.tag = value

	return err
}

// Value return the raw representation of this element.
func (elem *baseElemImpl) Value() interface{} {
	return elem.value
}
