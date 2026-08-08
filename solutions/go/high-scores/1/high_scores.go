package highscores

import (
	"math"
	"slices"
)

type HighScores struct {
	scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{
		scores,
	}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	if len(s.scores) == 0 {
		return 0
	}

	return s.scores[len(s.scores)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	high := 0

	for _, score := range s.scores {
		if score > high {
			high = score
		}
	}

	return high
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	sorted := append([]int{}, s.scores...)
	slices.SortFunc(sorted, func(a, b int) int {
		return b - a
	})

	return sorted[:int(math.Min(float64(len(sorted)), 3))]
}
