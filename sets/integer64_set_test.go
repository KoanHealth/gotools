package sets_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/koanhealth/gotools/sets"
)

var _ = Describe("Integer64Set", func() {

	var set = NewInteger64Set(1, 2, 3, 4)

	It("Stores unique integers", func() {
		input := []int64{1, 2, 3, 1, 2, 4, 4}
		output := []int64{1, 2, 3, 4}

		set := NewInteger64Set(input...)
		Expect(set.SortedItems()).To(Equal(output))
	})

	Describe("HasAny", func() {

		It("returns true if any values match", func() {
			Expect(set.HasAny(1)).To(BeTrue())
			Expect(set.HasAny(2, 4)).To(BeTrue())
			Expect(set.HasAny(2, 19)).To(BeTrue())
			Expect(set.HasAny(42, 26)).To(BeFalse())
		})

	})

	Describe("HasAll", func() {

		It("returns true if all values match", func() {
			Expect(set.HasAll(1)).To(BeTrue())
			Expect(set.HasAll(1, 2)).To(BeTrue())
			Expect(set.HasAll(1, 10)).To(BeFalse())
			Expect(set.HasAll(20, 30)).To(BeFalse())
		})

	})

	Describe("Stringer", func() {

		It("works", func() {
			Expect(set.String()).To(Equal("1, 2, 3, 4"))
		})

	})

})
