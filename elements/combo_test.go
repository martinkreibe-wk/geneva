package elements

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Elements in EDN", func() {
	Context("with the default marshaller", func() {
		It("should handle complex embedding", func() {
			part, err := NewKeywordElement("db.part/db")
			Ω(err).Should(BeNil())

			var id CollectionElement
			id, err = NewVector(part)
			Ω(err).Should(BeNil())

			err = id.SetTag("db/id")
			Ω(err).Should(BeNil())

			var str string
			str, err = id.Serialize()
			Ω(err).Should(BeNil())

			Ω(str).Should(BeEquivalentTo("#db/id [:db.part/db]"))

			key, err := NewKeywordElement("db/id")
			Ω(err).Should(BeNil())

			var pair Pair
			pair, err = NewPair(key, id)

			var attr CollectionElement
			attr, err = NewMap(pair)
			Ω(err).Should(BeNil())

			str, err = attr.Serialize()
			Ω(err).Should(BeNil())

			Ω(str).Should(HavePrefix("{"))
			Ω(str).Should(HaveSuffix("}"))

			Ω(str).Should(ContainSubstring(":db/id #db/id [:db.part/db]"))
		})
	})
})
