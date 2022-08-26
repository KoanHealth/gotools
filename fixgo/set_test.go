package fixgo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	Context("Integer Set", func() {

		var set = NewSet(1, 2, 3, 4)

		It("Stores unique integers", func() {
			input := []int{1, 2, 3, 1, 2, 4, 4}
			output := []int{1, 2, 3, 4}

			set := NewSet(input...)
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
	Context("Int64 Set", func() {
		var set = NewSet(int64(1), int64(2), int64(3), int64(4))

		It("Stores unique integers", func() {
			input := []int64{1, 2, 3, 1, 2, 4, 4}
			output := []int64{1, 2, 3, 4}

			set := NewSet(input...)
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

	Context("String Set", func() {
		var set = NewSet("", "A", "B", "C", "TOM")

		It("Stores unique strings (including blank)", func() {
			input := []string{"A", "B", "C", "A", "B", "TOM", ""}
			output := []string{"", "A", "B", "C", "TOM"}

			set := NewSet(input...)
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

		Describe("Stringer", func() {

			It("works", func() {
				Expect(set.String()).To(Equal(", A, B, C, TOM"))
			})

		})

	})

})
