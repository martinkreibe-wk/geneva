package elements

import (
	"github.com/mattrobenolt/gocql/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UUID in EDN", func() {
	Context("", func() {

		It("should initialize without issue", func() {
			delete(typeFactories, UUIDType)
			err := initUUID()
			Ω(err).Should(BeNil())
			_, has := typeFactories[UUIDType]
			Ω(has).Should(BeTrue())

			err = initUUID()
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidFactory))
		})

		It("should create elements from the factory", func() {
			uuidValue := "12345678-90ab-cdef-9876-0123456789ab"
			v, err := uuid.ParseUUID(uuidValue)
			Ω(err).Should(BeNil())

			elem, err := typeFactories[UUIDType](v)
			Ω(err).Should(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(UUIDType))
			Ω(elem.Value()).Should(BeEquivalentTo(v))
		})

		It("should not create elements from the factory if the input is not a the right type", func() {
			v := "foo"

			elem, err := typeFactories[UUIDType](v)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidInput))
			Ω(elem).Should(BeNil())
		})
	})

	Context("with the default marshaller", func() {

		uuidValue := "12345678-90ab-cdef-9876-0123456789ab"
		testValue, err := uuid.ParseUUID(uuidValue)
		if err != nil {
			panic(err)
		}

		It("should create an uuid value with no error", func() {

			elem, err := NewUUIDElement(testValue)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeEquivalentTo(UUIDType))
			Ω(elem.Value()).Should(BeEquivalentTo(testValue))
		})

		It("should serialize the uuid without an issue", func() {

			elem, err := NewUUIDElement(testValue)
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			edn, err := elem.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo("#uuid " + uuidValue))
		})
	})
})
