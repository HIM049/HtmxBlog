package utils

import "time"

const DateTimeLocalFormat = "2006-01-02T15:04"

// FormatDateTimeLocal formats a time.Time to a string compatible with <input type="datetime-local">.
func FormatDateTimeLocal(t time.Time) string {
	return t.Format(DateTimeLocalFormat)
}

// ParseDateTimeLocal parses a string from <input type="datetime-local"> into a time.Time.
func ParseDateTimeLocal(s string) (time.Time, error) {
	return time.Parse(DateTimeLocalFormat, s)
}
