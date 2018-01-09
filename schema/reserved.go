package schema

import "github.com/martinkreibe-wk/geneva/elements"

const (

	// IdentAttribute specifies the unique name of an attribute. It's value is a namespaced keyword with the lexical
	// form :<namespace>/<name>. It is possible to define a name without a namespace, as in :<name>, but a namespace is
	// preferred in order to avoid naming collisions. Namespaces can be hierarchical, with segments separated by ".",
	// as in :<namespace>.<nested-namespace>/<name>. The :db namespace is reserved for use by Datomic itself.
	IdentAttribute = ReservedDbNamespace + elements.SymbolSeparator + "ident"

	// specifies the type of value that can be associated with an attribute. The type is expressed as a keyword.
	// Allowable values are listed below.
	ValueTypeAttribute = ReservedDbNamespace + elements.SymbolSeparator + "valueType"

	// IdAttribute defines the element id.
	IdAttribute = ReservedDbNamespace + elements.SymbolSeparator + "id"

	// DocAttribute specifies a documentation string.
	DocAttribute = ReservedDbNamespace + elements.SymbolSeparator + "doc"

	// IndexAttribute specifies a boolean value indicating that an index should be generated for this attribute.
	// Defaults to false.
	IndexAttribute = ReservedDbNamespace + elements.SymbolSeparator + "index"

	// FulltextAttribute specifies a boolean value indicating that an eventually consistent fulltext search index should
	// be generated for the attribute. Defaults to false.
	// Fulltext search is constrained by several defaults (which cannot be altered): searches are case insensitive,
	// remove apostrophe or apostrophe and s sequences, and filter out the following common English stop words:
	//
	// "a", "an", "and", "are", "as", "at", "be", "but", "by", "for", "if", "in", "into", "is", "it", "no", "not", "of",
	// "on", "or", "such", "that", "the", "their", "then", "there", "these", "they", "this", "to", "was", "will", "with"
	FulltextAttribute = ReservedDbNamespace + elements.SymbolSeparator + "fulltext"

	// IsComponentAttribute specifies a boolean value indicating that an attribute whose type is :db.type/ref refers to
	// a subcomponent of the entity to which the attribute is applied. When you retract an entity with
	// :db.fn/retractEntity, all subcomponents are also retracted. When you touch an entity, all its subcomponent
	// entities are touched recursively. Defaults to false.
	// The purpose of :db/noHistory is to conserve storage, not to make semantic guarantees about removing information.
	// The effect of :db/noHistory happens in the background, and some amount of history may be visible even for
	// attributes with :db/noHistory set to true
	IsComponentAttribute = ReservedDbNamespace + elements.SymbolSeparator + "isComponent"

	// NoHistoryAttribute specifies a boolean value indicating whether past values of an attribute should not be
	// retained. Defaults to false.
	NoHistoryAttribute = ReservedDbNamespace + elements.SymbolSeparator + "noHistory"

	// TODO: Its own thing.
	// TxInstantAttribute is the id given to a transaction. You can set :db/txInstant explicitly, overriding the
	// transactor's clock time. When you do, you must choose a :db/txInstant value that is not older than any existing
	// transaction, and not newer than the transactor's clock time. This capability enables initial imports of existing
	// data that has its own timestamps.
	TxInstantAttribute = ReservedDbNamespace + elements.SymbolSeparator + "txInstant"
)

// reservedAttributes defines the collection reserved attributes, the boolean says if it is settable or not.
var reservedAttributes = map[string]bool{
	IdAttribute:          false,
	IdentAttribute:       false,
	ValueTypeAttribute:   false,
	CardinalityAttribute: false, // technically this is, but we aren't supporting it.
	TxInstantAttribute:   true,
	DocAttribute:         true,
	IndexAttribute:       true,
	FulltextAttribute:    true,
	IsComponentAttribute: true,
	NoHistoryAttribute:   true,
	UniqueAttribute:      true,
}
