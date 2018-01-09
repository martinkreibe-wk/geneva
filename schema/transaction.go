package schema

// TransactionId is a point in time in a database. Every transaction is assigned a numeric t value greater than any
// previous t in the database, and all processes see a consistent succession of ts. A transaction value t can be
// converted to a tx id with Peer.toTx.
type TransactionId int64
