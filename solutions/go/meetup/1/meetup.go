package meetup

// cSpell:ignore teenth wSched

import "time"

type WeekSchedule string

const (
	First  WeekSchedule = "first"
	Second WeekSchedule = "second"
	Third  WeekSchedule = "third"
	Fourth WeekSchedule = "fourth"
	Teenth WeekSchedule = "teenth"
	Last   WeekSchedule = "last"
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	cursor := time.Date(year, month, 1, 0, 0, 0, 0, time.Now().Location())

	seen := 0
	day := 0

	for cursor.Month() == month {
		if cursor.Weekday() == wDay {
			day = cursor.Day()
			seen++
		}

		if wSched == First && seen == 1 {
			return day
		}

		if wSched == Second && seen == 2 {
			return day
		}

		if wSched == Third && seen == 3 {
			return day
		}

		if wSched == Fourth && seen == 4 {
			return day
		}

		if wSched == Teenth && day >= 13 && day <= 19 {
			return day
		}

		cursor = cursor.Add(time.Hour * 24)
	}

	if wSched == Last {
		return day
	}

	return 0
}
