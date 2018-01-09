package schema

// An immutable, point-in-time fact: [entity, attribute, value, transaction, added]
type Datom interface {

	// EntityId returns this datom's entity id.
	EntityId() int64

	// AttributeId is datom's attribute id.
	AttributeId() int64

	// Value of the datom
	Value() interface{}

	// TransactionId is this datom's transaction id.
	TransactionId() TransactionId

	// Added identifies this datom as added or retracted?
	Added() bool
}
