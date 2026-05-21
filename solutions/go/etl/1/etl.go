package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	scores := make(map[string]int, 26)

	for score, letters := range in {
		for _, letter := range letters {
			scores[strings.ToLower(letter)] = score
		}
	}

	return scores
}
