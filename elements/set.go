package elements

const (

	// SetStartLiteral is the start of an EDN group element.
	SetStartLiteral = "#{"

	// SetEndLiteral is the end of an EDN group element.
	SetEndLiteral = "}"

	// GroupingSeparatorLiteral is the separator between item in a collection
	SetSeparatorLiteral = " "
)

// NewSet creates a new vector
func NewSet(elements ...Element) (elem CollectionElement, err error) {

	// check for errors
	for _, child := range elements {
		if child == nil {
			err = ErrInvalidElement
			break
		}
	}

	if err == nil {
		coll := &collectionElemImpl{
			startSymbol:     SetStartLiteral,
			endSymbol:       SetEndLiteral,
			separatorSymbol: SetSeparatorLiteral,
			collection:      []Element{},
		}

		var base *baseElemImpl
		if base, err = makeBaseElement(coll, SetType, collectionSerialization(false)); err == nil {
			coll.baseElemImpl = base
			elem = coll
			err = elem.Append(elements...)
		}
	}

	return elem, err
}
