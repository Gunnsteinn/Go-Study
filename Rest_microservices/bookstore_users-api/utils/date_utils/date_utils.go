package date_utils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:08:05Z"
)

// GetNow is a func to returns the current local time with the location set to UTC.
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString uses the GetNow function and returns a textual representation of the formatted time value.
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
