package geom

func dimension(pts []Coordinate) int {
	if pts == nil || len(pts) == 0 {
		return 3 //unknown, assume default
	}
	dim := int(0)
	for _, c := range pts {
		dim = MaxInt(dim, c.dimensions)
	}
	return dim
}

func measures(pts []Coordinate) int {
	if pts == nil || len(pts) == 0 {
		return 0 //unknown, assume default
	}
	measure := int(0)
	for _, c := range pts {
		measure = MaxInt(measure, c.measure())
	}
	return measure
}

func enforceConsistency(array []Coordinate) {
	// TODO: implement consistency enforcement when more Coordinate types will come on board
}

func enforceConsistencyWithDimensionAndMeasure(array []Coordinate, dimension, measure int) {
	// TODO: implement consistency enforcement when more Coordinate types will come on board
}

// Tests whether an array of Coordinate(s) forms a ring,
//  by checking length and closure.
//  Self-intersection is not checked.
func isRing(coordinates []Coordinate) bool {
	l := len(coordinates)
	if l < 4 {
		return false
	}
	if !coordinates[0].equals2D(&coordinates[l-1]) {
		return false
	}
	return true
}

// Finds a point in a list of points which is not contained in another list of points
func ptNotInList(testPts, pts []Coordinate) *Coordinate {
	for _, c := range testPts {
		if indexOf(c, pts) < 0 {
			return &c
		}
	}
	return nil
}

// Compares two Coordinate arrays
// in the forward direction of their coordinates,
// using lexicographic ordering.
func compare(pts1, pts2 []Coordinate) int {
	i := int(0)
	l1 := len(pts1)
	l2 := len(pts2)
	for i < l1 && i < l2 {
		compare := pts1[i].compareTo(pts2[i])
		if compare != 0 {
			return compare
		}
		i++
	}
	//handle situation when arrays are of different length
	if i < l2 {
		return -1
	}
	if i < l1 {
		return 1
	}
	return 0
}

// Determines which orientation of the Coordinate array
// is (overall) increasing.
// In other words, determines which end of the array is "smaller"
// (using the standard ordering on Coordinate).
// Returns an integer indicating the increasing direction.
// If the sequence is a palindrome, it is defined to be
// oriented in a positive direction.
//
// returns 1 if the array is smaller at the start
// or is a palindrome,
// -1 if smaller at the end
func increasingDirection(pts []Coordinate) int {
	for i, c := range pts {
		j := len(pts) - 1 - i
		// skip equal points on both ends
		comp := c.compareTo(pts[j])
		if comp != 0 {
			return comp
		}
	}
	// array must be a palindrome - defined to be in positive direction
	return 1
}

// Determines whether two Coordinate arrays of equal length
// are equal in opposite directions.
func isEqualReversed(pts1, pts2 []Coordinate) bool {
	for i, c := range pts1 {
		if c.compareTo(pts2[i]) != 0 {
			return false
		}
	}
	return true
}

func copyDeepCoordinateArray(coordinates []Coordinate) []Coordinate {
	result := make([]Coordinate, len(coordinates))
	copy(result, coordinates) // TODO: currently since all types in Coordinate are primitive, this shall result in deep copy, but to be verified later
	return result
}

// Returns whether two consecutive Coordinates equal
func hasRepeatedPoints(pts []Coordinate) bool {
	for i, _ := range pts {
		if pts[i-1] == pts[i] {
			return true
		}
	}
	return false
}

// If the coordinate array argument has repeated points,
// constructs a new array containing no repeated points.
// Otherwise, returns the argument.
func removeRepeatedPoints(pts []Coordinate) []Coordinate {
	if !hasRepeatedPoints(pts) {
		return pts
	}
	keys := make(map[Coordinate]bool)
	var result []Coordinate
	for _, c := range pts {
		if _, v := keys[c]; !v {
			keys[c] = true
			result = append(result, c)
		}
	}
	return result
}

// Collapses a coordinate array to remove all null elements.
func removeNull(pts []*Coordinate) []*Coordinate {
	var result []*Coordinate
	for _, c := range pts {
		if c != nil {
			result = append(result, c)
		}
	}
	return result
}

// Reverses the coordinates in an array in-place.
func reverse(pts []Coordinate) {
	last := len(pts) - 1
	mid := int(last / 2)
	for i := 0; i <= mid; i++ {
		tmp := pts[i]
		pts[i] = pts[last-i]
		pts[last-i] = tmp
	}
}

// Shifts the positions of the coordinates until the coordinate
// at indexOfFirstCoordinate is first.
//
// If ensureRing is true, first and last
// coordinate of the returned array are equal.
func scroll(coordinates []Coordinate, indexOfFirstCoordinate int, ensureRing bool) {
	i := indexOfFirstCoordinate
	if i <= 0 {
		return
	}
	l := len(coordinates)
	result := make([]Coordinate, l, l)
	if !ensureRing {
		copy(result[:l-i], coordinates[i:l])
		copy(result[l-i:l], coordinates[0:i])
	} else {
		last := l - 1
		// fill values
		var j int
		for j = 0; j < last; j++ {
			result[j] = coordinates[(i+j)%last]
		}
		// Fix the ring (first == last)
		result[j] = result[0].clone()
	}
	copy(coordinates, result)
}

func scrollWithRingCheck(coordinates []Coordinate, indexOfFirstCoordinate int) {
	scroll(coordinates, indexOfFirstCoordinate, isRing(coordinates))
}

// Returns the index of Coordinate in Coordinates slice.
// The first position is 0; the second, 1; etc.
func indexOf(coordinate Coordinate, coordinates []Coordinate) int {
	for i, c := range coordinates {
		if coordinate.equals(&c) {
			return i
		}
	}
	return -1
}

// Extracts a subsequence of the input Coordinate array
// from indices start to end (inclusive).
// The input indices are clamped to the array size;
// If the end index is less than the start index,
// the extracted array will be empty.
func extract(pts []Coordinate, start, end int) []Coordinate {
	start = clamp(start, 0, len(pts))
	end = clamp(end, -1, len(pts))
	npts := end - start + 1
	if end < 0 {
		npts = 0
	}
	if start >= len(pts) {
		npts = 0
	}
	if end < start {
		npts = 0
	}
	var extractPts []Coordinate
	if npts == 0 {
		return extractPts
	}
	return pts[start:end]
}

// Computes the Envelope of the coordinates.
func envelope(coordinates []Coordinate) Envelope {
	env := NewEmptyEnvelope()
	for _, c := range coordinates {
		env.expandToIncludeCoordinate(c)
	}
	return env
}

// Extracts the coordinates which intersect an Envelope
func intersection(coordinates []Coordinate, env Envelope) []Coordinate {
	result := make([]Coordinate, 0, len(coordinates))
	for _, c := range coordinates {
		if env.intersectsCoordinate(c) {
			result = append(result, c)
		}
	}
	return result
}

func equalArrays(coord1, coord2 []Coordinate) bool {
	if (coord1 == nil || len(coord1) == 0) && (coord1 == nil || len(coord2) == 0) {
		return true
	}
	if len(coord1) != len(coord2) {
		return false
	}
	for i, _ := range coord1 {
		if !coord1[i].equals(&coord2[i]) {
			return false
		}
	}
	return true
}
