package geom

import "math"

func EqualsWithTolerance(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
