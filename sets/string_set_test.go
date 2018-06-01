package sets_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/koanhealth/gotools/sets"
)

var _ = Describe("StringSet", func() {
	var set = NewStringSet("", "A", "B", "C", "TOM")

	It("Stores unique strings (including blank)", func() {
		input := []string{"A", "B", "C", "A", "B", "TOM", ""}
		output := []string{"", "A", "B", "C", "TOM"}

		set := NewStringSet(input...)
		Expect(set.SortedItems()).To(Equal(output))
	})

	Describe("HasAny", func() {

		It("returns true if any values match", func() {
			Expect(set.HasAny("A")).To(BeTrue())
			Expect(set.HasAny("A", "B")).To(BeTrue())
			Expect(set.HasAny("A", "X")).To(BeTrue())
			Expect(set.HasAny("X", "Y")).To(BeFalse())
		})

	})

	Describe("HasAll", func() {

		It("returns true if all values match", func() {
			Expect(set.HasAll("A")).To(BeTrue())
			Expect(set.HasAll("A", "B")).To(BeTrue())
			Expect(set.HasAll("A", "X")).To(BeFalse())
			Expect(set.HasAll("X", "Y")).To(BeFalse())
		})

	})

})
