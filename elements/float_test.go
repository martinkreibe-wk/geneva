package elements

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Float in EDN", func() {
	Context("", func() {

		It("should initialize without issue", func() {
			delete(typeFactories, FloatType)
			err := initFloat()
			Ω(err).Should(BeNil())
			_, has := typeFactories[FloatType]
			Ω(has).Should(BeTrue())

			err = initFloat()
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidFactory))
		})

		It("should create elements from the factory", func() {
			v := float64(1.234)

			elem, err := typeFactories[FloatType](v)
			Ω(err).Should(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(FloatType))
			Ω(elem.Value()).Should(BeEquivalentTo(v))
		})

		It("should not create elements from the factory if the input is not a the right type", func() {
			v := "foo"

			elem, err := typeFactories[FloatType](v)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidInput))
			Ω(elem).Should(BeNil())
		})
	})

	Context("with the default marshaller", func() {

		testValue := float64(12345.67)

		It("should create an float value with no error", func() {
			elem, err := NewFloatElement(testValue)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(FloatType))
			Ω(elem.Value()).Should(BeEquivalentTo(testValue))
		})

		It("should serialize the float without an issue", func() {
			elem, err := NewFloatElement(testValue)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo("1.234567E+04"))
		})
	})
})
