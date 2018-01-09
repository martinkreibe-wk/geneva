package elements

import (
	"fmt"
	"regexp"
	"strings"
)

const (

	// SymbolSeparator defines the symbol for separating the prefix with the name. If there is no separator, then the
	// symbol is just a name value.
	SymbolSeparator = "/"

	// NamespaceSeparator defines the symbol for separating the namespaces from each other. Note that this is not the
	// same as the namespace/name separator (SymbolSeparator).
	NamespaceSeparator = "."

	// ErrInvalidSymbol defines an invalid symbol
	ErrInvalidSymbol = Error("Invalid Symbol")

	// symbols that can modify a numeric
	numericModifierSymbols = `\.|\+|-`

	// symbols other then alphanumeric and mumeric modifiers that are legal
	legalFirstSymbols = `\*|!|_|\?|\$|%|&|=|<|>`

	// symbols that are marked as not being allowed to be first characters other then numeric
	specialSymbols = KeywordPrefix + `|` + TagPrefix

	// symbolRegex defines the valid symbols.
	// Symbols begin with a non-numeric character and can contain alphanumeric characters and . * + ! - _ ? $ % & = < >.
	// If -, + or . are the first character, the second character (if any) must be non-numeric. Additionally, : # are
	// allowed as constituent characters in symbols other than as the first character.
	symbolRegex = `^((` + numericModifierSymbols + `)|((((` + numericModifierSymbols + `)(` + legalFirstSymbols + `|[[:alpha:]]))|(` + legalFirstSymbols + `|[[:alpha:]]))+(` + numericModifierSymbols + `|` + legalFirstSymbols + `|` + specialSymbols + `|[[:alnum:]])*))$`
)

// symbolMatcher is the matching mechanism for symbols
var symbolMatcher = regexp.MustCompile(symbolRegex).MatchString

// Symbols are used to represent identifiers, and should map to something other than strings, if possible.
type SymbolElement interface {
	Element

	// Modifier for this symbol
	Modifier() string

	// Prefix to this symbol
	Prefix() string

	// Name to this symbol
	Name() string
}

// symbolElemImpl implements the symbolElemImpl
type symbolElemImpl struct {
	*baseElemImpl
	prefix   string
	name     string
	modifier string
}

// NewSymbolElement creates a new character element or an error.
func NewSymbolElement(parts ...string) (elem SymbolElement, err error) {

	var prefix string
	var name string

	switch len(parts) {
	case 1:

		switch name = parts[0]; {

		// handle the case where the name was really sent in with the separator
		case name == SymbolSeparator:
			// Fine, break

		case strings.Contains(name, SymbolSeparator):
			if parts = strings.Split(name, SymbolSeparator); len(parts) == 2 {
				if prefix = parts[0]; len(prefix) != 0 && symbolMatcher(prefix) {
					if name = parts[1]; len(name) == 0 || !symbolMatcher(name) {
						err = ErrInvalidSymbol
					}
				} else {
					err = ErrInvalidSymbol
				}
			} else {
				err = ErrInvalidSymbol
			}
		default:
			if !symbolMatcher(name) {
				err = ErrInvalidSymbol
			}
		}

	case 2:
		if prefix = parts[0]; len(prefix) != 0 && symbolMatcher(prefix) {
			if name = parts[1]; !symbolMatcher(name) {
				err = ErrInvalidSymbol
			}
		} else {
			err = ErrInvalidSymbol
		}
	default:
		err = ErrInvalidSymbol
	}

	if err == nil {

		symElem := &symbolElemImpl{
			prefix: prefix,
			name:   name,
		}

		var base *baseElemImpl
		if base, err = makeBaseElement(symElem, SymbolType, func(value interface{}) (out string, err error) {
			if elem, ok := value.(*symbolElemImpl); ok {
				symbol := elem.Name()
				if len(prefix) > 0 {
					symbol = fmt.Sprintf("%s%s%s", elem.Prefix(), SymbolSeparator, symbol)
				}

				out = elem.Modifier() + symbol
			}

			return out, err
		}); err == nil {

			symElem.baseElemImpl = base

			// equality for symbols are different then the normal path.
			symElem.baseElemImpl.equality = func(left, right Element) (result bool) {
				if leftSym, has := left.(SymbolElement); has {
					if rightSym, has := right.(SymbolElement); has {
						if leftSym.Name() == rightSym.Name() && leftSym.Prefix() == rightSym.Prefix() && leftSym.Modifier() == rightSym.Modifier() {
							result = true
						}
					}
				}

				return result
			}

			elem = symElem
		}
	}

	return elem, err
}

// Equals checks if the input element is equal to this element.
func (elem *symbolElemImpl) Equals(e Element) (result bool) {
	if elem.ElementType() == e.ElementType() {
		if elem.Tag() == e.Tag() {
			result = elem.baseElemImpl.equality(elem, e)
		}
	}
	return result
}

// Prefix to this symbol
func (elem *symbolElemImpl) Prefix() string {
	return elem.prefix
}

// Name to this symbol
func (elem *symbolElemImpl) Name() string {
	return elem.name
}

// Modifier for this symbol
func (elem *symbolElemImpl) Modifier() string {
	return elem.modifier
}
