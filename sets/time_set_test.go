package sets_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/koanhealth/gotools/sets"
	"time"
	kt "github.com/koanhealth/gotools/time"
)

var _ = Describe("TimeSet", func() {

	time1 := kt.Date(1, 1, 1)
	time2 := kt.Date(2018, 1, 1)
	time3 := kt.Date(2018, 2, 1)
	time4 := kt.Date(2018, 3, 1)
	time5 := kt.Date(2019, 12, 15)
	time6 := kt.Date(2019, 12, 31)

	set := NewTimeSet(time1, time2, time3, time4)

	It("Stores unique, non-zero times", func() {
		input := []time.Time{
			kt.Date(2018, 1, 1),
			kt.Date(2018, 2, 1),
			kt.Date(2018, 3, 1),
			kt.Date(2018, 1, 1),
			kt.Date(2018, 2, 1),
			kt.Date(2019, 12, 31),
			{}, // zero time object
		}
		output := []time.Time{
			kt.Date(1, 1, 1),
			kt.Date(2018, 1, 1),
			kt.Date(2018, 2, 1),
			kt.Date(2018, 3, 1),
			kt.Date(2019, 12, 31),
		}

		set := NewTimeSet(input...)
		Expect(set.SortedItems()).To(Equal(output))
	})

	Describe("HasAny", func() {

		It("returns true if any values match", func() {
			Expect(set.HasAny(time1)).To(BeTrue())
			Expect(set.HasAny(time2, time3)).To(BeTrue())
			Expect(set.HasAny(time1, time5)).To(BeTrue())
			Expect(set.HasAny(time5, time6)).To(BeFalse())
		})

	})

	Describe("HasAll", func() {

		It("returns true if all values match", func() {
			Expect(set.HasAll(time1)).To(BeTrue())
			Expect(set.HasAll(time1, time2)).To(BeTrue())
			Expect(set.HasAll(time3, time6)).To(BeFalse())
			Expect(set.HasAll(time5, time6)).To(BeFalse())
		})

	})

})
