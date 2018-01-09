package elements

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Instant in EDN", func() {
	Context("", func() {

		It("should initialize without issue", func() {
			delete(typeFactories, InstantType)
			err := initInstant()
			Ω(err).Should(BeNil())
			_, has := typeFactories[InstantType]
			Ω(has).Should(BeTrue())

			err = initInstant()
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidFactory))
		})

		It("should create elements from the factory", func() {
			v := time.Date(2017, 12, 28, 22, 20, 30, 450, time.UTC)

			elem, err := typeFactories[InstantType](v)
			Ω(err).Should(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(InstantType))
			Ω(elem.Value()).Should(BeEquivalentTo(v))
		})

		It("should not create elements from the factory if the input is not a the right type", func() {
			v := "foo"

			elem, err := typeFactories[InstantType](v)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidInput))
			Ω(elem).Should(BeNil())
		})
	})

	Context("with the default marshaller", func() {

		testValue := time.Date(2017, 12, 28, 22, 20, 30, 450, time.UTC)

		It("should create an instant value with no error", func() {
			elem, err := NewInstantElement(testValue)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(InstantType))
			Ω(elem.Value()).Should(BeEquivalentTo(testValue))
		})

		It("should serialize the instant without an issue", func() {
			elem, err := NewInstantElement(testValue)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo("#inst 2017-12-28T22:20:30Z"))
		})
	})
})
