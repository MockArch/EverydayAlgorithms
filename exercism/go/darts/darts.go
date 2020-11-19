package darts

import (
	"math"
)

func Score(x, y float64) int {
	target := math.Pow(x, 2) + math.Pow(y, 2)

	switch {
	case target <= 100 && target > 25:
		return 1
	case target <= 25 && target > 1:
		return 5
	case target <= 1 && target >= 0:
		return 10
	default:

		return 0
	}
}
