package slices_test

import (
	. "github.com/KoanHealth/gotools/slices"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StringSlice", func() {

	var (
		slice          = StringSlice{"a", "b", "c", "1", "2", "3"}
		emptySlice     = StringSlice{}
		sliceWithEmpty = StringSlice{" ", "a", "b", "c", "1", "", "2", "3", "   "}
	)

	Describe("Select", func() {
		It("returns items matching filter", func() {
			filter := func(s string) bool { return s == "a" || s == "1" }
			Expect(slice.Select(filter)).To(Equal(StringSlice{"a", "1"}))
		})
	})

	Describe("DeleteIf/Reject", func() {
		It("Removes items matching filter", func() {
			filter := func(s string) bool { return s == "a" || s == "1" }
			Expect(slice.DeleteIf(filter)).To(Equal(StringSlice{"b", "c", "2", "3"}))
			Expect(slice.Reject(filter)).To(Equal(StringSlice{"b", "c", "2", "3"}))
		})
	})

	Describe("Any", func() {
		It("returns true if any items match filter", func() {
			filter := func(s string) bool { return s == "a" || s == "1" }
			Expect(slice.Any(filter)).To(BeTrue())
		})

		It("returns false if no items match filter", func() {
			filter := func(s string) bool { return s == "ZZZ" }
			Expect(slice.Any(filter)).To(BeFalse())
		})
	})

	Describe("All", func() {
		It("returns true if all items match filter", func() {
			filter := func(s string) bool { return len(s) > 0 }
			Expect(slice.All(filter)).To(BeTrue())
		})

		It("returns false if any items fail filter", func() {
			filter := func(s string) bool { return s == "a" }
			Expect(slice.All(filter)).To(BeFalse())
		})
	})

	Describe("Contains", func() {

		Context("with a slice containing data", func() {
			It("returns true if string is in the slice", func() {
				Expect(slice.Contains("b")).To(BeTrue())
			})

			It("returns false if string is not in the slice", func() {
				Expect(slice.Contains("z")).To(BeFalse())
			})

			It("is case sensitive", func() {
				Expect(slice.Contains("A")).To(BeFalse())
			})
		})

		Context("with empty slice", func() {
			It("will always return false", func() {
				attempts := []string{"", " ", "a"}
				for _, attempt := range attempts {
					Expect(emptySlice.Contains(attempt)).To(BeFalse())
				}
			})
		})

	})

	Describe("Index", func() {
		It("returns index of matching item", func() {
			Expect(slice.Index("b")).To(Equal(1))
		})

		It("returns -1 for non-matching item", func() {
			Expect(slice.Index("ZZZ")).To(Equal(-1))
		})
	})

	Describe("Each", func() {
		It("Call function for each item", func() {
			counter := 0
			f := func(s string) { counter++ }
			slice.Each(f)
			Expect(counter).To(Equal(len(slice)))
		})
	})

	Describe("Compact", func() {
		It("Removes empty strings", func() {
			Expect(sliceWithEmpty.Compact()).To(Equal(slice))
		})
	})

	Describe("First", func() {
		It("Returns first item", func() {
			s, ok := slice.First()
			Expect(s).To(Equal(slice[0]))
			Expect(ok).To(BeTrue())
		})

		It("Empty slice returns not ok", func() {
			s, ok := emptySlice.First()
			Expect(s).To(Equal(""))
			Expect(ok).To(BeFalse())
		})
	})

	Describe("Last", func() {
		It("Returns last item", func() {
			s, ok := slice.Last()
			Expect(s).To(Equal(slice[len(slice)-1]))
			Expect(ok).To(BeTrue())
		})

		It("Empty slice returns not ok", func() {
			s, ok := emptySlice.Last()
			Expect(s).To(Equal(""))
			Expect(ok).To(BeFalse())
		})
	})

})
