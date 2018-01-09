package elements

// ElementType indicates the EDN element construct
type ElementType string

// ElementTypeFactory defines the factory for an element.
type ElementTypeFactory func(interface{}) (Element, error)

const (

	// typeNamespace hold the type namespace
	typeNamespace = KeywordPrefix + "db.type"

	// UnknownType TODO
	UnknownType = ElementType("")

	// bytes, uri, ref
	// KeywordType is the value type for keywords. Keywords are used as names, and are interned for efficiency.
	// Keywords map to the native interned-name type in languages that support them.
	NilType = ElementType(typeNamespace + SymbolSeparator + "nil")

	// CharacterType TODO
	CharacterType = ElementType(typeNamespace + SymbolSeparator + "character")

	// KeywordType is the value type for keywords. Keywords are used as names, and are interned for efficiency.
	// Keywords map to the native interned-name type in languages that support them.
	KeywordType = ElementType(typeNamespace + SymbolSeparator + "keyword")

	// SymbolType TODO
	SymbolType = ElementType(typeNamespace + SymbolSeparator + "symbol")

	// StringType is the value type for strings.
	StringType = ElementType(typeNamespace + SymbolSeparator + "string")

	// BooleanType value type.
	BooleanType = ElementType(typeNamespace + SymbolSeparator + "boolean")

	// IntegerType is the fixed integer value type. Same semantics as a Java long: 64 bits wide, two's complement binary
	// representation.
	IntegerType = ElementType(typeNamespace + SymbolSeparator + "long")

	// BigIntType is the value type for arbitrary precision integers. Maps to java.math.BigInteger on Java platforms.
	BigIntType = ElementType(typeNamespace + SymbolSeparator + "bigint")

	// FloatType is the floating point value type. Same semantics as a Java float: single-precision 32-bit IEEE 754
	// floating point.
	FloatType = ElementType(typeNamespace + SymbolSeparator + "float")

	// DoubleType is the floating point value type. Same semantics as a Java double: double-precision 64-bit IEEE 754
	// floating point.
	DoubleType = ElementType(typeNamespace + SymbolSeparator + "double")

	// BigDecType is the value type for arbitrary precision floating point numbers. Maps to java.math.BigDecimal on Java
	// platforms.
	BigDecType = ElementType(typeNamespace + SymbolSeparator + "bigdec")

	// RefType is the value type for references. All references from one entity to another are through attributes with
	// this value type.
	RefType = ElementType(typeNamespace + SymbolSeparator + "ref")

	// TODO
	GroupingType = ElementType(typeNamespace + SymbolSeparator + "group")
	VectorType   = ElementType(typeNamespace + SymbolSeparator + "vector")
	MapType      = ElementType(typeNamespace + SymbolSeparator + "map")
	SetType      = ElementType(typeNamespace + SymbolSeparator + "set")

	// InstantType is the value type for instants in time. Stored internally as a number of milliseconds since midnight,
	// January 1, 1970 UTC. Maps to java.util.Date on Java platforms.
	InstantType = ElementType(typeNamespace + SymbolSeparator + "instant")

	// UUIDType is the value type for UUIDs. Maps to java.util.UUID on Java platforms.
	UUIDType = ElementType(typeNamespace + SymbolSeparator + "uuid")

	// URIType is the value type for URIs. Maps to java.net.URI on Java platforms.
	URIType = ElementType(typeNamespace + SymbolSeparator + "uri")

	// BytesType is the value type for small binary data. Maps to byte array on Java platforms. See limitations.
	BytesType = ElementType(typeNamespace + SymbolSeparator + "bytes")

	// ErrInvalidFactory defines the factory error
	ErrInvalidFactory = Error("Invalid factory")

	// ErrInvalidInput defines the input error
	ErrInvalidInput = Error("Invalid input")
)

// typeFactories hold the collection of element factories.
var typeFactories = map[ElementType]ElementTypeFactory{}

type elementDefinition struct {
	isCollection bool
	initializer  func() error
}

// typeDefinitions holds the type to name/initializer mappings
var typeDefinitions = map[ElementType]*elementDefinition{
	UnknownType:   {false, nil},
	NilType:       {false, initNil},
	BooleanType:   {false, initBoolean},
	StringType:    {false, initString},
	CharacterType: {false, initCharacter},
	SymbolType:    {false, nil},
	KeywordType:   {false, initKeyword},
	IntegerType:   {false, initInteger},
	FloatType:     {false, initFloat},
	InstantType:   {false, initInstant},
	UUIDType:      {false, initUUID},
	GroupingType:  {true, nil},
	VectorType:    {true, nil},
	MapType:       {true, nil},
	SetType:       {true, nil},

	// TODO
	URIType:    {false, nil},
	BytesType:  {false, nil},
	BigIntType: {false, nil},
	BigDecType: {false, nil},
	DoubleType: {false, nil},
	RefType:    {false, nil},
}

// init will initialize the package - NOTE this is not testable
func init() {
	initAll()
}

// AddElementTypeFactory adds an element factory to the factory collection.
func AddElementTypeFactory(elemType ElementType, elemFactory ElementTypeFactory) (err error) {
	if _, has := typeFactories[elemType]; !has {
		typeFactories[elemType] = elemFactory
	} else {
		err = ErrInvalidFactory
	}

	return err
}

// init the package
func initAll() {
	for _, def := range typeDefinitions {
		if def.initializer != nil {
			if e := def.initializer(); e != nil {
				panic(e)
			}
		}
	}
}

// IsCollection indicates that this type is a collection
func (t ElementType) IsCollection() bool {
	var def *elementDefinition
	var found bool
	if def, found = typeDefinitions[t]; !found || def == nil {
		def = typeDefinitions[UnknownType]
	}

	return def.isCollection
}

// Name returns the name of the set.
func (t ElementType) Name() (name string) {
	return string(t)
}
