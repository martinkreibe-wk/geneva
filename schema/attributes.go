package schema

import (
	"github.com/martinkreibe-wk/geneva/elements"
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

	// element returns the attribute as an element
	asCollection() (elements.CollectionElement, error)
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
	impl elements.CollectionElement
}

func appendPair(keySymbol string, value elements.Element, pairs *elements.Pairs, err error) error {
	if err == nil {
		var key elements.SymbolElement
		if key, err = elements.NewKeywordElement(keySymbol); err == nil {
			err = pairs.Append(key, value)
		}
	}

	return err
}

func keywordCreator(in string) (elements.Element, error) {
	return elements.NewKeywordElement(in)
}

func appendPairWithString(keySymbol string, creator func(string) (elements.Element, error), value string, pairs *elements.Pairs, err error) error {
	if err == nil {
		var val elements.Element
		if val, err = creator(value); err == nil {
			err = appendPair(keySymbol, val, pairs, err)
		}
	}
	return err
}

func NewAttribute(name string, elementType elements.ElementType, cardinality AttributeCardinality, doc ...string) (attr Attribute, err error) {
	if len(doc) <= 1 {
		pairs := &elements.Pairs{}

		var part elements.SymbolElement
		var id elements.CollectionElement
		if part, err = elements.NewKeywordElement(DbPartition); err == nil {
			if id, err = elements.NewVector(part); err == nil {
				err = id.SetTag("db/id")
			}
		}

		err = appendPair(IdAttribute, id, pairs, err)
		err = appendPairWithString(IdentAttribute, keywordCreator, name, pairs, err)
		err = appendPairWithString(ValueTypeAttribute, keywordCreator, string(elementType), pairs, err)
		err = appendPairWithString(CardinalityAttribute, keywordCreator, string(cardinality), pairs, err)
		if len(doc) == 1 {
			err = appendPairWithString(DocAttribute, elements.NewStringElement, doc[0], pairs, err)
		}

		if err == nil {
			var impl elements.CollectionElement
			if impl, err = elements.NewMap(pairs.Raw()...); err == nil {
				attr = &attrImpl{
					impl: impl,
				}
			}
		}
	} else {
		err = elements.ErrInvalidInput
	}

	return attr, err
}

// asCollection returns the attribute as a collection
func (attr *attrImpl) asCollection() (elements.CollectionElement, error) {
	return attr.impl, nil
}

// Ident for this attribute
func (attr *attrImpl) Id() (ident int64) {

	// TODO: We dont have a clean error handling process here... Should we return an err?
	if elem, err := attr.impl.Get(IdAttribute); err == nil {
		ident, _ = elem.Value().(int64)
	}
	return ident
}

// Name for this attribute
func (attr *attrImpl) Name() (name string) {
	// TODO: We dont have a clean error handling process here... Should we return an err?
	if elem, err := attr.impl.Get(IdentAttribute); err == nil {
		v := elem.Value()
		if raw, convert := v.(elements.SymbolElement); convert {
			name = raw.Name()
		}
	} else {
		panic(err)
	}
	return name
}

// Type this attribute supports
func (attr *attrImpl) Type() (elemType elements.ElementType) {
	// TODO: We dont have a clean error handling process here... Should we return an err?
	if elem, err := attr.impl.Get(ValueTypeAttribute); err == nil {
		v := elem.Value()
		if raw, convert := v.(elements.SymbolElement); convert {
			if v, e := raw.Serialize(); e == nil {
				elemType = elements.ElementType(v)
			}
		}
	}
	return elemType
}

// Cardinality of the attribute
func (attr *attrImpl) Cardinality() (card AttributeCardinality) {
	// TODO: We dont have a clean error handling process here... Should we return an err?
	if elem, err := attr.impl.Get(CardinalityAttribute); err == nil {
		v := elem.Value()
		if raw, convert := v.(elements.SymbolElement); convert {
			if v, e := raw.Serialize(); e == nil {
				card = AttributeCardinality(v)
			}
		}
	}
	return card
}

// Document of the attribute
func (attr *attrImpl) Document() (doc string) {
	// TODO: We dont have a clean error handling process here... Should we return an err?
	if elem, err := attr.impl.Get(DocAttribute); err == nil {
		doc, _ = elem.Value().(string)
	}
	return doc
}
