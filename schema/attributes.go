package schema

import (
	"github.com/martinkreibe-wk/geneva/elements"
)

const (

	// IdTag is the database id.
	IdTag = "db/id"

	// UnsetId defines an id that is not set.
	UnsetId = 0
)

// Each database has a schema that describes the set of attributes that can be associated with entities. A schema only
// defines the characteristics of the attributes themselves. It does not define which attributes can be associated with
// which entities. Decisions about which attributes apply to which entities are made by an application.

// Schema attributes are defined using the same data model used for application data. That is, attributes are themselves
// entities with associated attributes. Datomic defines a set of built-in system attributes that are used to define new
// attributes.

// Attribute is a dimension on the entity. ie: Something that can be said about an entity. An attribute has a name,
// e.g. :firstName, and a value type, e.g. :db.type/long, and a cardinality. The actual value of the attribute may have
// a tag which stereotypes the value.
type Attribute interface {
	elements.Serializer
	elements.CollectionBuilder

	// Ident defines the attributes identity
	Id() int64

	// Name for this attribute
	Name() string

	// Type this attribute supports
	Type() elements.ElementType

	// Cardinality of the attribute
	Cardinality() AttributeCardinality

	// Document associated with this attribute.
	Document() string
}

// AttributeIdent is the name of the attribute. `:db/ident` specifies the unique name of an attribute. It's value is a
// namespaced keyword with the lexical form :<namespace>/<name>. It is possible to define a name without a namespace,
// as in :<name>, but a namespace is preferred in order to avoid naming collisions. Namespaces can be hierarchical,
// with segments separated by ".", as in :<namespace>.<nested-namespace>/<name>.
//
// IdentAttribute specifies the unique name of an attribute. It's value is a namespaced keyword with the lexical
// form :<namespace>/<name>. It is possible to define a name without a namespace, as in :<name>, but a namespace is
// preferred in order to avoid naming collisions. Namespaces can be hierarchical, with segments separated by ".", as
// in :<namespace>.<nested-namespace>/<name>. The :db namespace is reserved for use by Datomic itself.
// IdentAttribute = AttributeIdent(ReservedDbNamespace + NamespaceFieldSeparator + "ident")

// attrImpl defines the actual implementation for the attribute in eva.
type attrImpl struct {
	AttrId          int64
	AttrName        string
	AttrType        elements.ElementType
	AttrCardinality AttributeCardinality
	AttrDocument    string
}

// NewAttribute creates a new attribute pair.
func NewAttribute(name string, elementType elements.ElementType, cardinality AttributeCardinality, doc ...string) (attr Attribute, err error) {
	if len(doc) <= 1 {

		var d string
		if len(doc) == 1 {
			d = doc[0]
		}

		attr = &attrImpl{
			AttrId:          UnsetId,
			AttrName:        name,
			AttrType:        elementType,
			AttrCardinality: cardinality,
			AttrDocument:    d,
		}
	} else {
		err = elements.ErrInvalidInput
	}

	return attr, err
}

// Serialize the element into a string or return the appropriate error.
func (attr *attrImpl) Serialize() (composition string, err error) {

	var elem elements.CollectionElement
	if elem, err = attr.BuildCollection(); err == nil {
		composition, err = elem.Serialize()
	}

	return composition, err
}

// Ident for this attribute
func (attr *attrImpl) Id() (ident int64) {
	return attr.AttrId
}

// Name for this attribute
func (attr *attrImpl) Name() (name string) {
	return attr.AttrName
}

// Type this attribute supports
func (attr *attrImpl) Type() (elemType elements.ElementType) {
	return attr.AttrType
}

// Cardinality of the attribute
func (attr *attrImpl) Cardinality() (card AttributeCardinality) {
	return attr.AttrCardinality
}

// Document of the attribute
func (attr *attrImpl) Document() (doc string) {
	return attr.AttrDocument
}

// appendPair will create a pair and push it onto the pair collection
func appendPair(keySymbol string, value elements.Element, pairs *elements.Pairs, err error) error {
	if err == nil {
		var key elements.SymbolElement
		if key, err = elements.NewKeywordElement(keySymbol); err == nil {
			err = pairs.Append(key, value)
		}
	}

	return err
}

// keywordCreator creates the keyword from a single string.
func keywordCreator(in string) (elements.Element, error) {
	return elements.NewKeywordElement(in)
}

// appendPairWithString will create a string pair.
func appendPairWithString(keySymbol string, creator func(string) (elements.Element, error), value string, pairs *elements.Pairs, err error) error {
	if err == nil {
		var val elements.Element
		if val, err = creator(value); err == nil {
			err = appendPair(keySymbol, val, pairs, err)
		}
	}
	return err
}

// BuildCollection will create an element collection.
func (attr *attrImpl) BuildCollection() (elem elements.CollectionElement, err error) {

	pairs := &elements.Pairs{}

	var part elements.SymbolElement
	var id elements.Element

	if attr.AttrId == UnsetId {
		if part, err = elements.NewKeywordElement(DbPartition); err == nil {
			id, err = elements.NewVector(part)
		}
	} else {
		id, err = elements.NewIntegerElement(attr.AttrId)
	}

	if err == nil {
		err = id.SetTag(IdTag)
	}
	err = appendPair(IdAttribute, id, pairs, err)
	err = appendPairWithString(IdentAttribute, keywordCreator, attr.AttrName, pairs, err)
	err = appendPairWithString(ValueTypeAttribute, keywordCreator, string(attr.AttrType), pairs, err)
	err = appendPairWithString(CardinalityAttribute, keywordCreator, string(attr.AttrCardinality), pairs, err)
	if len(attr.AttrDocument) > 0 {
		err = appendPairWithString(DocAttribute, elements.NewStringElement, attr.AttrDocument, pairs, err)
	}

	if err == nil {
		elem, err = elements.NewMap(pairs.Raw()...)
	}

	return elem, err
}
