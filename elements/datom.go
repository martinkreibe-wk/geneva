package elements

const (

	// DatomTag defines the tag for the datom structure.
	DatomTag = "datom"
)

// Datom is an atomic fact in a database, composed of entity/attribute/value/transaction/added. Pronounced like "datum",
// but pluralizes as datoms. Example: #datom[17592186045422 63 "The Goonies" 13194139534317 true]
type Datom interface {

	// Serializer mixin
	Serializer

	// EntityId for this datom.
	EntityId() int64

	// AttributeId for this datom.
	AttributeId() int64

	// Value for this datom.
	Value() interface{}

	// Transaction
	Transaction() T

	// Added will indicate if this datom was added or retracted.
	Added() bool
}

// T is a timestamp in Eva
type T int64

// datomImpl implements the datom as defined.
type datomImpl struct {

	// entityId holds the entity id.
	entityId int64

	// attributeId holds the attribute id.
	attributeId int64

	// value holds the datom value.
	value interface{}

	// transaction holds the datom transaction stamp.
	transaction T

	// added to the db?
	added bool
}

// NewDatom will create a new datom.
func NewDatom(entityId int64, attributeId int64, value interface{}, transaction T, added bool) (datom Datom, err error) {
	datom = &datomImpl{
		entityId:    entityId,
		attributeId: attributeId,
		value:       value,
		transaction: transaction,
		added:       added,
	}

	return datom, err
}

// Serialize the element into a string or return the appropriate error.
func (datom *datomImpl) Serialize() (composition string, err error) {

	var internal CollectionElement

	if internal, err = NewVector(); err == nil {
		if err = internal.SetTag(DatomTag); err == nil {
			var eid, aid, v, t, a Element

			eid, err = NewIntegerElement(datom.entityId)
			if err == nil {
				aid, err = NewIntegerElement(datom.attributeId)
			}
			if err == nil {
				v, err = NewElement(datom.value)
			}
			if err == nil {
				t, err = NewIntegerElement(int64(datom.transaction))
			}
			if err == nil {
				a, err = NewBooleanElement(datom.added)
			}

			if err == nil {
				if err = internal.Append(eid, aid, v, t, a); err == nil {
					composition, err = internal.Serialize()
				}
			}
		}
	}

	return composition, err
}

// EntityId for this datom.
func (datom *datomImpl) EntityId() int64 {
	return datom.entityId
}

// AttributeId for this datom.
func (datom *datomImpl) AttributeId() int64 {
	return datom.attributeId
}

// Value for this datom.
func (datom *datomImpl) Value() interface{} {
	return datom.value
}

// Transaction for this datom.
func (datom *datomImpl) Transaction() T {
	return datom.transaction
}

// Added will indicate if this datom was added or retracted.
func (datom *datomImpl) Added() bool {
	return datom.added
}
