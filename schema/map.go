package schema

// MapPatterns can explicitly specify the handling of referenced entities by using a map instead of just an attribute
// name. The simplest map specification is a map specifying a specific pattern for a particular attr-name.
//
// map-spec           = { ((attr-name | limit-expr) (pattern | recursion-limit))+ }
// limit-expr         = [("limit" | 'limit') attr-name (positive-number | nil)]
// recursion-limit    = positive-number | '...'
type MapPatterns interface {
	Pattern

	AppendAttribute(attribute Attribute, pattern Pattern) (err error)
	AppendLimit(attribute Attribute)
}
