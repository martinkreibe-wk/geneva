package elements

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Character in EDN", func() {
	Context("", func() {

		It("should initialize without issue", func() {
			delete(typeFactories, CharacterType)
			err := initCharacter()
			Ω(err).Should(BeNil())
			_, has := typeFactories[CharacterType]
			Ω(has).Should(BeTrue())

			err = initCharacter()
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidFactory))
		})

		It("should create elements from the factory", func() {
			v := 'g'

			elem, err := typeFactories[CharacterType](v)
			Ω(err).Should(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(CharacterType))
			Ω(elem.Value()).Should(BeEquivalentTo(v))
		})

		It("should not create elements from the factory if the input is not a the right type", func() {
			v := "foo"

			elem, err := typeFactories[CharacterType](v)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidInput))
			Ω(elem).Should(BeNil())
		})
	})

	Context("with the default marshaller", func() {

		c := 'c'
		runes := map[rune]string{
			c:    "\\c",
			'\n': "\\newline",
			'\r': "\\return",
			' ':  "\\space",
			'\t': "\\tab",
			'⌘':  "\\u2318",
		}

		It("should create an character value with no error", func() {
			elem, err := NewCharacterElement(c)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(CharacterType))
			Ω(elem.Value()).Should(BeEquivalentTo(c))
		})

		It("should serialize the character without an issue", func() {

			for r, ser := range runes {
				elem, err := NewCharacterElement(r)
				Ω(err).Should(BeNil())
				Ω(elem).ShouldNot(BeNil())
				Ω(elem.Value()).Should(BeEquivalentTo(r), fmt.Sprintf("For rune: %+q", r))

				var edn string
				edn, err = elem.Serialize()
				Ω(err).Should(BeNil())
				Ω(edn).Should(BeEquivalentTo(ser), fmt.Sprintf("For rune: %+q", r))
			}
		})
	})
})
