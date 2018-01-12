package schema

import (
	"github.com/martinkreibe-wk/geneva/elements"
)

const (
	DbPartition = "db.part/db"

	// InstallOperation defines the installation operation
	InstallOperation = "db.install/_attribute"
)

// Schema defines an eva schema.
type Schema interface {
	elements.Serializer
	elements.CollectionBuilder

	// AddAttribute will add an attribute to the schema
	AddAttribute(name string, elementType elements.ElementType, cardinality AttributeCardinality, doc ...string) (Attribute, error)
}

// schemaImpl implements the schema
type schemaImpl struct {

	// name of the schema.
	name string

	// attributes holds the collection of attributes this schema defines.
	attributes []Attribute
}

// NewSchema will create a new schema
func NewSchema(name string) (schema Schema, err error) {

	// TODO: there may be restrictions on the names.
	schema = &schemaImpl{
		name: name,
	}

	return schema, err
}

// Serialize the element into a string or return the appropriate error.
func (schema *schemaImpl) BuildCollection() (elem elements.CollectionElement, err error) {
	var n elements.SymbolElement
	if n, err = elements.NewSymbolElement(schema.name); err == nil {
		var def elements.SymbolElement
		if def, err = elements.NewSymbolElement("def"); err == nil {
			var attrs elements.CollectionElement
			if attrs, err = elements.NewVector(); err == nil {
				if elem, err = elements.NewGroup(def, n, attrs); err == nil {

					for _, attr := range schema.attributes {

						var coll elements.CollectionElement
						if coll, err = attr.BuildCollection(); err == nil {

							var kw elements.Element
							if kw, err = elements.NewKeywordElement(InstallOperation); err == nil {

								var val elements.Element
								if val, err = elements.NewKeywordElement(DbPartition); err == nil {
									err = coll.Append(kw, val)
								}
							}

							err = attrs.Append(coll)
						}

						if err != nil {
							break
						}
					}
				}
			}
		}
	}

	return elem, err
}

// BuildCollection this object into a collection of elements.
func (schema *schemaImpl) Serialize() (composition string, err error) {

	var elem elements.Element
	if elem, err = schema.BuildCollection(); err == nil {
		composition, err = elem.Serialize()
	}

	return composition, err
}

// AddAttribute will add an attribute
func (schema *schemaImpl) AddAttribute(name string, elementType elements.ElementType, cardinality AttributeCardinality, doc ...string) (attr Attribute, err error) {

	if attr, err = NewAttribute(name, elementType, cardinality, doc...); err == nil {
		schema.attributes = append(schema.attributes, attr)
	}

	return attr, err
}
