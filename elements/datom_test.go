package elements

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Datoms in EDN", func() {
	Context("with the default marshaller", func() {
		It("should handle datom serialization", func() {

			var err error
			var datom Datom

			eid := int64(123)
			aid := int64(345)
			v := "this is a test"
			trans := T(567)
			added := true

			datom, err = NewDatom(eid, aid, v, trans, added)
			Ω(err).Should(BeNil())

			Ω(datom.EntityId()).Should(BeEquivalentTo(eid))
			Ω(datom.AttributeId()).Should(BeEquivalentTo(aid))
			Ω(datom.Value()).Should(BeEquivalentTo(v))
			Ω(datom.Transaction()).Should(BeEquivalentTo(trans))
			Ω(datom.Added()).Should(BeEquivalentTo(added))

			var str string
			str, err = datom.Serialize()
			Ω(str).Should(HavePrefix("#datom ["))
			Ω(str).Should(HaveSuffix("]"))

			Ω(str).Should(ContainSubstring(fmt.Sprintf("[%d %d \"%s\" %d %t]", eid, aid, v, trans, added)))
		})
	})
})
