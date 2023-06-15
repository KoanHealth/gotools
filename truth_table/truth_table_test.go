package truth_table

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Truth Table", func() {
	Context("Empty Table", func() {
		truthTable := Create("ECMO", "Trach", "MV96", "PDX", "MajOR")
		It("Doesn't find anything", func() {
			found, result := truthTable.Result(true, false, true, false, true)
			Expect(found).To(BeFalse())
			Expect(result).To(BeNil())
		})
	})

	Context("Basic Operations", func() {
		truthTable := Create("ECMO", "Trach", "MV96", "PDX", "MajOR").
			Line("D003", true, nil, nil, nil, nil).
			Line("D003", false, true, true, nil, true).
			Line("D003", false, true, nil, true, true).
			Line("D004", false, true, true, nil, false).
			Line("D004", false, true, nil, true, false)

		It("Returns the expected result with boolean arguments", func() {
			found, result := truthTable.Result(true, false, true, false, true)
			Expect(found).To(BeTrue())
			Expect(result).To(Equal("D003"))
		})

		It("Returns the expected result with table arguments", func() {
			found, result := truthTable.Result(Yes, No, Yes, No, Yes)
			Expect(found).To(BeTrue())
			Expect(result).To(Equal("D003"))
		})

		It("Returns the expected result with different arguments", func() {
			found, result := truthTable.Result(false, true, true, false, false)
			Expect(found).To(BeTrue())
			Expect(result).To(Equal("D004"))
		})

		It("Display String", func() {
			str := truthTable.String()
			fmt.Println()
			fmt.Println(str)
		})
	})

})
