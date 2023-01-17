package time

import "time"

func Date(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

func Coalesce(times ...time.Time) time.Time {
	for _, t := range times {
		if !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}

func Min(l, r time.Time) time.Time {
	if l.Before(r) {
		return l
	} else {
		return r
	}
}

func Max(l, r time.Time) time.Time {
	if l.After(r) {
		return l
	} else {
		return r
	}
}

func EarliestIndex(times ...time.Time) int {
	result := time.Time{}
	index := -1
	for i, t := range times {
		if t.IsZero() {
			continue
		}
		if result.IsZero() || t.Before(result) {
			result = t
			index = i
		}
	}
	return index
}

func Earliest(times ...time.Time) time.Time {
	idx := EarliestIndex(times...)
	if idx >= 0 {
		return times[idx]
	} else {
		return time.Time{}
	}
}

func LatestIndex(times ...time.Time) int {
	result := time.Time{}
	index := -1
	for i, t := range times {
		if t.IsZero() {
			continue
		}
		if result.IsZero() || t.After(result) {
			result = t
			index = i
		}
	}
	return index
}

func Latest(times ...time.Time) time.Time {
	idx := LatestIndex(times...)
	if idx >= 0 {
		return times[idx]
	} else {
		return time.Time{}
	}
}

func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

func FormatDateP(t *time.Time) string {
	if t == nil {
		return "<nil>"
	} else {
		return (*t).Format(DateFormat)
	}
}

const (
	DateFormat = "1/2/2006"
)
