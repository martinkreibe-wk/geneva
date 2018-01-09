package elements

const (

	// MapStartLiteral is the start of an EDN group element.
	MapStartLiteral = "{"

	// MapEndLiteral is the end of an EDN group element.
	MapEndLiteral = "}"

	// MapSeparatorLiteral is the separator between item in a collection
	MapSeparatorLiteral = ", "

	// MapKeyValueSeparatorLiteral is the separator for keys and values
	MapKeyValueSeparatorLiteral = " "

	// ErrDuplicateKey defines the duplicate key error
	ErrDuplicateKey = Error("Duplicate key found")
)

// NewMap creates a new vector
func NewMap(pairs ...Pair) (elem CollectionElement, err error) {

	// check for errors
	keys := []Element{}
	children := map[Element]Element{}
	for _, pair := range pairs {
		if pair == nil || pair.Key() == nil {
			err = ErrInvalidPair
		} else {

			key := pair.Key()
			for _, k := range keys {

				if key.Equals(k) {
					err = ErrDuplicateKey
					break
				}
			}

			if err == nil {
				keys = append(keys, key)
				children[key] = pair.Value()
			}
		}

		if err != nil {
			break
		}
	}

	if err == nil {
		coll := &collectionElemImpl{
			startSymbol:             MapStartLiteral,
			endSymbol:               MapEndLiteral,
			separatorSymbol:         MapSeparatorLiteral,
			keyValueSeparatorSymbol: MapKeyValueSeparatorLiteral,
			collection:              map[string]Element{},
		}

		var base *baseElemImpl
		if base, err = makeBaseElement(coll, MapType, collectionSerialization(true)); err == nil {
			coll.baseElemImpl = base
			elem = coll
			for key, value := range children {
				if err = elem.Append(key, value); err != nil {
					break
				}
			}
		}
	}

	return elem, err
}
