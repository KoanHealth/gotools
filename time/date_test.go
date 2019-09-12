package time

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
})
