package schema

// The Attributes With Options specification provides control of various aspects of the values returned by the pull of an attribute.
type OptionsPattern interface {
	Pattern

	// The following pattern uses an :as option to pull an :artist/name, replacing the key in the result map with the string "Band Name", pulling on ledZeppelin:
	// The :as option allows replacement of the key for an attribute result map with an arbitrary value.
	As() (interface{}, bool)

	SetAs(interface{}) error

	// The :limit option controls how many values will be returned for a cardinality-many attribute. A limit can be a positive number or nil. A nil limit causes all values to be returned, and should be used with caution.
	// In the absence of an explicit limit, pull will return the first 1000 values for a cardinality-many attribute.
	Limit() (int, bool)

	SetLimit(limit ...int) error

	// The :default option specifies a value to use if an attribute is not present for an entity.
	Default() (interface{}, bool)

	SetDefault(interface{}) error
}
