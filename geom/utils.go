package geom

import "math"

func EqualsWithTolerance(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func clamp(x, min, max int) int {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
