package elements

import (
	"fmt"
	"strconv"
)

const (

	// ErrNoValue is returned when no value is found in the collection.
	ErrNoValue = Error("No value found")
)

// ChildIterator is the iterator for children elements, if this is a list based item, the key will be an index of the
// current item, the value is mapped value to the key. To stop a loop mid iteration, set the error to non nil.
type ChildIterator func(key Element, value Element) (err error)

// GroupElement defines the element for the EDN grouping construct. A group is a sequence of values. Groups are
// represented by zero or more elements enclosed in parentheses (). Note that Group can be heterogeneous.
type CollectionElement interface {
	Element

	// Len return the quantity of items in this collection.
	Len() int

	// IterateChildren will iterator over all children.
	IterateChildren(iterator ChildIterator) (err error)

	// Append the elements into this collection.
	Append(children ...Element) (err error)

	// Get the key from the collection.
	Get(key interface{}) (Element, error)
}

// collectionElemImpl is the implementation to the GroupElement interface.
type collectionElemImpl struct {
	*baseElemImpl

	// startSymbol defines the start symbol
	startSymbol string

	// endSymbol defines the end symbol
	endSymbol string

	// separatorSymbol for the elements
	separatorSymbol string

	// keyValueSeparatorSymbol for the element
	keyValueSeparatorSymbol string

	// collection of elements
	collection interface{}
}

// Len return the quantity of items in this collection.
func (elem *collectionElemImpl) Len() (l int) {
	switch v := elem.collection.(type) {
	case []Element:
		l = len(v)
	case map[string]Element:
		l = len(v)
	}
	return l
}

// IterateChildren will iterate over the child elements within this collection.
func (elem *collectionElemImpl) IterateChildren(iterator ChildIterator) (err error) {
	switch v := elem.collection.(type) {
	case []Element:
		for i, c := range v {
			iElem, _ := NewElement(int64(i))
			if err = iterator(iElem, c); err != nil {
				break
			}
		}
	case map[string]Element:
		for i, c := range v {

			var k Element
			if k, err = NewElement(i); err == nil {
				err = iterator(k, c)
			}

			if err != nil {
				break
			}
		}
	}
	return err
}

// collectionSerialization the element into a string or return the appropriate error.
func collectionSerialization(hasKey bool) func(value interface{}) (composition string, err error) {

	return func(value interface{}) (composition string, err error) {
		val := value.(*collectionElemImpl)
		composition = val.startSymbol

		first := true
		if err = val.IterateChildren(func(key Element, child Element) (e error) {
			if first {
				first = false
			} else {
				composition += val.separatorSymbol
			}

			var c string
			if hasKey {
				if c, e = key.Serialize(); e == nil {
					composition += c + val.keyValueSeparatorSymbol
				}
			}

			if e == nil {
				if c, e = child.Serialize(); e == nil {
					composition += c
				}
			}

			return e
		}); err == nil {
			composition += val.endSymbol
		}

		return composition, err
	}
}

// Append will add the appropriate children. Note that a map must have 2 parameters.
func (elem *collectionElemImpl) Append(children ...Element) (err error) {

	if len(children) != 0 {
		switch v := elem.collection.(type) {
		case []Element:
			elem.collection = append(v, children...)
		case map[string]Element:

			if len(children)%2 == 0 {
				for i := 0; i < len(children); i += 2 {
					var k string
					if str, is := children[i].(*baseElemImpl); is && str.elemType == StringType {
						k = str.value.(string)
					} else {
						k, err = children[i].Serialize()
					}
					if err == nil {
						v[k] = children[i+1]
					}
				}
			} else {
				err = ErrInvalidInput
			}

		default:
			err = ErrInvalidElement
		}
	}

	return err
}

// Get the value from the collection.
func (elem *collectionElemImpl) Get(key interface{}) (value Element, err error) {

	var realKey string

	switch k := key.(type) {
	case int, int32, int64:
		realKey = fmt.Sprintf("%d", k)
	case string:
		realKey = k
	case Element:
		realKey, err = k.Serialize()
	default:
		err = ErrInvalidInput
	}

	if err == nil {
		switch v := elem.collection.(type) {
		case []Element:
			var index int
			err = ErrNoValue
			if index, err = strconv.Atoi(realKey); err == nil {
				if index >= 0 && index < len(v) {
					value = v[index]
					err = nil
				}
			}
		case map[string]Element:
			var has bool
			if value, has = v[realKey]; !has {
				err = ErrNoValue
			}
		default:
			err = ErrInvalidElement
		}
	}

	return value, err
}
