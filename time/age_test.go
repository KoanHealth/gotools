package time

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type AgeTestCandidate struct {
	BirthDate    time.Time
	CheckingTime time.Time
	ExpectedAge  int
}

var _ = Describe("Age", func() {

	var AgeTestCandidates = []AgeTestCandidate{
		{time.Date(2000, 3, 14, 0, 0, 0, 0, time.UTC), time.Date(2010, 3, 14, 0, 0, 0, 0, time.UTC), 10},
		{time.Date(2001, 3, 14, 0, 0, 0, 0, time.UTC), time.Date(2009, 3, 14, 0, 0, 0, 0, time.UTC), 8},
		{time.Date(2004, 6, 18, 0, 0, 0, 0, time.UTC), time.Date(2005, 5, 12, 0, 0, 0, 0, time.UTC), 0},
	}

	It("Calculates age correctly", func() {
		for _, candidate := range AgeTestCandidates {
			gotAge := AgeAt(candidate.BirthDate, candidate.CheckingTime)
			Expect(gotAge).To(Equal(candidate.ExpectedAge))
		}
	})

	var AgeTestInMonthsCandidates = []AgeTestCandidate{
		{time.Date(2000, 3, 14, 0, 0, 0, 0, time.UTC), time.Date(2010, 3, 14, 0, 0, 0, 0, time.UTC), 120},
		{time.Date(2001, 3, 14, 0, 0, 0, 0, time.UTC), time.Date(2009, 3, 14, 0, 0, 0, 0, time.UTC), 96},
		{time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2018, 3, 14, 0, 0, 0, 0, time.UTC), 2},
		{time.Date(2004, 6, 18, 0, 0, 0, 0, time.UTC), time.Date(2005, 5, 12, 0, 0, 0, 0, time.UTC), 10},
	}

	It("Calculates age in months correctly", func() {
		for _, candidate := range AgeTestInMonthsCandidates {
			gotAgeInMonths := AgeAtInMonths(candidate.BirthDate, candidate.CheckingTime)
			Expect(gotAgeInMonths).To(Equal(candidate.ExpectedAge))
		}
	})

})
