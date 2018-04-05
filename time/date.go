package time

import "time"

func Date(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

const (
	DateFormat = "1/2/2006"
)
