package geom

import (
	assert2 "github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var COORDS_1 = []Coordinate{NewXYCoordinate(1, 1), NewXYCoordinate(2, 2), NewXYCoordinate(3, 3)}
var COORDS_EMPTY []Coordinate

func TestPtNotInList1(t *testing.T) {
	expected := NewXYCoordinate(2, 2)
	assert2.True(t,
		ptNotInList([]Coordinate{NewXYCoordinate(1, 1), NewXYCoordinate(2, 2), NewXYCoordinate(3, 3)},
			[]Coordinate{NewXYCoordinate(1, 1), NewXYCoordinate(1, 2), NewXYCoordinate(1, 3)}).equals2D(&expected))

}

func TestPtNotInList2(t *testing.T) {
	assert2.True(t,
		ptNotInList([]Coordinate{NewXYCoordinate(1, 1), NewXYCoordinate(2, 2), NewXYCoordinate(3, 3)},
			[]Coordinate{NewXYCoordinate(1, 1), NewXYCoordinate(2, 2), NewXYCoordinate(3, 3)}) == nil)

}

func TestEnvelope1(t *testing.T) {
	assert2.Equal(t, envelope(COORDS_1), NewEnvelope(1, 3, 1, 3))
}

func TestEnvelopeEmpty(t *testing.T) {
	assert2.Equal(t, envelope(COORDS_EMPTY), NewEmptyEnvelope())
}

func TestIntersectionEnvelope1(t *testing.T) {
	assert2.True(t, equalArrays(
		intersection(COORDS_1, NewEnvelope(1, 2, 1, 2)),
		[]Coordinate{NewXYCoordinate(1, 1), NewXYCoordinate(2, 2)},
	))
}

func TestIntersection_envelopeDisjoint(t *testing.T) {
	assert2.True(t, equalArrays(
		intersection(COORDS_1, NewEnvelope(10, 20, 10, 20)),
		COORDS_EMPTY,
	))
}

func TestIntersectionEmptyEnvelope(t *testing.T) {
	assert2.True(t, equalArrays(
		intersection(COORDS_EMPTY, NewEnvelope(1, 2, 1, 2)),
		COORDS_EMPTY,
	))
}

func TestIntersectionCoordsEmptyEnvelope(t *testing.T) {
	assert2.True(t, equalArrays(
		intersection(COORDS_1, NewEmptyEnvelope()),
		COORDS_EMPTY,
	))
}

func TestEnforceConsistency(t *testing.T) {
	// TODO: implement enforceConsistency test
}

func TestScrollRing(t *testing.T) {
	sequence := createCircle(NewXYCoordinate(10, 10), 9.0)
	scrolled := createCircle(NewXYCoordinate(10, 10), 9.0)
	scrollWithRingCheck(scrolled, 12)
	io := 12
	for is := 0; is < len(scrolled)-1; is++ {
		checkCoordinateAt(sequence, io, scrolled, is, t)
		io++
		io %= len(scrolled) - 1
	}
	checkCoordinateAt(scrolled, 0, scrolled, len(scrolled)-1, t)
}

func TestScroll(t *testing.T) {
	// arrange
	sequence := createCircularString(NewXYCoordinate(20, 20), 7.0, 0.1, 22)
	scrolled := createCircularString(NewXYCoordinate(20, 20), 7.0, 0.1, 22)
	scrollWithRingCheck(scrolled, 12)
	io := 12
	for is := 0; is < len(scrolled)-1; is++ {
		checkCoordinateAt(sequence, io, scrolled, is, t)
		io++
		io %= len(scrolled)
	}
}

func checkCoordinateAt(seq1 []Coordinate, pos1 int, seq2 []Coordinate, pos2 int, t *testing.T) {
	c1 := seq1[pos1]
	c2 := seq2[pos2]
	assert := assert2.New(t)
	assert.Equal(c1.x, c2.x, "unexpected x-ordinate at post %d", pos2)
	assert.Equal(c1.y, c2.y, "unexpected y-ordinate at post %d", pos2)
}

func createCircularString(center Coordinate, radius, startAngle float64, numpoints int) []Coordinate {
	numSegmentsCircle := 48
	angleCircle := 2 * math.Pi
	angleStep := angleCircle / float64(numSegmentsCircle)
	sequence := make([]Coordinate, numpoints)
	pm := NewFixedPrecisionModel(1000)
	angle := startAngle
	for i := 0; i < numpoints; i++ {
		dx := math.Cos(angle) * radius
		dy := math.Sin(angle) * radius
		sequence[i] = NewXYCoordinate(pm.makePrecise(center.x+dx), pm.makePrecise(center.y+dy))
		angle += angleStep
		angle = math.Mod(angle, angleCircle)
	}
	return sequence
}

func createCircle(center Coordinate, radius float64) []Coordinate {
	// Get a complete circular string
	result := createCircularString(center, radius, 0.0, 49)

	// ensure it is closed
	result[48] = result[0].clone()

	return result
}
