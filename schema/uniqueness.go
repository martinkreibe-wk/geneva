package schema

import "github.com/martinkreibe-wk/geneva/elements"

// Uniqueness specifies a uniqueness constraint for the values of an attribute. Setting an attribute :db/unique
// also implies :db/index. :db/unique defaults to nil.
type Uniqueness string

const (

	// uniqueNamespace
	uniqueNamespace = "unique"

	// UniqueAttribute specifies a uniqueness constraint for the values of an attribute. Setting an attribute :db/unique
	// also implies :db/index. :db/unique defaults to nil.
	UniqueAttribute = ReservedDbNamespace + elements.NamespaceSeparator + uniqueNamespace

	// UniqueValue specifies only one entity can have a given value for this attribute. Attempts to assert a duplicate
	// value for the same attribute for a different entity id will fail.
	UniqueValue = Uniqueness(ReservedDbNamespace + elements.NamespaceSeparator + uniqueNamespace + elements.SymbolSeparator + "value")

	// UniqueIdentity specifies only one entity can have a given value for this attribute and "upsert" is enabled;
	// attempts to insert a duplicate value for a temporary entity id will cause all attributes associated with that
	// temporary id to be merged with the entity already in the database.
	UniqueIdentity = Uniqueness(ReservedDbNamespace + elements.NamespaceSeparator + uniqueNamespace + elements.SymbolSeparator + "identity")
)
