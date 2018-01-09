package elements_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/martinkreibe-wk/geneva/elements"
)

var _ = Describe("Pair for maps in EDN", func() {
	Context("with the default usage", func() {
		It("should create a pair with no error", func() {
			key, err := NewStringElement("key")
			Ω(err).Should(BeNil())
			value, err := NewStringElement("value")
			Ω(err).Should(BeNil())

			pair, err := NewPair(key, value)
			Ω(err).Should(BeNil())
			Ω(pair).ShouldNot(BeNil())
			Ω(pair.Key()).Should(BeEquivalentTo(key))
			Ω(pair.Value()).Should(BeEquivalentTo(value))
		})

		It("should create an error with nil key", func() {
			value, err := NewStringElement("value")
			Ω(err).Should(BeNil())

			pair, err := NewPair(nil, value)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidPair))
			Ω(pair).Should(BeNil())
		})

		It("should create an error with nil value", func() {
			key, err := NewStringElement("key")
			Ω(err).Should(BeNil())

			pair, err := NewPair(key, nil)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidPair))
			Ω(pair).Should(BeNil())
		})
	})
})
