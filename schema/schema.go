package schema

import (
	"github.com/martinkreibe-wk/geneva/elements"
)

const (
	DbPartition = "db.part/db"
)

// Schema
type Schema interface {
	AddAttribute(name string, elementType elements.ElementType, cardinality AttributeCardinality, doc ...string) (Attribute, error)

	// AddAttributeWithId(id int, name string, elementType elements.ElementType, cardinality AttributeCardinality, doc ...string) (error)

	Serialize() (string, error)
}

// schemaImpl implements the schema
type schemaImpl struct {

	// name of the schema.
	name string

	// attributes holds the collection of attributes this schema defines.
	attributes []Attribute
}

// Serialize the element into a string or return the appropriate error.
func (schema *schemaImpl) Serialize() (composition string, err error) {
	var n elements.SymbolElement
	if n, err = elements.NewSymbolElement(schema.name); err == nil {
		var def elements.SymbolElement
		if def, err = elements.NewSymbolElement("def"); err == nil {
			var attrs elements.CollectionElement
			if attrs, err = elements.NewVector(); err == nil {
				var wrapper elements.CollectionElement
				if wrapper, err = elements.NewGroup(def, n, attrs); err == nil {

					for _, attr := range schema.attributes {

						var coll elements.CollectionElement
						if coll, err = attr.asCollection(); err == nil {

							var kw elements.Element
							if kw, err = elements.NewKeywordElement("db.install/_attribute"); err == nil {

								var val elements.Element
								if val, err = elements.NewKeywordElement(DbPartition); err == nil {
									err = coll.Append(kw, val)
								}
							}

							attrs.Append(coll)
						}

						if err != nil {
							break
						}
					}

					if err == nil {
						composition, err = wrapper.Serialize()
					}
				}
			}
		}
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

// NewSchema will create a new schema
func NewSchema(name string) (schema Schema, err error) {

	// TODO: there may be restrictions on the names.
	schema = &schemaImpl{
		name: name,
	}

	return schema, err
}
