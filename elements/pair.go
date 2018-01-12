package elements

const (

	// ErrInvalidPair defines the error for pair issues
	ErrInvalidPair = Error("Invalid pair")
)

// Pair of elements
type Pair interface {

	// Key the key part of the pair
	Key() Element

	// Value the value part of the pair
	Value() Element
}

// pairImpl implements the pair interface
type pairImpl struct {
	key   Element
	value Element
}

// Key the key part of the pair
func (pair *pairImpl) Key() Element {
	return pair.key
}

// Value the value part of the pair
func (pair *pairImpl) Value() Element {
	return pair.value
}

// NewPair creates a new pair from the key and value supplied.
func NewPair(key, value interface{}) (pair Pair, err error) {

	if key != nil && value != nil {
		var k, v Element
		if k, err = NewElement(key); err == nil {
			if v, err = NewElement(value); err == nil {
				pair = &pairImpl{
					key:   k,
					value: v,
				}
			}
		}
	} else {
		err = ErrInvalidPair
	}

	return pair, err
}

// Pairs defines a collection of pairs.
type Pairs struct {
	data []Pair
}

// Append will append the entities
func (pairs *Pairs) Append(key, value Element) (err error) {

	var pair Pair
	if pair, err = NewPair(key, value); err == nil {
		pairs.data = append(pairs.data, pair)
	}

	return err
}

// Raw returns the internal pair collection
func (pairs *Pairs) Raw() []Pair {
	return pairs.data
}

// Len returns the pair collection length
func (pairs *Pairs) Len() int {
	return len(pairs.data)
}
