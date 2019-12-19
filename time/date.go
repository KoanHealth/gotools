package time

import "time"

func Date(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

func Min(l,r time.Time) time.Time {
	if l.Before(r) {
		return l
	} else {
		return r
	}
}

func Max(l,r time.Time) time.Time {
	if l.After(r) {
		return l
	} else {
		return r
	}
}

func Earliest(times ...time.Time) time.Time {
	result := time.Time{}
	for _, t := range times {
		if t.IsZero() { continue }
		if result.IsZero() || t.Before(result) {
			result = t
		}
	}
	return result
}

func Latest(times ...time.Time) time.Time {
	result := time.Time{}
	for _, t := range times {
		if t.IsZero() { continue }
		if result.IsZero() || t.After(result) {
			result = t
		}
	}
	return result
}

func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

const (
	DateFormat = "1/2/2006"
)
