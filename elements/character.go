package elements

import (
	"fmt"
	"strings"
)

const (

	// CharacterPrefix defines the prefix for characters
	CharacterPrefix = "\\"
)

var specialCharacters = map[rune]string{
	'\r': CharacterPrefix + "return",
	'\n': CharacterPrefix + "newline",
	' ':  CharacterPrefix + "space",
	'\t': CharacterPrefix + "tab",
}

// init will add the element factory to the collection of factories
func initCharacter() error {
	return AddElementTypeFactory(CharacterType, func(input interface{}) (elem Element, err error) {
		if v, ok := input.(rune); ok {
			elem, err = NewCharacterElement(v)
		} else {
			err = ErrInvalidInput
		}
		return elem, err
	})
}

// NewCharacterElement creates a new character element or an error.
func NewCharacterElement(value rune) (elem Element, err error) {
	elem, err = makeBaseElement(value, CharacterType, func(value interface{}) (out string, e error) {

		r := value.(rune)

		var has bool
		if out, has = specialCharacters[r]; !has {

			// if there is no special character, then quote the rune, remove the single quotes around this, then
			// if it is an ASCII then make sure to prefix is intact.
			if out = strings.Trim(fmt.Sprintf("%+q", r), "'"); !strings.HasPrefix(out, CharacterPrefix) {
				out = CharacterPrefix + out
			}
		}

		return out, e
	})

	return elem, err
}
