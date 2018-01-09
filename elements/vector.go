package elements

const (

	// VectorStartLiteral is the start of an EDN group element.
	VectorStartLiteral = "["

	// VectorEndLiteral is the end of an EDN group element.
	VectorEndLiteral = "]"

	// GroupingSeparatorLiteral is the separator between item in a collection
	VectorSeparatorLiteral = " "
)

// NewVector creates a new vector
func NewVector(elements ...Element) (elem CollectionElement, err error) {

	// check for errors
	for _, child := range elements {
		if child == nil {
			err = ErrInvalidElement
			break
		}
	}

	if err == nil {
		coll := &collectionElemImpl{
			startSymbol:     VectorStartLiteral,
			endSymbol:       VectorEndLiteral,
			separatorSymbol: VectorSeparatorLiteral,
			collection:      []Element{},
		}

		var base *baseElemImpl
		if base, err = makeBaseElement(coll, VectorType, collectionSerialization(false)); err == nil {
			coll.baseElemImpl = base
			elem = coll
			err = elem.Append(elements...)
		}
	}

	return elem, err
}
