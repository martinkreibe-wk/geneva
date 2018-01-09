package schema

const (
	ReservedDbNamespace = ":db"
)

/*
// Database defines an immutable, point-in-time database value.
type Database interface {

	// AsOf will return the value of the database filtered to include data up to t, inclusive.
	AsOf(t TransactionId) (Database, error)

	// AsOfTransaction returns the latest transaction this database image performed.
	AsOfTransaction() (TransactionId)

	// Attribute returns information about an attribute.
	Attribute(attr AttributeIdent) (Attribute, error)

	// BasisTransaction returns value of the most recent transaction in this db.
	BasisTransaction() (TransactionId)

	// NextTransaction is the next t value that will be assigned by this database.
	NextTransaction()

	// Entity returns a lazy, dynamic associative view of datoms sharing an entity id.
	Entity(EntityId) Entity

	// Filter returns a value of the database containing only Datoms satisfying the predicate.
	Filter(predicate DatomPredicate) Database

	// IsFiltered returns if the database has a filter set
	IsFiltered() bool

	// History returns a history database value containing all assertions and retractions across time.
	History() Database

	// IsHistory returns true for databases created with history()
	IsHistory() bool

	// Id Opaque, globally unique database id
	Id() string

	// Ident returns the symbolic keyword associated with an id, or the key itself if passed.
	Ident(idOrKey interface{}) AttributeIdent

	// Since returns the value of the database filtered to include only data since t, exclusive
	Since(t TransactionId) (Database, error)

	// SinceTransaction returns the transaction id defined in the since call
	SinceTransaction() TransactionId


}


type DatomPredicate func(db Database, val interface{}) (bool, error)

java.lang.Iterable<Datom>	datoms(java.lang.Object index, java.lang.Object... components)
Implements the Datoms API for raw access to matching index data.

java.lang.Object	entid(java.lang.Object entityId)
Returns the entity id associated with any kind of entity identifier.

java.lang.Object	entidAt(java.lang.Object partition, java.lang.Object timePoint)
Returns a fabricated entity id in the supplied partition whose T component is at or after the supplied t

java.lang.Iterable<Datom>	indexRange(java.lang.Object attrid, java.lang.Object start, java.lang.Object end)
Returns a range of AVET-indexed datoms.

java.lang.Object	invoke(java.lang.Object entityId, java.lang.Object... args)
Look up the database function of the entity at entityId, and invoke the function with args.




java.util.Map	pull(java.lang.Object pattern, java.lang.Object entityId)
Returns a hierarchical selection of attributes for entityId.

java.util.List<java.util.Map>	pullMany(java.lang.Object pattern, java.util.List entityIds)
Returns hierarchical selections of attributes for entityIds.

java.lang.Iterable<Datom>	seekDatoms(java.lang.Object index, java.lang.Object... components)
Raw access to index data, starting at nearest match to input


java.util.Map	with(java.util.List txData)
Returns a database with txData applied locally in memory.

*/
