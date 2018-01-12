package elements

import (
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Elements in EDN", func() {
	Context("with the default marshaller", func() {
		It("should create an base element with no error", func() {

			t := ElementType(99)

			elem, err := makeBaseElement(nil, t, func(i interface{}) (string, error) {
				return "", nil
			})
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.ElementType()).Should(BeIdenticalTo(t))
		})

		It("should create an base element with no error", func() {

			t := ElementType(99)

			elem, err := makeBaseElement(nil, t, nil)
			Ω(err).ShouldNot(BeNil())
			Ω(elem).Should(BeNil())
			Ω(err).Should(BeIdenticalTo(ErrInvalidElement))
		})

		It("should equal the same thing if they are actually equal", func() {

			value := "42"

			elem, err := makeBaseElement(value, StringType, func(value interface{}) (out string, e error) {
				out = strconv.Quote(value.(string))
				return out, e
			})
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())
			Ω(elem.Value()).Should(BeEquivalentTo(value))

			elem2, err := makeBaseElement(value, StringType, func(value interface{}) (out string, e error) {
				out = strconv.Quote(value.(string))
				return out, e
			})
			Ω(err).Should(BeNil())
			Ω(elem2).ShouldNot(BeNil())
			Ω(elem2.Value()).Should(BeEquivalentTo(value))

			Ω(elem.Equals(elem2)).Should(BeTrue())
		})

	})
})
