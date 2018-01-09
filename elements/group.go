package elements

const (

	// GroupingStartLiteral is the start of an EDN group element.
	GroupingStartLiteral = "("

	// GroupingEndLiteral is the end of an EDN group element.
	GroupingEndLiteral = ")"

	// GroupingSeparatorLiteral is the separator between item in a collection
	GroupingSeparatorLiteral = " "
)

// NewGroup creates a new group
func NewGroup(elements ...Element) (elem CollectionElement, err error) {

	// check for errors
	for _, child := range elements {
		if child == nil {
			err = ErrInvalidElement
			break
		}
	}

	if err == nil {
		coll := &collectionElemImpl{
			startSymbol:     GroupingStartLiteral,
			endSymbol:       GroupingEndLiteral,
			separatorSymbol: GroupingSeparatorLiteral,
			collection:      []Element{},
		}

		var base *baseElemImpl
		if base, err = makeBaseElement(coll, GroupingType, collectionSerialization(false)); err == nil {
			coll.baseElemImpl = base
			elem = coll
			err = elem.Append(elements...)
		}
	}

	return elem, err
}
