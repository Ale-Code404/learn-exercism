package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	converted, err := time.Parse(
		"1/2/2006 15:04:05",
		date,
	)

	if err != nil {
		return time.Now()
	}

	return converted
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	schedule, _ := time.Parse(
		"January 2, 2006 15:04:05",
		date,
	)

	return schedule.Unix() < time.Now().Unix()
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	schedule, _ := time.Parse(
		"Monday, January 2, 2006 15:04:05",
		date,
	)

	return schedule.Hour() >= 12 && schedule.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	schedule := Schedule(date)
	formatted := schedule.Format("Monday, January 2, 2006, at 15:04.")

	return "You have an appointment on " + formatted
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now().UTC()

	return time.Date(now.Year(), 9, 15, 0, 0, 0, 0, now.Location())
}
