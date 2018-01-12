package schema

const (
	ReservedDbNamespace = ":db"
)

type Pattern interface {
	Compile() (string, error)
}

type Result interface {
}

// Database defines an immutable, point-in-time database value.
type Database interface {
	Pull(pattern Pattern, entity Entity) (Result, error)
}

// OptionPattern
// MapPattern
// ExpressionPattern
