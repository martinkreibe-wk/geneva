package elements

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map in EDN", func() {
	Context("with the default marshaller", func() {

		makePair := func(key string, value string) Pair {
			elem, err := NewStringElement(value)
			Ω(err).Should(BeNil())

			var pair Pair
			pair, err = NewPair(key, elem)
			Ω(err).Should(BeNil())
			Ω(pair.Key()).ShouldNot(BeNil())
			Ω(pair.Key().ElementType()).Should(BeEquivalentTo(KeywordType))
			Ω(pair.Key().Value().(SymbolElement).Name()).Should(BeEquivalentTo(key[1:]))

			Ω(pair.Value()).ShouldNot(BeNil())
			Ω(pair.Value().ElementType()).Should(BeEquivalentTo(StringType))
			Ω(pair.Value().Value()).Should(BeEquivalentTo(value))

			return pair
		}

		It("should create an empty map with no error", func() {
			group, err := NewMap()
			Ω(err).Should(BeNil())
			Ω(group).ShouldNot(BeNil())
			Ω(group.ElementType()).Should(BeEquivalentTo(MapType))
			Ω(group.Len()).Should(BeEquivalentTo(0))
		})

		It("should serialize an empty map correctly", func() {
			group, err := NewMap()
			Ω(err).Should(BeNil())

			var edn string
			edn, err = group.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo("{}"))
		})

		It("should error with a nil item", func() {
			group, err := NewMap(nil)
			Ω(err).Should(BeEquivalentTo(ErrInvalidPair))
			Ω(group).Should(BeNil())
		})

		It("should create a map element with the initial values", func() {
			elem, err := NewStringElement("foo")
			Ω(err).Should(BeNil())

			var pair Pair
			pair, err = NewPair(elem, elem)
			Ω(err).Should(BeNil())

			group, err := NewMap(pair)
			Ω(err).Should(BeNil())
			Ω(group).ShouldNot(BeNil())
			Ω(group.ElementType()).Should(BeEquivalentTo(MapType))
			Ω(group.Len()).Should(BeEquivalentTo(1))

			var v Element
			v, err = group.Get("foo")
			Ω(err).Should(BeNil())
			Ω(v).ShouldNot(BeNil())

			v, err = group.Get("notfound")
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrNoValue))
			Ω(v).Should(BeNil())
		})

		It("should serialize a single nil entry in a map correctly", func() {
			elem, err := NewNilElement()
			Ω(err).Should(BeNil())

			var pair Pair
			pair, err = NewPair(elem, elem)
			Ω(err).Should(BeNil())

			group, err := NewMap(pair)
			Ω(err).Should(BeNil())

			var edn string
			edn, err = group.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(BeEquivalentTo("{nil nil}"))
		})

		It("should serialize some pairs entries in a map correctly", func() {

			keys := map[string]string{
				":key1": "val1",
				":key2": "val2",
				":key3": "val3",
				":key4": "val3", // same values are ok
			}

			pairs := []Pair{}
			for k, v := range keys {
				pairs = append(pairs, makePair(k, v))
			}

			group, err := NewMap(pairs...)
			Ω(err).Should(BeNil())

			var edn string
			edn, err = group.Serialize()
			Ω(err).Should(BeNil())
			Ω(edn).Should(HavePrefix("{"))
			Ω(edn).Should(HaveSuffix("}"))

			for k, v := range keys {
				Ω(edn).Should(ContainSubstring(k + " " + "\"" + v + "\""))
			}
		})

		It("should not accept duplicate keys", func() {
			p1 := makePair(":key1", "val1")
			p2 := makePair(":key1", "val2")
			Ω(p1.Key().Equals(p2.Key())).Should(BeTrue())

			group, err := NewMap(
				p1,
				p2,
			)
			Ω(err).ShouldNot(BeNil())
			Ω(group).Should(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrDuplicateKey))
		})

		It("should break the iteration and return the error", func() {

			keys := map[string]string{
				":key1": "val1",
				":key2": "val2",
				":key3": "val3",
				":key4": "val3", // same values are ok
			}

			pairs := []Pair{}
			for k, v := range keys {
				pairs = append(pairs, makePair(k, v))
			}

			group, err := NewMap(pairs...)
			Ω(err).Should(BeNil())

			breakCount := len(keys) / 2
			Ω(len(keys) > breakCount).Should(BeTrue())

			templateError := Error("This is the expected error")
			err = group.IterateChildren(func(key, value Element) (e error) {
				if breakCount--; breakCount == 0 {
					e = templateError
				}
				return e
			})

			Ω(err).Should(BeEquivalentTo(templateError))
		})

		It("should break the iteration and return the error", func() {

			m, err := NewMap()
			Ω(err).Should(BeNil())

			var elem Element
			elem, err = NewNilElement()
			Ω(err).Should(BeNil())

			err = m.Append(elem)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidInput))
		})

		It("should error if somehow the collection was not the type we were expecting.", func() {

			m, err := NewMap()
			Ω(err).Should(BeNil())

			raw := m.(*collectionElemImpl)
			raw.collection = &struct{}{} // overwrite the actual data.

			var elem Element
			elem, err = NewNilElement()
			Ω(err).Should(BeNil())

			err = m.Append(elem)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidElement))

			_, err = m.Get("foo")
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidElement))
		})

		It("should break the creation if there is an error", func() {

			p, err := NewPair(":key1", nil)
			Ω(err).ShouldNot(BeNil())

			_, err = NewMap(p)
			Ω(err).ShouldNot(BeNil())
			Ω(err).Should(BeEquivalentTo(ErrInvalidPair))
		})
	})
})
