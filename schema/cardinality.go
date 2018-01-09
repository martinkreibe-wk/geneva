package schema

import "github.com/martinkreibe-wk/geneva/elements"

// AttributeCardinality specifies whether an attribute associates a single value or a set of values with an entity.
type AttributeCardinality string

const (
	cardinalityNamespace = "cardinality"

	// CardinalityAttribute specifies whether an attribute associates a single value or a set of values with an entity.
	// The values allowed for :db/cardinality are:
	CardinalityAttribute = ReservedDbNamespace + elements.SymbolSeparator + cardinalityNamespace

	// OneCardinality describes the attribute as single valued, it associates a single value with an entity.
	OneCardinality = AttributeCardinality(ReservedDbNamespace + elements.NamespaceSeparator + cardinalityNamespace + elements.SymbolSeparator + "one")

	// OneCardinality describes the attribute as multi valued, it associates a set of values with an entity.
	ManyCardinality = AttributeCardinality(ReservedDbNamespace + elements.NamespaceSeparator + cardinalityNamespace + elements.SymbolSeparator + "many")

	// ErrUnknownCardinality describes an error where the cardinality was not allowed for the :db/cardinality attribute.
	ErrUnknownCardinality = elements.Error("Encountered an unknown cardinality")
)

// checkCardinality makes sure the cardinality is appropriate.
func checkCardinality(cardinality AttributeCardinality) (err error) {
	switch cardinality {

	// The values allowed for :db/cardinality are:
	case OneCardinality, ManyCardinality:

		// All others are an error.
	default:
		err = ErrUnknownCardinality
	}
	return err
}
