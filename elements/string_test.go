package elements

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("String in EDN", func() {
	Context("", func() {

		It("should initialize without issue", func() {
			delete(typeFactories, StringType)
			err := initString()
			Ω(err).Should(BeNil())
			_, has := typeFactories[StringType]
			Ω(has).Should(BeTrue())

			err = initString()
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidFactory))
		})

		It("should create elements from the factory", func() {
			v := "Hello world"

			elem, err := typeFactories[StringType](v)
			Ω(err).Should(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(StringType))
			Ω(elem.Value()).Should(BeEquivalentTo(v))
		})

		It("should not create elements from the factory if the input is not a the right type", func() {
			v := 123

			elem, err := typeFactories[StringType](v)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidInput))
			Ω(elem).Should(BeNil())
		})
	})

	Context("with the default marshaller", func() {

		testValue := "This is my test value."

		It("should create a string value with no error", func() {
			elem, err := NewStringElement(testValue)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(StringType))
			Ω(elem.Value()).Should(BeEquivalentTo(testValue))
		})

		It("should serialize the string without an issue", func() {
			elem, err := NewStringElement(testValue)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo("\"" + testValue + "\""))
		})
	})
})
