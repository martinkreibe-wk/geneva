package elements

import (
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
	})
})
