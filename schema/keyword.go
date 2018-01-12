package schema

import "github.com/martinkreibe-wk/geneva/elements"

// KeywordPattern is an attribute spec names an attribute, with an optional leading underscore on the local name to
// reverse the direction of navigation.
type KeywordPattern interface {
	Pattern

	// SetKeyword will set the keyword
	SetKeyword(element elements.KeywordElement) (err error)

	// Keyword of this pattern
	Keyword() (element elements.KeywordElement)
}

// keywordPatternImpl defines the keyword pattern
type keywordPatternImpl struct {
	elements.KeywordElement
}

// SetKeyword will apply the keyword as its reference
func (pattern *keywordPatternImpl) SetKeyword(element elements.KeywordElement) (err error) {
	pattern.KeywordElement = element
	return nil
}

// Keyword returns the internal keyword
func (pattern *keywordPatternImpl) Keyword() (element elements.KeywordElement) {
	return pattern.KeywordElement
}

// Pattern representation of the keyword pattern.
func (pattern *keywordPatternImpl) Pattern() (string, error) {
	return pattern.KeywordElement.Serialize()
}
