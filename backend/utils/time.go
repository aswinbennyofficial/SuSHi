package utils

import (
	"time"
)

func RoundToNearestMinute(t time.Time) time.Time {
	// Get the seconds and nanoseconds part of the time
	sec := t.Second()
	

	// If seconds are 30 or more, round up to the next minute
	if sec >= 30 {
		// Truncate to the start of the minute and add one minute
		return t.Truncate(time.Minute).Add(time.Minute)
	} else {
		// Truncate to the start of the minute
		return t.Truncate(time.Minute)
	}
}