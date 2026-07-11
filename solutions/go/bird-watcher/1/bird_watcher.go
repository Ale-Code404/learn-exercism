package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0

	for _, birds := range birdsPerDay {
		total += birds
	}

	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	length := len(birdsPerDay) - 1
	weekStart := (week - 1) * 7
	weekEnd := weekStart + 7

	if weekStart > length {
		return 0
	}

	days := []int{}
	if weekEnd <= length {
		days = birdsPerDay[weekStart:weekEnd]
	} else {
		days = birdsPerDay[weekStart:]
	}

	return TotalBirdCount(days)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}

	return birdsPerDay
}
