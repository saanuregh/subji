package helper

import "time"

// Date layouts used in app for parsing and formating time.Time.
const (
	DateLayout     = "2006-01-02"
	DateTimeLayout = "2006-01-02 15:04:05"
)

// DaysLeft calculates date left between two time.Time.
func DaysLeft(date1, date2 time.Time) float64 {
	return date1.Sub(date2).Hours() / 24
}

// CurrentDate returns current date.
func CurrentDate() time.Time {
	t, _ := time.Parse(DateLayout, time.Now().Format(DateLayout))
	return t
}

// CurrentDateString returns current date in this format "2006-01-02T15:04:05" string.
func CurrentDateString() string {
	return CurrentDate().Format("2006-01-02T15:04:05")
}
