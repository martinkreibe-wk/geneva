package schema

// ConstantPattern defines a pattern that consists of a constant
type ConstantPattern string

const (

	// WildcardPattern is the specification '*' that pulls all attributes of an entity, and recursively pulls any
	// component attributes.
	WildcardPattern = ConstantPattern("*")
)

// Pattern from the constant pattern
func (pattern ConstantPattern) Compile() (string, error) {
	return string(pattern), nil
}
