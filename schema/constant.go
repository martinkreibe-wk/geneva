package schema

type ConstantPattern string

const (
	// Wid
	WildcardPattern = ConstantPattern("*")
)

// Pattern from the constant pattern
func (pattern ConstantPattern) Compile() (string, error) {
	return string(pattern), nil
}
