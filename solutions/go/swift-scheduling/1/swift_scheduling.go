package swiftscheduling

import (
	"math"
	"strconv"
	"time"
)

func DeliveryDate(start, delivery string) string {
	date, err := time.Parse("2006-01-02T15:04:05", start)
	if err != nil {
		panic("Error on parsing the start date")
	}

	end := date

	if delivery == "NOW" {
		end = date.Add(time.Hour * 2)
	}

	if delivery == "ASAP" {
		if date.Hour() < 13 {
			end = time.Date(date.Year(), date.Month(), date.Day(), 17, 0, 0, 0, date.Location())
		} else {
			end = date.Add(time.Hour * 24)
			end = time.Date(end.Year(), end.Month(), end.Day(), 13, 0, 0, 0, date.Location())
		}
	}

	if delivery == "EOW" {
		if date.Weekday() > time.Wednesday {
			days := time.Duration(7 - int(date.Weekday()))

			end = date.Add(time.Hour * 24 * days)
			end = time.Date(end.Year(), end.Month(), end.Day(), 20, 0, 0, 0, date.Location())
		} else {
			days := time.Duration(int(time.Friday) - int(date.Weekday()))

			end = date.Add(time.Hour * 24 * days)
			end = time.Date(end.Year(), end.Month(), end.Day(), 17, 0, 0, 0, date.Location())
		}
	}

	if delivery[len(delivery)-1] == 'M' {
		month, err := strconv.ParseInt(delivery[0:len(delivery)-1], 10, 0)
		if err != nil {
			panic("Cannot convert the month delivery")
		}

		if date.Month() < time.Month(month) {
			end = time.Date(
				date.Year(),
				time.Month(month),
				firstWorkday(date.Year(), time.Month(month)),
				8,
				0,
				0,
				0,
				date.Location(),
			)
		} else {
			end = time.Date(
				date.Year()+1,
				time.Month(month),
				firstWorkday(date.Year()+1, time.Month(month)),
				8,
				0,
				0,
				0,
				date.Location(),
			)
		}
	}

	if delivery[0] == 'Q' {
		quarter, err := strconv.ParseInt(delivery[1:], 10, 0)
		if err != nil {
			panic("Cannot convert the month on quarter delivery")
		}

		meetQuarter := int(math.Ceil(float64(date.Month()) / 3))
		month := time.Month(quarter * 3)

		if time.Month(meetQuarter) <= time.Month(quarter) {
			end = time.Date(
				date.Year(),
				month,
				lastWorkday(date.Year(), month),
				8,
				0,
				0,
				0,
				date.Location(),
			)
		} else {
			end = time.Date(
				date.Year()+1,
				month,
				lastWorkday(date.Year()+1, month),
				8,
				0,
				0,
				0,
				date.Location(),
			)
		}
	}

	return end.Format("2006-01-02T15:04:05")
}

func firstWorkday(year int, month time.Month) int {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	if date.Weekday() >= time.Monday && date.Weekday() <= time.Friday {
		return 1
	} else if date.Weekday() == time.Sunday {
		return 2
	} else {
		return 3
	}
}

func lastWorkday(year int, month time.Month) int {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	date = date.AddDate(0, 1, 0)
	date = date.AddDate(0, 0, -1)

	day := date.Day()

	if date.Weekday() >= time.Monday && date.Weekday() <= time.Friday {
		return day
	} else if date.Weekday() == time.Saturday {
		return day - 1
	} else {
		return day - 2
	}
}
