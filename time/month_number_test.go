package time

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Month Number", func() {

	It("returns the correct year", func() {
		Expect(MonthNumber(201312).Year()).To(Equal(2013))
	})

	It("returns the correct month", func() {
		Expect(MonthNumber(201306).Month()).To(Equal(time.June))
	})

	Context("Validations", func() {

		invalidMonths := []int{-1, 42, 10000000}

		It("Identifies invalid months", func() {
			for _, m := range invalidMonths {
				GinkgoWriter.Write([]byte(fmt.Sprintf("Month: %v\n", m)))
				Expect(MonthNumber(m).IsValid()).To(BeFalse())
			}
		})

		It("Default Value is not valid", func() {
			var unset int
			invalid := MonthNumber(unset)
			Expect(invalid.IsValid()).To(BeFalse())
		})

	})

	Context("Date Conversions", func() {
		It("Parses a date", func() {
			Expect(int(MonthNumberForDate(Date(2023, 5, 15)))).To(Equal(202305))
		})
		It("ToDate sets the day of month", func() {
			expected := time.Date(2013, 12, 13, 0, 0, 0, 0, time.UTC)
			Expect(MonthNumber(201312).ToDate(13)).To(Equal(expected))
		})

		It("FirstDay sets the first day of the month", func() {
			expected := time.Date(2013, 12, 1, 0, 0, 0, 0, time.UTC)
			Expect(MonthNumber(201312).FirstDay()).To(Equal(expected))
		})

		It("FirstDayOfFollowingMonth sets the first day of the next month", func() {
			expected := MonthNumber(201401).FirstDay()
			Expect(MonthNumber(201312).FirstDayOfFollowingMonth()).To(Equal(expected))
		})

		It("LastDay sets the first day of the month", func() {
			expected := time.Date(2013, 12, 31, 0, 0, 0, 0, time.UTC)
			Expect(MonthNumber(201312).LastDay()).To(Equal(expected))
		})

		It("LastDay catches leap years", func() {
			expected := time.Date(2012, 2, 29, 0, 0, 0, 0, time.UTC)
			Expect(MonthNumber(201202).LastDay()).To(Equal(expected))
		})

		It("Makes a Date Range", func() {
			dr := NewMonthNumber(2023, 6).DateRange()
			Expect(dr.Min).To(Equal(Date(2023, 06, 01)))
			Expect(dr.Max).To(Equal(Date(2023, 06, 30)))

		})

	})

	Context("Next/Previous", func() {

		It("NextMonths returns array containing current month and next (num - 1) months", func() {
			expected := []MonthNumber{201211, 201212, 201301, 201302}
			Expect(MonthNumber(201211).NextMonths(4)).To(Equal(expected))
		})

		It("NextYear returns array containing next 12 months", func() {
			expected := []MonthNumber{201211, 201212, 201301, 201302, 201303, 201304, 201305, 201306, 201307, 201308, 201309, 201310}
			Expect(MonthNumber(201211).NextYear()).To(Equal(expected))
		})

		It("PreviousMonths returns array containing current month and previous (num - 1) months", func() {
			expected := []MonthNumber{201211, 201212, 201301, 201302}
			Expect(MonthNumber(201302).PreviousMonths(4)).To(Equal(expected))
		})

		It("PreviousYear returns array containing previous 12 months", func() {
			expected := []MonthNumber{201211, 201212, 201301, 201302, 201303, 201304, 201305, 201306, 201307, 201308, 201309, 201310}
			Expect(MonthNumber(201310).PreviousYear()).To(Equal(expected))
		})
	})

	Context("FirstMonthOfYear/LastMonthOfYear", func() {

		It("FirstMonthOfYear returns January of the current year", func() {
			Expect(MonthNumber(201211).FirstMonthOfYear()).To(Equal(MonthNumber(201201)))
		})

		It("LastMonthOfYear returns December of the current year", func() {
			Expect(MonthNumber(201211).LastMonthOfYear()).To(Equal(MonthNumber(201212)))
		})

	})

	Context("Ago/FromNow", func() {
		It("MonthsFromNow works", func() {
			start := MonthNumber(201406)
			Expect(start.MonthsFromNow(1)).To(Equal(MonthNumber(201407)))
			Expect(start.MonthsFromNow(5)).To(Equal(MonthNumber(201411)))
			Expect(start.MonthsFromNow(7)).To(Equal(MonthNumber(201501)))
			Expect(start.MonthsFromNow(0)).To(Equal(MonthNumber(201406)))
			// Boneheaded, but it should handle it
			Expect(start.MonthsFromNow(-1)).To(Equal(MonthNumber(201405)))
			Expect(start.MonthsFromNow(-7)).To(Equal(MonthNumber(201311)))
		})
		It("Next/PreviousMonth", func() {
			start := MonthNumber(201406)
			Expect(start.NextMonth()).To(Equal(MonthNumber(201407)))
			Expect(start.PreviousMonth()).To(Equal(MonthNumber(201405)))
		})
		It("YearsFromNow works", func() {
			start := MonthNumber(201406)
			Expect(start.YearsFromNow(1)).To(Equal(MonthNumber(201506)))
			Expect(start.YearsFromNow(2)).To(Equal(MonthNumber(201606)))
			// Boneheaded, but it should handle it
			Expect(start.YearsFromNow(-1)).To(Equal(MonthNumber(201306)))
			Expect(start.YearsFromNow(-2)).To(Equal(MonthNumber(201206)))
		})
		It("MonthsAgo works", func() {
			start := MonthNumber(201406)
			Expect(start.MonthsAgo(1)).To(Equal(MonthNumber(201405)))
			Expect(start.MonthsAgo(5)).To(Equal(MonthNumber(201401)))
			Expect(start.MonthsAgo(7)).To(Equal(MonthNumber(201311)))
			Expect(start.MonthsAgo(0)).To(Equal(MonthNumber(201406)))
			// Boneheaded, but it should handle it
			Expect(start.MonthsAgo(-1)).To(Equal(MonthNumber(201407)))
			Expect(start.MonthsAgo(-7)).To(Equal(MonthNumber(201501)))
		})
		It("YearsAgo works", func() {
			start := MonthNumber(201406)
			Expect(start.YearsAgo(1)).To(Equal(MonthNumber(201306)))
			Expect(start.YearsAgo(2)).To(Equal(MonthNumber(201206)))
			// Boneheaded, but it should handle it
			Expect(start.YearsAgo(-1)).To(Equal(MonthNumber(201506)))
			Expect(start.YearsAgo(-2)).To(Equal(MonthNumber(201606)))
		})
	})

	Context("Range", func() {

		It("Current month number < other", func() {
			expected := []MonthNumber{201211, 201212, 201301, 201302}
			Expect(MonthNumber(201211).Range(201302)).To(Equal(expected))
		})

		It("Current month number > other", func() {
			expected := []MonthNumber{201211, 201212, 201301, 201302}
			Expect(MonthNumber(201302).Range(201211)).To(Equal(expected))
		})

		It("Current month number = other", func() {
			expected := []MonthNumber{201302}
			Expect(MonthNumber(201302).Range(201302)).To(Equal(expected))
		})
	})

	Context("Equal", func() {

		It("Returns true if they are equal", func() {
			Expect(MonthNumber(202012).Equal(MonthNumber(202012))).To(BeTrue())
		})

		It("Returns false if they are not equal", func() {
			Expect(MonthNumber(202012).Equal(MonthNumber(202011))).To(BeFalse())
		})

	})

	Context("Before", func() {

		It("Returns true if month before another", func() {
			Expect(MonthNumber(202011).Before(MonthNumber(202012))).To(BeTrue())
		})

		It("Returns false if month equals another", func() {
			Expect(MonthNumber(202012).Before(MonthNumber(202012))).To(BeFalse())
		})

		It("Returns false if month after another", func() {
			Expect(MonthNumber(202012).Before(MonthNumber(202011))).To(BeFalse())
		})

	})

	Context("After", func() {

		It("Returns true if month after another", func() {
			Expect(MonthNumber(202012).After(MonthNumber(202011))).To(BeTrue())
		})

		It("Returns false if month equals another", func() {
			Expect(MonthNumber(202012).After(MonthNumber(202012))).To(BeFalse())
		})

		It("Returns false if month before another", func() {
			Expect(MonthNumber(202011).After(MonthNumber(202012))).To(BeFalse())
		})

	})

})
