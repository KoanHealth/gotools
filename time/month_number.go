package time

import (
	"time"
)

// 6 digit number representing a particular month
// Example: 201312 represents Dec 2013
type MonthNumber int

func NewMonthNumber(year int, month time.Month) MonthNumber {
	n := year*100 + int(month)
	return MonthNumber(n)
}

func (m MonthNumber) IsValid() bool {
	year := m.Year()
	if year < 0 || year > 9999 {
		return false
	}

	month := m.Month()
	if month < time.January || month > time.December {
		return false
	}

	return true
}
func (m MonthNumber) YearsFromNow(num int) MonthNumber {
	year := m.Year()
	month := int(m.Month())
	return NewMonthNumber(year+num, time.Month(month))
}

func (m MonthNumber) YearsAgo(num int) MonthNumber {
	return m.YearsFromNow(-1 * num)
}

func (m MonthNumber) MonthsFromNow(num int) MonthNumber {
	if num < 0 {
		return m.MonthsAgo(-1 * num)
	}

	year := m.Year()
	month := int(m.Month())
	for i := 0; i < num; i++ {
		month += 1
		if month == 13 {
			year += 1
			month = 1
		}
	}
	return NewMonthNumber(year, time.Month(month))
}

func (m MonthNumber) MonthsAgo(num int) MonthNumber {
	if num < 0 {
		return m.MonthsFromNow(-1 * num)
	}

	year := m.Year()
	month := int(m.Month())
	for i := 0; i < num; i++ {
		month -= 1
		if month == 0 {
			year -= 1
			month = 12
		}
	}
	return NewMonthNumber(year, time.Month(month))
}

func (m MonthNumber) NextMonth() MonthNumber {
	return m.MonthsFromNow(1)
}

func (m MonthNumber) PreviousMonth() MonthNumber {
	return m.MonthsAgo(1)
}

func (m MonthNumber) Year() int {
	return int(m) / 100
}

func (m MonthNumber) Month() time.Month {
	return time.Month(int(m) % 100)
}

func (m MonthNumber) ToDate(dayOfMonth int) time.Time {
	return time.Date(m.Year(), m.Month(), dayOfMonth, 0, 0, 0, 0, time.UTC)
}

func (m MonthNumber) FirstDay() time.Time {
	return m.ToDate(1)
}

func (m MonthNumber) FirstDayOfFollowingMonth() time.Time {
	return m.NextMonth().FirstDay()
}

func (m MonthNumber) LastDay() time.Time {
	return m.FirstDay().AddDate(0, 1, -1) // Add 1 month, Subtract 1 day
}

func (m MonthNumber) NextMonths(num int) []MonthNumber {
	return m.Range(m.MonthsFromNow(num - 1))
}

func (m MonthNumber) NextYear() []MonthNumber {
	return m.NextMonths(12)
}

func (m MonthNumber) PreviousMonths(num int) []MonthNumber {
	return m.Range(m.MonthsAgo(num - 1))
}

func (m MonthNumber) PreviousYear() []MonthNumber {
	return m.PreviousMonths(12)
}

func (m MonthNumber) Range(other MonthNumber) []MonthNumber {
	var min, max, current MonthNumber
	if m < other {
		min, max = m, other
	} else {
		min, max = other, m
	}

	current = min
	months := []MonthNumber{current}
	year := current.Year()
	month := int(current.Month())
	for current < max {
		month += 1
		if month == 13 {
			year += 1
			month = 1
		}

		current = NewMonthNumber(year, time.Month(month))
		months = append(months, current)
	}

	return months
}
