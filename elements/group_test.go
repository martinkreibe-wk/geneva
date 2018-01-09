package elements_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/martinkreibe-wk/geneva/elements"
)

var _ = Describe("Grouping in EDN", func() {
	Context("with the default marshaller", func() {
		It("should create an empty group with no error", func() {
			group, err := NewGroup()
			Ω(err).Should(BeNil())
			Ω(group).ShouldNot(BeNil())
			Ω(group.ElementType()).Should(BeEquivalentTo(GroupingType))
			Ω(group.Len()).Should(BeEquivalentTo(0))
		})

		It("should serialize an empty grouping correctly", func() {
			group, err := NewGroup()
			Ω(err).Should(BeNil())

			var edn string
			edn, err = group.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo("()"))
		})

		It("should error with a nil item", func() {
			group, err := NewGroup(nil)
			Ω(err).Should(BeEquivalentTo(ErrInvalidElement))
			Ω(group).Should(BeNil())
		})

		It("should create a group element with the initial values", func() {
			elem, err := NewStringElement("foo")
			Ω(err).Should(BeNil())

			group, err := NewGroup(elem)
			Ω(err).Should(BeNil())
			Ω(group).ShouldNot(BeNil())
			Ω(group.ElementType()).Should(BeEquivalentTo(GroupingType))
			Ω(group.Len()).Should(BeEquivalentTo(1))
		})

		It("should serialize a single nil entry in a grouping correctly", func() {
			elem, err := NewNilElement()
			Ω(err).Should(BeNil())

			group, err := NewGroup(elem)
			Ω(err).Should(BeNil())

			var edn string
			edn, err = group.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo("(nil)"))
		})

		It("should serialize some nil entries in a grouping correctly", func() {

			var err error
			var elem1, elem2, elem3 Element
			var group CollectionElement

			elem1, err = NewStringElement("foo")
			Ω(err).Should(BeNil())

			elem2, err = NewStringElement("bar")
			Ω(err).Should(BeNil())

			elem3, err = NewStringElement("faz")
			Ω(err).Should(BeNil())

			group, err = NewGroup(elem1, elem2, elem3)
			Ω(err).Should(BeNil())

			var edn string
			edn, err = group.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo("(\"foo\" \"bar\" \"faz\")"))

			breakCount := 2
			templateError := Error("This is the expected error")
			err = group.IterateChildren(func(key, value Element) (e error) {
				if breakCount--; breakCount == 0 {
					e = templateError
				}
				return e
			})

			Ω(err).Should(BeEquivalentTo(templateError))
		})
	})
})
