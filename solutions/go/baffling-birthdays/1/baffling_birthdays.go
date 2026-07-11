package bafflingbirthdays

import (
	"fmt"
	"math/rand"
	"time"
)

func SharedBirthday(dates []time.Time) bool {
	shared := make(map[string]bool)

	for _, date := range dates {
		key := fmt.Sprintf("%d-%d", date.Month(), date.Day())

		_, exists := shared[key]
		if exists {
			return true
		} else {
			shared[key] = true
		}
	}

	return false
}

func RandomBirthdates(size int) []time.Time {
	min := time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.UTC).Unix()

	delta := max - min
	dates := []time.Time{}

	for i := 0; i < size; i++ {
		date := time.Unix(rand.Int63n(delta)+min, 0)
		dates = append(dates, date)
	}

	return dates
}

func EstimatedProbability(size int) float64 {
	probability := 1.0

	for i := range size {
		probability *= (365 - float64(i)) / 365.0
	}

	return (1 - probability) * 100
}
