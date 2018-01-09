package elements

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Nil in EDN", func() {
	Context("", func() {

		It("should initialize without issue", func() {
			delete(typeFactories, NilType)
			err := initNil()
			Ω(err).Should(BeNil())
			_, has := typeFactories[NilType]
			Ω(has).Should(BeTrue())

			err = initNil()
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidFactory))
		})

		It("should create elements from the factory", func() {
			var v interface{}

			elem, err := typeFactories[NilType](v)
			Ω(err).Should(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(NilType))
			Ω(elem.Value()).Should(BeNil())
		})

		It("should not create elements from the factory if the input is not a the right type", func() {
			v := "foo"

			elem, err := typeFactories[NilType](v)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidInput))
			Ω(elem).Should(BeNil())
		})
	})

	Context("with the default marshaller", func() {

		It("should create an nil with no error", func() {
			elem, err := NewNilElement()
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(NilType))
		})

		It("should serialize without an issue", func() {
			elem, err := NewNilElement()
			Ω(err).Should(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo(NilLiteral))
		})
	})
})
