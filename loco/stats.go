package loco

import (
	"math"
)

// Sum -
func Sum(values []float64) float64 {
	sum := 0.0
	for _, val := range values {
		sum += val
	}
	return sum
}

// Mean -
func Mean(values []float64) float64 {
	return Sum(values) / float64(len(values))
}

// StDev -
func StDev(values []float64) (float64, float64) {
	count := float64(len(values))
	mean := Mean(values)
	var sd float64

	for _, val := range values {
		sd += math.Pow(val-mean, 2)
	}
	sd = math.Sqrt(sd / count)

	return mean, sd
}
