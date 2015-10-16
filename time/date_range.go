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
		return "[ Empty }"
	} else {
		return fmt.Sprintf("[%s - %s]", r.Min.Format(DateFormat), r.Max.Format(DateFormat))
	}
}

func (r DateRange) IsEmpty() bool {
	return r.Min.IsZero() && r.Max.IsZero()
}

func (r DateRange) DaysBetween() int {
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
