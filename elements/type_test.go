package elements

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Types in EDN", func() {
	Context("with the default usage", func() {
		It("initialization should not panic the first iteration, but should on the second.", func() {

			for key := range typeFactories {
				delete(typeFactories, key)
			}
			Ω(initAll).ShouldNot(Panic())
			Ω(initAll).Should(Panic())
		})

		typeCollectionMap := map[ElementType]struct {
			value bool
			name  string
		}{
			NilType:       {false, ":db.type/nil"},
			BooleanType:   {false, ":db.type/boolean"},
			StringType:    {false, ":db.type/string"},
			CharacterType: {false, ":db.type/character"},
			SymbolType:    {false, ":db.type/symbol"},
			KeywordType:   {false, ":db.type/keyword"},
			IntegerType:   {false, ":db.type/long"},
			FloatType:     {false, ":db.type/float"},
			InstantType:   {false, ":db.type/instant"},
			UUIDType:      {false, ":db.type/uuid"},
			GroupingType:  {true, ":db.type/group"},
			VectorType:    {true, ":db.type/vector"},
			MapType:       {true, ":db.type/map"},
			SetType:       {true, ":db.type/set"},
		}

		It("distinguish collections from non collections", func() {

			for t, data := range typeCollectionMap {
				Ω(t.IsCollection()).Should(BeEquivalentTo(data.value), fmt.Sprintf("Expected %s to be: %T", data.name, data.value))
			}
		})

		It("should have the right name", func() {
			for t, data := range typeCollectionMap {
				Ω(t.Name()).Should(BeEquivalentTo(data.name), fmt.Sprintf("Expected %s to be: %T", data.name, data.value))
			}
		})

		It("should have the right name for unknown types", func() {
			Ω(UnknownType.Name()).Should(BeEquivalentTo(""))

			testType := ElementType("foo")
			Ω(testType.Name()).Should(BeEquivalentTo("foo"))
		})
	})
})
