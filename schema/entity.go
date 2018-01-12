package schema

// Navigation to other entities is recursive, so with thoroughly connected data such a social network graph it may be
// possible to navigate an entire dataset starting with a single entity.

// Entity provides a lazy, associative view of all the information that can be reached from an entity id. The Entity
// interface is completely generic, and can navigate any and all other data.
//
// The Entity interface provides associative access to
//   all the attribute/value pairs associated with an entity id E
//   all other entities reachable as values V from E
//   all other entities that can reach E through their values V
type Entity interface {

	// attributes
	// value(s)
	// other entities
}
