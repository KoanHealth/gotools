package time

import (
	"fmt"
	"time"
)

type DateRange struct {
	Min time.Time
	Max time.Time
}

func NewEmptyDateRange() DateRange {
	return DateRange{}
}

func NewUnaryDateRange(start time.Time) DateRange {
	return NewDateRange(start, start)
}

func NewDateRange(start, end time.Time) DateRange {
	if start.IsZero() || end.IsZero() {
		return NewEmptyDateRange()
	}

	if start.Before(end) {
		return DateRange{Min: start, Max: end}
	} else {
		return DateRange{Min: end, Max: start}
	}
}

func (r DateRange) String() string {
	if r.IsEmpty() {
		return "[ Empty ]"
	} else {
		return fmt.Sprintf("[%s - %s]", r.Min.Format(DateFormat), r.Max.Format(DateFormat))
	}
}
func (r DateRange) IncludeNextYears(n int) DateRange {
	return NewDateRange(r.Min, r.Max.AddDate(n, 0, 0))
}

func (r DateRange) IncludePreviousYears(n int) DateRange {
	return NewDateRange(r.Min.AddDate(-n, 0, 0), r.Max)
}

func (r DateRange) IsEmpty() bool {
	return r.Min.IsZero() && r.Max.IsZero()
}

func (r DateRange) Days() int {
	if r.IsEmpty() {
		return 0
	}
	return r.DaysBetween() + 1
}

func (r DateRange) DaysBetween() int {
	return r.Duration()
}

func (r DateRange) Duration() int {
	// Since the time isn't specified for the Min/Max day, the evaluation takes place in UTC
	// daylight savings time is not an issue
	return int(r.Max.Sub(r.Min).Hours() / 24.0)
}

func (r DateRange) Equals(other DateRange) bool {
	return r.Min.Equal(other.Min) && r.Max.Equal(other.Max)
}

func (r DateRange) Includes(moment time.Time) bool {
	return (r.Min.Equal(moment) || r.Min.Before(moment)) && (r.Max.Equal(moment) || r.Max.After(moment))
}

func (r DateRange) CompletelyIncludes(dr DateRange) bool {
	if r.IsEmpty() {
		return false
	} else if dr.IsEmpty() {
		return true
	} else {
		return (r.Min.Equal(dr.Min) || r.Min.Before(dr.Min)) && (r.Max.Equal(dr.Max) || r.Max.After(dr.Max))
	}
}

func (r DateRange) Overlaps(dr DateRange) bool {
	return r.Includes(dr.Min) || r.Includes(dr.Max) || dr.Includes(r.Min) || dr.Includes(r.Max)
}

func (r DateRange) Intersection(dr DateRange) DateRange {
	if r.IsEmpty() || dr.IsEmpty() {
		return NewEmptyDateRange()
	}

	if r.Overlaps(dr) {
		min := r.Min
		if r.Min.Before(dr.Min) {
			min = dr.Min
		}
		max := r.Max
		if r.Max.After(dr.Max) {
			max = dr.Max
		}
		return NewDateRange(min, max)

	} else {
		return NewEmptyDateRange()
	}
}

func (r DateRange) Union(dr DateRange) DateRange {
	if r.IsEmpty() && dr.IsEmpty() {
		return NewEmptyDateRange()
	}

	if r.IsEmpty() {
		return dr
	}

	if dr.IsEmpty() {
		return r
	}

	min := r.Min
	if dr.Min.Before(r.Min) {
		min = dr.Min
	}
	max := r.Max
	if dr.Max.After(r.Max) {
		max = dr.Max
	}
	return NewDateRange(min, max)
}

func (r DateRange) IsAdjacentTo(dr DateRange) bool {
	return r.IsImmediatelyAfter(dr) || r.IsImmediatelyBefore(dr)
}

func (r DateRange) IsAfter(dr DateRange) bool {
	return TimeGreaterThanOrEqualTo(r.Min, dr.Max)
}

func (r DateRange) IsImmediatelyAfter(dr DateRange) bool {
	return r.IsAfter(dr) && TimeGreaterThanOrEqualTo(dr.Max.AddDate(0, 0, 1), r.Min)
}

func (r DateRange) IsBefore(dr DateRange) bool {
	return TimeLessThanOrEqualTo(r.Max, dr.Min)
}

func (r DateRange) IsImmediatelyBefore(dr DateRange) bool {
	return r.IsBefore(dr) && TimeLessThanOrEqualTo(dr.Min.AddDate(0, 0, -1), r.Max)
}

func TimeLessThanOrEqualTo(lhs, rhs time.Time) bool {
	return lhs.Equal(rhs) || lhs.Before(rhs)
}

func TimeGreaterThanOrEqualTo(lhs, rhs time.Time) bool {
	return lhs.Equal(rhs) || lhs.After(rhs)
}

func EarliestStart(drs ...DateRange) DateRange {
	var dates []time.Time
	for _, r := range drs {
		dates = append(dates, r.Min)
	}
	idx := EarliestIndex(dates...)
	if idx >= 0 {
		return drs[idx]
	} else {
		return NewEmptyDateRange()
	}
}

func LatestStart(drs ...DateRange) DateRange {
	var dates []time.Time
	for _, r := range drs {
		dates = append(dates, r.Min)
	}
	idx := LatestIndex(dates...)
	if idx >= 0 {
		return drs[idx]
	} else {
		return NewEmptyDateRange()
	}
}
func EarliestEnd(drs ...DateRange) DateRange {
	var dates []time.Time
	for _, r := range drs {
		dates = append(dates, r.Max)
	}
	idx := EarliestIndex(dates...)
	if idx >= 0 {
		return drs[idx]
	} else {
		return NewEmptyDateRange()
	}
}

func LatestEnd(drs ...DateRange) DateRange {
	var dates []time.Time
	for _, r := range drs {
		dates = append(dates, r.Max)
	}
	idx := LatestIndex(dates...)
	if idx >= 0 {
		return drs[idx]
	} else {
		return NewEmptyDateRange()
	}
}

func (r DateRange) MonthNumbers() []MonthNumber {
	if r.IsEmpty() {
		return nil
	} else {
		return MonthNumberForDate(r.Min).Range(MonthNumberForDate(r.Max))
	}
}
