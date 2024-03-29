package time

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DateRange", func() {
	Context("Initialization", func() {
		var (
			today    = Date(2012, 2, 1)
			tomorrow = Date(2012, 2, 2)
		)

		It("Initializes with one value", func() {
			dr := NewUnaryDateRange(today)
			Expect(dr.Min).To(Equal(today))
			Expect(dr.Max).To(Equal(today))
		})

		It("Initializes with two values", func() {
			dr := NewDateRange(today, tomorrow)
			Expect(dr.Min).To(Equal(today))
			Expect(dr.Max).To(Equal(tomorrow))
		})

		It("Initializes with a year", func() {
			dr := NewYearDateRange(2023)
			Expect(dr.Min).To(Equal(Date(2023, 1, 1)))
			Expect(dr.Max).To(Equal(Date(2023, 12, 31)))
		})

		It("Initializes with any value zero gives empty range", func() {
			zero := time.Time{}
			Expect(NewUnaryDateRange(zero).IsEmpty()).To(BeTrue())
			Expect(NewDateRange(zero, tomorrow).IsEmpty()).To(BeTrue())
			Expect(NewDateRange(today, zero).IsEmpty()).To(BeTrue())
		})

		It("Initializes with reversed arguments, range should be corrected", func() {
			dr := NewDateRange(tomorrow, today)
			Expect(dr.Min).To(Equal(today))
			Expect(dr.Max).To(Equal(tomorrow))
		})
	})

	Context("Extension", func() {
		It("Empty Range extended forward is still empty", func() {
			r0 := NewEmptyDateRange()
			r1 := r0.IncludeNextYears(1)
			Expect(r1.IsEmpty()).To(BeTrue())
		})
		It("Empty Range extended backwards is still empty", func() {
			r0 := NewEmptyDateRange()
			r1 := r0.IncludePreviousYears(1)
			Expect(r1.IsEmpty()).To(BeTrue())
		})
		It("Extend date range forward", func() {
			r0 := NewDateRange(Date(2012, 1, 1), Date(2012, 12, 31))
			r1 := r0.IncludeNextYears(1)
			Expect(r1.IsEmpty()).To(BeFalse())
			Expect(r1.Min).To(Equal(Date(2012, 1, 1)))
			Expect(r1.Max).To(Equal(Date(2013, 12, 31)))
		})
		It("Extend date range backwards", func() {
			r0 := NewDateRange(Date(2012, 1, 1), Date(2012, 12, 31))
			r1 := r0.IncludePreviousYears(1)
			Expect(r1.IsEmpty()).To(BeFalse())
			Expect(r1.Min).To(Equal(Date(2011, 1, 1)))
			Expect(r1.Max).To(Equal(Date(2012, 12, 31)))
		})

		It("Converts to Month Number Range", func() {
			r := NewDateRange(Date(2023, 1, 1), Date(2023, 12, 31))
			rng := r.MonthNumbers()
			Expect(rng).To(HaveLen(12))
			for m := 1; m <= 12; m++ {
				Expect(rng).To(ContainElement(NewMonthNumber(2023, time.Month(m))))
			}
		})
	})

	Context("Includes", func() {
		It("returns true if date is inside range", func() {
			r1 := NewDateRange(Date(2012, 1, 2), Date(2012, 1, 5))
			Expect(r1.Includes(Date(2012, 1, 1))).To(BeFalse())
			Expect(r1.Includes(Date(2012, 1, 2))).To(BeTrue())
			Expect(r1.Includes(Date(2012, 1, 3))).To(BeTrue())
			Expect(r1.Includes(Date(2012, 1, 4))).To(BeTrue())
			Expect(r1.Includes(Date(2012, 1, 5))).To(BeTrue())
			Expect(r1.Includes(Date(2012, 1, 6))).To(BeFalse())
		})
	})
	Context("Completely Includes", func() {
		It("returns true if range is inside another", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 1, 15), Date(2012, 1, 15))
			Expect(r1.CompletelyIncludes(r2)).To(BeTrue())
		})

		It("returns true if same range ", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			Expect(r1.CompletelyIncludes(r2)).To(BeTrue())
		})

		It("returns true if same start ", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 1, 1), Date(2012, 1, 15))
			Expect(r1.CompletelyIncludes(r2)).To(BeTrue())
		})

		It("returns true if same end ", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 1, 15), Date(2012, 2, 1))
			Expect(r1.CompletelyIncludes(r2)).To(BeTrue())
		})

		It("returns true if other is empty", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewEmptyDateRange()
			Expect(r1.CompletelyIncludes(r2)).To(BeTrue())
		})

		It("returns false if this is empty", func() {
			r1 := NewEmptyDateRange()
			r2 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			Expect(r1.CompletelyIncludes(r2)).To(BeFalse())
		})

		It("returns false if range overlaps min", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2011, 12, 15), Date(2012, 1, 15))
			Expect(r1.CompletelyIncludes(r2)).To(BeFalse())
		})

		It("returns false if range overlaps max", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 1, 15), Date(2012, 2, 15))
			Expect(r1.CompletelyIncludes(r2)).To(BeFalse())
		})

		It("returns false if range doesnt overlap", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 3, 1), Date(2012, 3, 15))
			Expect(r1.CompletelyIncludes(r2)).To(BeFalse())
		})
	})

	Context("Overlap", func() {
		It("returns true if range is inside another", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 1, 15), Date(2012, 1, 15))
			Expect(r1.Overlaps(r2)).To(BeTrue())
		})

		It("returns true if range overlaps min", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2011, 12, 15), Date(2012, 1, 15))
			Expect(r1.Overlaps(r2)).To(BeTrue())
		})

		It("returns true if range overlaps max", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 1, 15), Date(2012, 2, 15))
			Expect(r1.Overlaps(r2)).To(BeTrue())
		})

		It("returns false if range doesnt overlap", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 3, 1), Date(2012, 3, 15))
			Expect(r1.Overlaps(r2)).To(BeFalse())
		})
	})

	Context("Duration (Days Between)", func() {
		It("Evaluates duration correctly", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 1, 15))
			Expect(r1.DaysBetween()).To(Equal(14))
		})

		It("Evaluates duration correctly", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 1, 15))
			Expect(r1.Duration()).To(Equal(14))
		})

		It("Evaluates duration correctly", func() {
			r1 := NewDateRange(Date(2012, 1, 31), Date(2012, 2, 17))
			Expect(r1.Duration()).To(Equal(17))
		})

		It("Empty range has zero duration", func() {
			Expect(NewEmptyDateRange().Duration()).To(Equal(0))
		})
	})

	Context("Duration (Days)", func() {
		It("Evaluates duration correctly", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 1, 15))
			Expect(r1.Days()).To(Equal(15))
		})

		It("Evaluates duration correctly", func() {
			r1 := NewDateRange(Date(2012, 1, 31), Date(2012, 2, 17))
			Expect(r1.Days()).To(Equal(18))
		})

		It("One day range has 1 day", func() {
			r1 := NewDateRange(Date(2019, 1, 1), Date(2019, 1, 1))
			Expect(r1.Days()).To(Equal(1))
		})

		It("Empty range has zero duration", func() {
			Expect(NewEmptyDateRange().Days()).To(Equal(0))
		})

	})

	Context("Operations", func() {
		It("Equals evaluates correctly", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))

			Expect(r1.Equals(NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1)))).To(BeTrue())
			Expect(r1.Equals(NewDateRange(Date(2011, 1, 1), Date(2012, 2, 1)))).To(BeFalse())
			Expect(r1.Equals(NewDateRange(Date(2012, 1, 1), Date(2012, 2, 2)))).To(BeFalse())
		})
		It("Intersection returns new range where they intersect", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 1, 15), Date(2012, 2, 15))
			i := r1.Intersection(r2)
			Expect(i.IsEmpty()).To(BeFalse())
			Expect(i.Equals(NewDateRange(r2.Min, r1.Max))).To(BeTrue())
		})
		It("Intersection with empty range yields empty range", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			i := r1.Intersection(NewEmptyDateRange())
			Expect(i.IsEmpty()).To(BeTrue())
		})
		It("Intersection returns empty range when they don't intersect", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 2, 15), Date(2012, 2, 16))
			i := r1.Intersection(r2)
			Expect(i.IsEmpty()).To(BeTrue())
		})
		It("Union returns new range with max of upper and lower bounds", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 1, 15), Date(2012, 2, 15))
			Expect(r1.Union(r2).Equals(NewDateRange(r1.Min, r2.Max))).To(BeTrue())
		})
		It("Union returns new range even if ranges overlap", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			r2 := NewDateRange(Date(2012, 3, 15), Date(2012, 4, 15))
			Expect(r1.Union(r2).Equals(NewDateRange(r1.Min, r2.Max))).To(BeTrue())
		})
		It("Union with empty is self", func() {
			r1 := NewDateRange(Date(2012, 1, 1), Date(2012, 2, 1))
			Expect(r1.Union(NewEmptyDateRange()).Equals(r1)).To(BeTrue())
			Expect(NewEmptyDateRange().Union(r1).Equals(r1)).To(BeTrue())
			Expect(NewEmptyDateRange().Union(NewEmptyDateRange()).IsEmpty()).To(BeTrue())
		})
	})
	Context("relations", func() {
		var (
			r1 = NewDateRange(Date(2012, 1, 1), Date(2012, 1, 15))
			r2 = NewDateRange(Date(2012, 1, 20), Date(2012, 1, 25))
		)
		It("after", func() {
			Expect(r1.IsAfter(r2)).To(BeFalse())
			Expect(r2.IsAfter(r1)).To(BeTrue())
			Expect(NewDateRange(Date(2012, 1, 15), Date(2012, 1, 20)).IsAfter(r1)).To(BeTrue())
		})

		It("before", func() {
			Expect(r1.IsBefore(r2)).To(BeTrue())
			Expect(r2.IsBefore(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2011, 12, 20), Date(2012, 1, 1)).IsBefore(r1)).To(BeTrue())
		})

		It("ImmediatelyAfter", func() {
			Expect(NewDateRange(Date(2011, 12, 1), Date(2011, 12, 15)).IsImmediatelyAfter(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2011, 12, 1), Date(2012, 1, 2)).IsImmediatelyAfter(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2012, 1, 3), Date(2012, 1, 7)).IsImmediatelyAfter(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2012, 1, 14), Date(2012, 1, 17)).IsImmediatelyAfter(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2012, 1, 20), Date(2012, 1, 27)).IsImmediatelyAfter(r1)).To(BeFalse())

			Expect(NewDateRange(Date(2012, 1, 15), Date(2012, 1, 31)).IsImmediatelyAfter(r1)).To(BeTrue())
			Expect(NewDateRange(Date(2012, 1, 16), Date(2012, 1, 31)).IsImmediatelyAfter(r1)).To(BeTrue())
		})

		It("ImmediatelyBefore", func() {
			Expect(NewDateRange(Date(2012, 1, 20), Date(2012, 1, 27)).IsImmediatelyBefore(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2012, 1, 14), Date(2012, 1, 17)).IsImmediatelyBefore(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2012, 1, 3), Date(2012, 1, 7)).IsImmediatelyBefore(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2011, 12, 1), Date(2011, 12, 15)).IsImmediatelyBefore(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2011, 12, 1), Date(2012, 1, 2)).IsImmediatelyBefore(r1)).To(BeFalse())

			Expect(NewDateRange(Date(2011, 12, 15), Date(2011, 12, 31)).IsImmediatelyBefore(r1)).To(BeTrue())
			Expect(NewDateRange(Date(2011, 12, 15), Date(2012, 1, 1)).IsImmediatelyBefore(r1)).To(BeTrue())
		})

		It("AdjacentTo", func() {
			Expect(NewDateRange(Date(2012, 1, 20), Date(2012, 1, 27)).IsAdjacentTo(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2012, 1, 14), Date(2012, 1, 17)).IsAdjacentTo(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2012, 1, 3), Date(2012, 1, 7)).IsAdjacentTo(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2011, 12, 1), Date(2011, 12, 15)).IsAdjacentTo(r1)).To(BeFalse())
			Expect(NewDateRange(Date(2011, 12, 1), Date(2012, 1, 2)).IsAdjacentTo(r1)).To(BeFalse())

			Expect(NewDateRange(Date(2012, 1, 15), Date(2012, 1, 31)).IsAdjacentTo(r1)).To(BeTrue())
			Expect(NewDateRange(Date(2012, 1, 16), Date(2012, 1, 31)).IsAdjacentTo(r1)).To(BeTrue())
			Expect(NewDateRange(Date(2011, 12, 15), Date(2011, 12, 31)).IsAdjacentTo(r1)).To(BeTrue())
			Expect(NewDateRange(Date(2011, 12, 15), Date(2012, 1, 1)).IsAdjacentTo(r1)).To(BeTrue())
		})
	})
	Context("Earliest/Latest", func() {
		defaultRanges := []DateRange{
			NewEmptyDateRange(),
			NewDateRange(Date(2012, 1, 20), Date(2012, 1, 27)),
			NewDateRange(Date(2012, 1, 25), Date(2012, 1, 26)),
			NewDateRange(Date(2012, 1, 24), Date(2012, 1, 28)),
		}

		It("Earliest End", func() {
			Expect(EarliestEnd(defaultRanges...)).To(Equal(defaultRanges[2]))
			Expect(EarliestEnd().IsEmpty()).To(BeTrue())
			Expect(EarliestEnd(NewEmptyDateRange(), NewEmptyDateRange()).IsEmpty()).To(BeTrue())
		})

		It("Latest End", func() {
			Expect(LatestEnd(defaultRanges...)).To(Equal(defaultRanges[3]))
			Expect(LatestEnd().IsEmpty()).To(BeTrue())
			Expect(LatestEnd(NewEmptyDateRange(), NewEmptyDateRange()).IsEmpty()).To(BeTrue())
		})

		It("EarliestStart", func() {
			Expect(EarliestStart(defaultRanges...)).To(Equal(defaultRanges[1]))
			Expect(EarliestStart().IsEmpty()).To(BeTrue())
			Expect(EarliestStart(NewEmptyDateRange(), NewEmptyDateRange()).IsEmpty()).To(BeTrue())
		})

		It("LatestStart", func() {
			Expect(LatestStart(defaultRanges...)).To(Equal(defaultRanges[2]))
			Expect(LatestStart().IsEmpty()).To(BeTrue())
			Expect(LatestStart(NewEmptyDateRange(), NewEmptyDateRange()).IsEmpty()).To(BeTrue())
		})

	})
})
