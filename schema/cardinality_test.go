package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Schema cardinality", func() {
	Context("with the default marshaller", func() {
		It("should know which cardinality values are valid", func() {
			var err error

			err = checkCardinality(OneCardinality)
			Ω(err).Should(BeNil())

			err = checkCardinality(ManyCardinality)
			Ω(err).Should(BeNil())

			err = checkCardinality(AttributeCardinality("someMadeUpStuff"))
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrUnknownCardinality))

		})
	})
})
