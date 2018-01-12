package elements

import (
	"strings"
)

const (

	// KeywordPrefix defines the prefix for keywords
	KeywordPrefix = ":"

	// ErrInvalidKeyword defines the error for invalid keywords
	ErrInvalidKeyword = Error("Invalid keyword")

	// ForwardDirection
	ForwardDirection KeywordDirection = ""

	// ReverseDirection
	ReverseDirection KeywordDirection = "_"
)

// KeywordDirection
type KeywordDirection string

// init will add the element factory to the collection of factories
func initKeyword() error {
	return AddElementTypeFactory(KeywordType, func(input interface{}) (elem Element, err error) {
		if v, ok := input.(string); ok {
			elem, err = NewKeywordElement(v)
		} else {
			err = ErrInvalidInput
		}
		return elem, err
	})
}

// Keywords are identifiers that typically designate themselves. They are semantically akin to enumeration values.
// Keywords follow the rules of symbols, except they can (and must) begin with :, e.g. :fred or :my/fred. If the target
// platform does not have a keyword type distinct from a symbol type, the same type can be used without conflict, since
// the mandatory leading : of keywords is disallowed for symbols. Per the symbol rules above, :/ and :/anything are not
// legal keywords. A keyword cannot begin with ::

// If the target platform supports some notion of interning, it is a further semantic of keywords that all instances of
// the same keyword yield the identical object.
type KeywordElement interface {
	SymbolElement

	// SetDirection of the keyword
	SetDirection(KeywordDirection)
}

// NewKeywordElement creates a new character element or an error.
func NewKeywordElement(parts ...string) (elem KeywordElement, err error) {

	// remove the : symbol if it is the first character.
	switch len(parts) {
	case 0:
		err = ErrInvalidKeyword

	default:
		if strings.HasPrefix(parts[0], KeywordPrefix) {
			parts[0] = strings.TrimPrefix(parts[0], KeywordPrefix)
		}

		// Per the symbol rules above, :/ and :/anything are not legal keywords.
		if strings.HasPrefix(parts[0], SymbolSeparator) {
			err = ErrInvalidKeyword
		}
	}

	if err == nil {
		var symbol SymbolElement
		if symbol, err = NewSymbolElement(parts...); err == nil {

			impl := symbol.(*symbolElemImpl)
			impl.baseElemImpl.elemType = KeywordType
			impl.modifier = KeywordPrefix

			elem = impl
		}
	}

	if err == ErrInvalidSymbol {
		err = ErrInvalidKeyword
	}

	return elem, err
}

// SetDirection of the keyword
func (elem *symbolElemImpl) SetDirection(direction KeywordDirection) {
	elem.direction = direction
}
