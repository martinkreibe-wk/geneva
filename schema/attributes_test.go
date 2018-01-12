package schema

import (
	"fmt"

	"github.com/martinkreibe-wk/geneva/elements"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Attributes", func() {
	Context("with the default marshaller", func() {
		It("should not create a schema with an attribute with multiple documentations", func() {

			attrName := "test"
			attrType := elements.StringType
			attrCard := OneCardinality

			var attr Attribute
			var err error

			attr, err = NewAttribute(attrName, attrType, attrCard)
			Ω(err).Should(BeNil())
			Ω(attr).ShouldNot(BeNil())

			Ω(attr.Name()).Should(BeEquivalentTo(attrName))
			Ω(attr.Type()).Should(BeEquivalentTo(attrType))
			Ω(attr.Cardinality()).Should(BeEquivalentTo(attrCard))

			var edn string
			edn, err = attr.Serialize()
			Ω(edn).Should(ContainSubstring(":db/id #db/id [:db.part/db]"))
			Ω(edn).Should(ContainSubstring(":db/ident :" + attrName))
			Ω(edn).Should(ContainSubstring(":db/valueType " + string(attrType)))
			Ω(edn).Should(ContainSubstring(":db/cardinality " + string(attrCard)))
		})

		It("should not create a schema with an attribute with multiple documentations", func() {

			attrName := "test"
			attrType := elements.StringType
			attrCard := OneCardinality
			id := int64(123)

			var attr Attribute
			var err error

			attr, err = NewAttribute(attrName, attrType, attrCard)
			Ω(err).Should(BeNil())
			Ω(attr).ShouldNot(BeNil())

			attr.(*attrImpl).AttrId = id
			Ω(attr.Id()).Should(BeEquivalentTo(id))

			var elem elements.CollectionElement
			elem, err = attr.BuildCollection()
			Ω(err).Should(BeNil())
			Ω(elem).ShouldNot(BeNil())

			var e elements.Element
			e, err = elem.Get(IdAttribute)
			Ω(err).Should(BeNil())
			Ω(e.Value()).Should(BeEquivalentTo(id))

			var edn string
			edn, err = attr.Serialize()
			Ω(edn).Should(ContainSubstring(fmt.Sprintf(":db/id #db/id %d", id)))
			Ω(edn).Should(ContainSubstring(":db/ident :" + attrName))
			Ω(edn).Should(ContainSubstring(":db/valueType " + string(attrType)))
			Ω(edn).Should(ContainSubstring(":db/cardinality " + string(attrCard)))
		})
	})
})
