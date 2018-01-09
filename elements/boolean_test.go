package elements

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Boolean in EDN", func() {
	Context("", func() {

		It("should initialize without issue", func() {
			delete(typeFactories, BooleanType)
			err := initBoolean()
			Ω(err).Should(BeNil())
			_, has := typeFactories[BooleanType]
			Ω(has).Should(BeTrue())

			err = initBoolean()
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidFactory))
		})

		It("should create elements from the factory", func() {
			v := true

			elem, err := typeFactories[BooleanType](v)
			Ω(err).Should(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(BooleanType))
			Ω(elem.Value()).Should(BeEquivalentTo(v))
		})

		It("should not create elements from the factory if the input is not a the right type", func() {
			v := "true"

			elem, err := typeFactories[BooleanType](v)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidInput))
			Ω(elem).Should(BeNil())
		})
	})

	Context("with the default marshaller", func() {

		It("should create an true value with no error", func() {
			v := true

			elem, err := NewBooleanElement(v)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(BooleanType))
			Ω(elem.Value()).Should(BeEquivalentTo(v))
		})

		It("should fail if we hack the value on serialization", func() {
			v := true

			elem, err := NewBooleanElement(v)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(BooleanType))
			Ω(elem.Value()).Should(BeEquivalentTo(v))
		})

		It("should create an false value with no error", func() {
			v := false

			elem, err := NewBooleanElement(v)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(BooleanType))
			Ω(elem.Value()).Should(BeEquivalentTo(v))
		})

		It("should serialize true without an issue", func() {
			v := true

			elem, err := NewBooleanElement(v)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo("true"))
		})

		It("should serialize false without an issue", func() {
			v := false

			elem, err := NewBooleanElement(v)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo("false"))
		})
	})
})
