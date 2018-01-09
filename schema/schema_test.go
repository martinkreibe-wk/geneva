package schema

import (
	"github.com/martinkreibe-wk/geneva/elements"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Schema building", func() {
	Context("with the default marshaller", func() {
		It("should create an empty schema", func() {
			schema, err := NewSchema("my/foo")
			Ω(err).Should(BeNil())

			var edn string
			edn, err = schema.Serialize()
			Ω(err).Should(BeNil())

			Ω(edn).Should(BeEquivalentTo("(def my/foo [])"))
		})

		It("should create a schema with one attribute", func() {

			schemaName := "my/foo"
			attrName := "test"
			attrType := elements.StringType
			attrCard := OneCardinality

			schema, err := NewSchema(schemaName)
			Ω(err).Should(BeNil())

			var attr Attribute
			attr, err = schema.AddAttribute(attrName, attrType, attrCard)
			Ω(err).Should(BeNil())
			Ω(attr).ShouldNot(BeNil())
			Ω(attr.Name()).Should(BeEquivalentTo(attrName))
			Ω(attr.Type()).Should(BeEquivalentTo(attrType))
			Ω(attr.Cardinality()).Should(BeEquivalentTo(attrCard))

			var edn string
			edn, err = schema.Serialize()
			Ω(err).Should(BeNil())

			Ω(edn).Should(HavePrefix("(def " + schemaName + " ["))
			Ω(edn).Should(HaveSuffix("])"))
			Ω(edn).Should(ContainSubstring(":db.install/_attribute :db.part/db"))
			Ω(edn).Should(ContainSubstring(":db/id #db/id [:db.part/db]"))
			Ω(edn).Should(ContainSubstring(":db/ident :" + attrName))
			Ω(edn).Should(ContainSubstring(":db/valueType " + string(attrType)))
			Ω(edn).Should(ContainSubstring(":db/cardinality " + string(attrCard)))
			Ω(edn).ShouldNot(ContainSubstring(":db/doc"))
		})

		It("should create a schema with an attribute with documentation", func() {

			schemaName := "my/foo"
			attrName := "test"
			attrType := elements.StringType
			attrCard := OneCardinality
			attrDoc := "a description"

			schema, err := NewSchema(schemaName)
			Ω(err).Should(BeNil())

			var attr Attribute
			attr, err = schema.AddAttribute(attrName, attrType, attrCard, attrDoc)
			Ω(err).Should(BeNil())
			Ω(attr).ShouldNot(BeNil())
			Ω(attr.Name()).Should(BeEquivalentTo(attrName))
			Ω(attr.Type()).Should(BeEquivalentTo(attrType))
			Ω(attr.Cardinality()).Should(BeEquivalentTo(attrCard))
			Ω(attr.Document()).Should(BeEquivalentTo(attrDoc))

			var edn string
			edn, err = schema.Serialize()
			Ω(err).Should(BeNil())

			Ω(edn).Should(HavePrefix("(def " + schemaName + " ["))
			Ω(edn).Should(HaveSuffix("])"))
			Ω(edn).Should(ContainSubstring(":db.install/_attribute :db.part/db"))
			Ω(edn).Should(ContainSubstring(":db/id #db/id [:db.part/db]"))
			Ω(edn).Should(ContainSubstring(":db/ident :" + attrName))
			Ω(edn).Should(ContainSubstring(":db/ident :" + attrName))
			Ω(edn).Should(ContainSubstring(":db/valueType " + string(attrType)))
			Ω(edn).Should(ContainSubstring(":db/cardinality " + string(attrCard)))
			Ω(edn).Should(ContainSubstring(":db/doc \"" + attrDoc + "\""))
		})

		It("should not create a schema with an attribute with multiple documentations", func() {

			schemaName := "my/foo"
			attrName := "test"
			attrType := elements.StringType
			attrCard := OneCardinality
			attrDoc := "a description"
			attrDoc2 := "and again a description"

			schema, err := NewSchema(schemaName)
			Ω(err).Should(BeNil())

			var attr Attribute
			attr, err = schema.AddAttribute(attrName, attrType, attrCard, attrDoc, attrDoc2)
			Ω(err).ShouldNot(BeNil())
			Ω(attr).Should(BeNil())
			Ω(err).Should(BeEquivalentTo(elements.ErrInvalidInput))
		})
	})
})
