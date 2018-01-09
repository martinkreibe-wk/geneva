package elements

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Integer in EDN", func() {
	Context("", func() {

		It("should initialize without issue", func() {
			delete(typeFactories, IntegerType)
			err := initInteger()
			Ω(err).Should(BeNil())
			_, has := typeFactories[IntegerType]
			Ω(has).Should(BeTrue())

			err = initInteger()
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidFactory))
		})

		It("should create elements from the factory", func() {
			v := int64(123)

			elem, err := typeFactories[IntegerType](v)
			Ω(err).Should(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(IntegerType))
			Ω(elem.Value()).Should(BeEquivalentTo(v))
		})

		It("should not create elements from the factory if the input is not a the right type", func() {
			v := "foo"

			elem, err := typeFactories[IntegerType](v)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidInput))
			Ω(elem).Should(BeNil())
		})
	})

	Context("with the default marshaller", func() {

		testValue := int64(12345)

		It("should create an integer value with no error", func() {
			elem, err := NewIntegerElement(testValue)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(IntegerType))
			Ω(elem.Value()).Should(BeEquivalentTo(testValue))
		})

		It("should serialize the integer without an issue", func() {
			elem, err := NewIntegerElement(testValue)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo("12345"))
		})
	})
})
