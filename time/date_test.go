package time

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Date", func() {

	Context("Min", func() {
		It("Compares", func() {
			Expect(Min(Date(2019, 12, 1), Date(2019, 12, 2))).To(Equal(Date(2019, 12, 1)))
		})
		It("Compares", func() {
			Expect(Min(Date(2019, 12, 2), Date(2019, 12, 1))).To(Equal(Date(2019, 12, 1)))
		})
		It("Compares", func() {
			Expect(Min(Date(2019, 12, 2), Date(2019, 12, 2))).To(Equal(Date(2019, 12, 2)))
		})
	})
	Context("Max", func() {
		It("Compares", func() {
			Expect(Max(Date(2019, 12, 1), Date(2019, 12, 2))).To(Equal(Date(2019, 12, 2)))
		})
		It("Compares", func() {
			Expect(Max(Date(2019, 12, 2), Date(2019, 12, 1))).To(Equal(Date(2019, 12, 2)))
		})
		It("Compares", func() {
			Expect(Max(Date(2019, 12, 2), Date(2019, 12, 2))).To(Equal(Date(2019, 12, 2)))
		})
	})

	Context("Earliest", func() {
		It("Compares several", func() {
			Expect(Earliest(
				Date(2019, 12, 1),
				Date(2019, 12, 2),
				Date(2019, 11, 5),
				Date(2019, 12, 5),
			)).To(Equal(Date(2019, 11, 5)))
		})
		It("Compares several - rejects zero", func() {
			Expect(Earliest(
				Date(2019, 12, 1),
				time.Time{},
				Date(2019, 12, 2),
				Date(2019, 11, 5),
				time.Time{},
				Date(2019, 12, 5),
			)).To(Equal(Date(2019, 11, 5)))
		})
		It("Empty - returns zero", func() {
			Expect(Earliest(
			)).To(Equal(time.Time{}))
		})
		It("Only zero - returns zero", func() {
			Expect(Earliest(
				time.Time{},
				time.Time{},
			)).To(Equal(time.Time{}))
		})
	})
	Context("Latest", func() {
		It("Compares several", func() {
			Expect(Latest(
				Date(2019, 12, 1),
				Date(2019, 12, 2),
				Date(2019, 11, 5),
				Date(2019, 12, 5),
			)).To(Equal(Date(2019, 12, 5)))
		})
		It("Compares several - rejects zero", func() {
			Expect(Latest(
				Date(2019, 12, 1),
				time.Time{},
				Date(2019, 12, 2),
				Date(2019, 12, 5),
				Date(2019, 11, 5),
				time.Time{},
			)).To(Equal(Date(2019, 12, 5)))
		})
		It("Empty - returns zero", func() {
			Expect(Latest(
			)).To(Equal(time.Time{}))
		})
		It("Only zero - returns zero", func() {
			Expect(Latest(
				time.Time{},
				time.Time{},
			)).To(Equal(time.Time{}))
		})
	})
})
