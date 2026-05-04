package darts

import "math"

type Circle struct {
    Limit int
    Score int
}

func Score(x, y float64) int {
	var circles = [3]Circle {
        {Limit: 1, Score: 10},
        {Limit: 5, Score: 5},
        {Limit: 10, Score: 1},
    }

    var distance float64 = math.Sqrt(x * x + y * y)

    for _, circle := range circles {
        if distance <= float64(circle.Limit) {
        	return circle.Score
        }
    }

    return 0
}
