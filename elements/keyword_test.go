package elements

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keyword in EDN", func() {
	Context("", func() {

		It("should initialize without issue", func() {
			delete(typeFactories, KeywordType)
			err := initKeyword()
			Ω(err).Should(BeNil())
			_, has := typeFactories[KeywordType]
			Ω(has).Should(BeTrue())

			err = initKeyword()
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidFactory))
		})

		It("should create elements from the factory", func() {
			v := "testKeyword"

			elem, err := typeFactories[KeywordType](v)
			Ω(err).Should(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(KeywordType))

			symbol, is := elem.(SymbolElement)
			Ω(is).Should(BeTrue())
			Ω(symbol.Name()).Should(BeEquivalentTo(v))
		})

		It("should not create elements from the factory if the input is not a the right type", func() {
			v := 123

			elem, err := typeFactories[KeywordType](v)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidInput))
			Ω(elem).Should(BeNil())
		})
	})

	Context("with the default marshaller", func() {

		badKeywords := []string{
			":/",
			"::",
			":",
			"/",
			":1",
			"1",
			"bad1/1worse",
			"/bad",
			"bad/",
			"bad/worse/wrong",
			"worse/+1bad",
			"+1bad",
			"-1bad",
			".1bad",
			".1",
		}

		It("should create a keyword with just one parameter value with no error", func() {
			prefix := ""
			name := "foobar"

			elem, err := NewKeywordElement(name)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(KeywordType))
			Ω(elem.Prefix()).Should(BeEquivalentTo(prefix))
			Ω(elem.Name()).Should(BeEquivalentTo(name))
		})

		It("should create a keyword with two parameter value with no error", func() {
			prefix := "namespace"
			name := "foobar"

			elem, err := NewKeywordElement(prefix, name)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(KeywordType))
			Ω(elem.Prefix()).Should(BeEquivalentTo(prefix))
			Ω(elem.Name()).Should(BeEquivalentTo(name))
		})

		It("should create a keyword with one parameter (but with the separator) value with no error", func() {
			prefix := "namespace"
			name := "foobar"

			elem, err := NewKeywordElement(prefix + SymbolSeparator + name)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(KeywordType))
			Ω(elem.Prefix()).Should(BeEquivalentTo(prefix))
			Ω(elem.Name()).Should(BeEquivalentTo(name))
		})

		It("should create a keyword with just one parameter (first : prefixed) value with no error", func() {
			prefix := ""
			name := "foobar"

			elem, err := NewKeywordElement(KeywordPrefix + name)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(KeywordType))
			Ω(elem.Prefix()).Should(BeEquivalentTo(prefix))
			Ω(elem.Name()).Should(BeEquivalentTo(name))
		})

		It("should create a keyword with two parameter (first : prefixed) value with no error", func() {
			prefix := "namespace"
			name := "foobar"

			elem, err := NewKeywordElement(KeywordPrefix+prefix, name)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(KeywordType))
			Ω(elem.Prefix()).Should(BeEquivalentTo(prefix))
			Ω(elem.Name()).Should(BeEquivalentTo(name))
		})

		It("should create a keyword with one parameter (: prefixed and with the separator) value with no error", func() {
			prefix := "namespace"
			name := "foobar"

			elem, err := NewKeywordElement(KeywordPrefix + prefix + SymbolSeparator + name)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(KeywordType))
			Ω(elem.Prefix()).Should(BeEquivalentTo(prefix))
			Ω(elem.Name()).Should(BeEquivalentTo(name))
		})

		It("should create a keyword with zero parameter value with an error", func() {
			elem, err := NewKeywordElement()
			Ω(err).ShouldNot(BeNil())
			Ω(elem).Should(BeNil())
			Ω(err).Should(BeIdenticalTo(ErrInvalidKeyword))
		})

		It("should create a keyword with three parameter value with an error", func() {
			elem, err := NewKeywordElement("a", "b", "c")
			Ω(err).ShouldNot(BeNil())
			Ω(elem).Should(BeNil())
			Ω(err).Should(BeIdenticalTo(ErrInvalidKeyword))
		})

		It("should serialize the keyword with one parameter without an issue", func() {
			name := "foobar"

			elem, err := NewKeywordElement(name)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo(KeywordPrefix + name))
		})

		It("should serialize the keyword with two parameter without an issue", func() {
			prefix := "namespace"
			name := "foobar"

			elem, err := NewKeywordElement(prefix, name)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo(KeywordPrefix + prefix + SymbolSeparator + name))
		})

		It("should serialize the keyword with one (but with the separator) parameter without an issue", func() {
			prefix := "namespace"
			name := "foobar"

			elem, err := NewKeywordElement(prefix + SymbolSeparator + name)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo(KeywordPrefix + prefix + SymbolSeparator + name))
		})

		It("should serialize the keyword with one parameter (with : prefix) without an issue", func() {
			name := "foobar"

			elem, err := NewKeywordElement(KeywordPrefix + name)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo(KeywordPrefix + name))
		})

		It("should serialize the keyword with two parameter (with : prefix) without an issue", func() {
			prefix := "namespace"
			name := "foobar"

			elem, err := NewKeywordElement(KeywordPrefix+prefix, name)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo(KeywordPrefix + prefix + SymbolSeparator + name))
		})

		It("should serialize the keyword with one (with : prefix and with the separator) parameter without an issue", func() {
			prefix := "namespace"
			name := "foobar"

			elem, err := NewKeywordElement(KeywordPrefix + prefix + SymbolSeparator + name)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo(KeywordPrefix + prefix + SymbolSeparator + name))
		})

		It("should not process all odd invalid keywords", func() {

			for _, keyword := range badKeywords {
				elem, err := NewKeywordElement(keyword)
				Ω(elem).Should(BeNil())
				Ω(err).ShouldNot(BeNil())
				Ω(err).Should(BeIdenticalTo(ErrInvalidKeyword))
			}
		})
	})
})
