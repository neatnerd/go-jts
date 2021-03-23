package geom

import (
	"fmt"
	"math"
)

// Defines a rectangular region of the 2D coordinate plane.
// It is often used to represent the bounding box of a Geometry
// e.g. the minimum and maximum x and y values of the Coordinate
type Envelope struct {
	minX float64
	maxX float64
	minY float64
	maxY float64
}

// Test the point q to see whether it intersects the Envelope defined by p1-p2
func intersectPoint(p1, p2, q *Coordinate) bool {
	return (q.x >= math.Min(p1.x, p2.x) && q.x <= math.Max(p1.x, p2.x)) &&
		(q.y >= math.Min(p1.y, p2.y) && q.y <= math.Max(p1.y, p2.y))
}

// Tests whether the envelope defined by p1-p2
// and the envelope defined by q1-q2 intersect.
func intersect(p1, p2, q1, q2 Coordinate) bool {
	minQ := math.Min(q1.x, q2.x)
	maxQ := math.Max(q1.x, q2.x)
	minP := math.Min(p1.x, p2.x)
	maxP := math.Max(p1.x, p2.x)

	if minP > maxQ {
		return false
	}
	if maxP < minQ {
		return false
	}

	minQ = math.Min(q1.y, q2.y)
	maxQ = math.Max(q1.y, q2.y)
	minP = math.Min(p1.y, p2.y)
	maxP = math.Max(p1.y, p2.y)

	if minP > maxQ {
		return false
	}
	if maxP < minQ {
		return false
	}
	return true
}

// Initialize an Envelope for a region defined by maximum and minimum values.
func NewEnvelope(x1, x2, y1, y2 float64) Envelope {
	result := Envelope{}
	if x1 < x2 {
		result.minX = x1
		result.maxX = x2
	} else {
		result.minX = x2
		result.maxX = x1
	}
	if y1 < y2 {
		result.minY = y1
		result.maxY = y2
	} else {
		result.minY = y2
		result.maxY = y1
	}
	return result
}

// Creates an Envelope for a region defined by two Coordinate(s).
func NewEnvelopeFromCoordinates(p1, p2 Coordinate) Envelope {
	return NewEnvelope(p1.x, p2.x, p1.y, p2.y)
}

// Creates an Envelope for a region defined by a single Coordinate.
func NewPointEnvelope(p Coordinate) Envelope {
	return NewEnvelope(p.x, p.x, p.y, p.y)
}

// Makes this Envelope a "null" envelope, that is, the envelope
// of the empty geometry.
func NewEmptyEnvelope() Envelope {
	return Envelope{
		minX: 0,
		maxX: -1,
		minY: 0,
		maxY: -1,
	}
}

// Copy constructor
func CopyEnvelope(other Envelope) Envelope {
	return Envelope{
		minX: other.minX,
		maxX: other.maxX,
		minY: other.minY,
		maxY: other.maxY,
	}
}

// Makes this Envelope a "null" envelope, that is, the envelope
// of the empty geometry.
func (e *Envelope) setToNull() {
	e.minX = 0
	e.maxX = -1
	e.minY = 0
	e.maxY = -1
}

// Returns true, if this Envelope is a "null" Envelope
func (e *Envelope) isNull() bool {
	return e.maxX < e.minX
}

// Initialize an Envelope from an existing Envelope.
func (e *Envelope) init(env Envelope) {
	e.minX = env.minX
	e.maxX = env.maxX
	e.minY = env.minY
	e.maxY = env.maxY
}

// Returns the difference between the maximum and minimum x values.
func (e *Envelope) width() float64 {
	if e.isNull() {
		return 0
	}
	return e.maxX - e.minX
}

// Returns the difference between the maximum and minimum y values.
func (e *Envelope) height() float64 {
	if e.isNull() {
		return 0
	}
	return e.maxY - e.minY
}

// Returns the Envelope minimum x-value. min x > max x
// indicates that this is a null Envelope.
func (e *Envelope) diameter() float64 {
	if e.isNull() {
		return 0
	}
	w := e.width()
	h := e.height()
	return math.Sqrt(w*w + h*h)
}

// Gets the area of this envelope.
func (e *Envelope) area() float64 {
	return e.width() * e.height()
}

// Gets the minimum extent of this Envelope across both dimensions.
func (e *Envelope) minExtent() float64 {
	if e.isNull() {
		return 0
	}
	w := e.width()
	h := e.height()
	if w < h {
		return w
	}
	return h
}

// Gets the maximum extent of this Envelope across both dimensions.
func (e *Envelope) maxExtent() float64 {
	if e.isNull() {
		return 0
	}
	w := e.width()
	h := e.height()
	if w > h {
		return w
	}
	return h
}

// Enlarges this Envelope so that it contains the given point.
//  Has no effect if the point is already on or within the Envelope.
func (e *Envelope) expandToInclude(x, y float64) {
	if e.isNull() {
		e.minX = x
		e.maxX = x
		e.minY = y
		e.maxY = y
	} else {
		if x < e.minX {
			e.minX = x
		}
		if x > e.maxX {
			e.maxX = x
		}
		if y < e.minY {
			e.minY = y
		}
		if y > e.maxY {
			e.maxY = y
		}
	}
}

// Enlarges this Envelope so that it contains
// the given Coordinate
// Has no effect if the point is already on or within the envelope.
func (e *Envelope) expandToIncludeCoordinate(c Coordinate) {
	e.expandToInclude(c.x, c.y)
}

// Expands this envelope by a given distance in all directions.
// Both positive and negative distances are supported.
func (e *Envelope) expandBy(deltaX, deltaY float64) {
	if e.isNull() {
		return
	}
	e.minX -= deltaX
	e.maxX += deltaX
	e.minY -= deltaY
	e.maxY += deltaY
	if e.minX > e.maxX || e.minY > e.maxY {
		e.setToNull()
	}
}

// Enlarges this Envelope so that it contains the other Envelope.
// Has no effect if other is wholly on or within the envelope.
func (e *Envelope) expandToIncludeEnvelope(other Envelope) {
	if other.isNull() {
		return
	}
	if e.isNull() {
		e.minX = other.minX
		e.maxX = other.maxX
		e.minY = other.minY
		e.maxY = other.maxY
	} else {
		if other.minX < e.minX {
			e.minX = other.minX
		}
		if other.maxX > e.maxX {
			e.maxX = other.maxX
		}
		if other.minY < e.minY {
			e.minY = other.minY
		}
		if other.maxY > e.maxY {
			e.maxY = other.maxY
		}
	}
}

// Translates this envelope by given amounts in the X and Y direction.
func (e *Envelope) translate(transX, transY float64) {
	if e.isNull() {
		return
	}
	e.minX += transX
	e.maxX += transX
	e.minY += transY
	e.maxY += transY
}

// Computes the Coordinate of the centre of this Envelope (as long as it is non-null)
func (e *Envelope) centre() *Coordinate {
	if e.isNull() {
		return nil
	}
	result := NewXYCoordinate(
		(e.minX+e.maxX)/2,
		(e.minY+e.maxY)/2,
	)
	return &result
}

// Computes the intersection of two Envelope(s).
func (e *Envelope) intersection(env Envelope) Envelope {
	if e.isNull() || env.isNull() || !e.intersectsEnvelope(env) {
		return NewEmptyEnvelope()
	}
	intMinX := math.Max(e.minX, env.minX)
	intMinY := math.Max(e.minY, env.minY)
	intMaxX := math.Min(e.maxX, env.maxX)
	intMaxY := math.Min(e.maxY, env.maxY)
	return NewEnvelope(intMinX, intMaxX, intMinY, intMaxY)
}

// Tests if the region defined by other Envelope
// intersects the region of this Envelope.
func (e *Envelope) intersectsEnvelope(other Envelope) bool {
	if e.isNull() || other.isNull() {
		return false
	}
	return !(other.minX > e.maxX ||
		other.maxX < e.minX ||
		other.minY > e.maxY ||
		other.maxY < e.minY)
}

// Tests if the extent defined by two extremal Coordinates
// intersects the extent of this Envelope.
func (e *Envelope) intersectsExtent(a, b Coordinate) bool {
	if e.isNull() {
		return false
	}
	envMinX := math.Min(a.x, b.x)
	if envMinX > e.maxX {
		return false
	}
	envMaxX := math.Max(a.x, b.x)
	if envMaxX < e.minX {
		return false
	}
	envMinY := math.Min(a.y, b.y)
	if envMinY < e.maxY {
		return false
	}
	envMaxY := math.Max(a.y, b.y)
	if envMaxY < e.minY {
		return false
	}
	return true
}

// Tests if the region defined by other Envelope is
// disjoint from the region of this Envelope.
func (e *Envelope) disjoint(other Envelope) bool {
	if e.isNull() || other.isNull() {
		return true
	}
	return other.minX > e.maxX ||
		other.maxX < e.minX ||
		other.minY > e.maxY ||
		other.maxY < e.minY
}

// Check if the point (x,y) intersects (lies inside)
// the region of this Envelope
func (e *Envelope) intersects(x, y float64) bool {
	if e.isNull() {
		return false
	}
	return !(x > e.maxX ||
		x < e.minX ||
		y > e.maxY ||
		y < e.minY)
}

// Check if the Coordinate intersects (lies inside)
// the region of this Envelope
func (e *Envelope) intersectsCoordinate(p Coordinate) bool {
	return e.intersects(p.x, p.y)
}

// Tests if the given point lies in or on the envelope.
// Returns true if (x, y) lies in the interior or
// on the boundary of this Envelope.
func (e *Envelope) covers(x, y float64) bool {
	if e.isNull() {
		return false
	}
	return x >= e.minX &&
		x <= e.maxX &&
		y >= e.minY &&
		y <= e.maxY
}

// Tests if the given point lies in or on the Envelope.
func (e *Envelope) coversCoordinate(p Coordinate) bool {
	return e.covers(p.x, p.y)
}

// Tests if the other Envelope
// lies wholly inside this Envelope (inclusive of the boundary).
func (e *Envelope) coversEnvelope(other Envelope) bool {
	if e.isNull() || other.isNull() {
		return false
	}
	return other.minX >= e.minX &&
		other.maxX <= e.maxX &&
		other.minY >= e.minY &&
		other.maxY <= e.maxY
}

// Tests if the given point lies in or on the Envelope.
//
// Note that this is not the same definition as the SFS contains,
// which would exclude the envelope boundary.
func (e *Envelope) contains(x, y float64) bool {
	return e.covers(x, y)
}

// Tests if the given point lies in or on the envelope.
//
// Note that this is not the same definition as the SFS contains,
// which would exclude the envelope boundary.
func (e *Envelope) containsCoordinate(p Coordinate) bool {
	return e.contains(p.x, p.y)
}

// Tests if the given point lies in or on the Envelope.
//
// Note that this is not the same definition as the SFS contains,
// which would exclude the envelope boundary.
func (e *Envelope) containsEnvelope(other Envelope) bool {
	return e.coversEnvelope(other)
}

// Computes the distance between this and another Envelope.
// The distance between overlapping Envelopes is 0.  Otherwise, the
// distance is the Euclidean distance between the closest points.
func (e *Envelope) distance(env Envelope) float64 {
	if e.intersectsEnvelope(env) {
		return 0
	}
	dx := 0.0
	if e.maxX < env.minX {
		dx = env.minX - e.maxX
	} else if e.minX > env.maxX {
		dx = e.minX - env.maxX
	}
	dy := 0.0
	if e.maxY < env.minY {
		dy = env.minY - e.maxY
	} else if e.minY > env.maxY {
		dy = e.minY - env.maxY
	}
	if dx == 0 {
		return dy
	}
	if dy == 0 {
		return dy
	}
	return math.Sqrt(dx*dx + dy*dy)
}

func (e *Envelope) String() string {
	return fmt.Sprintf("Env[%g:%g,%g:%g]",
		e.minX, e.maxX, e.minY, e.maxY)
}

// Compares two envelopes using lexicographic ordering.
// The ordering comparison is based on the usual numerical
// comparison between the sequence of ordinates.
// Null envelopes are less than all non-null envelopes.
func (e *Envelope) compareTo(other Envelope) int {
	if e.isNull() {
		if other.isNull() {
			return 0
		}
		return -1
	} else {
		if other.isNull() {
			return 1
		}
	}
	if e.minX < other.minX {
		return -1
	}
	if e.minX > other.minX {
		return 1
	}
	if e.minY < other.minY {
		return -1
	}
	if e.minY > other.minY {
		return 1
	}
	if e.maxX < other.maxX {
		return -1
	}
	if e.maxX > other.maxX {
		return 1
	}
	if e.maxY < other.maxY {
		return -1
	}
	if e.maxY > other.maxY {
		return 1
	}
	return 0
}

func (e *Envelope) copy() Envelope {
	return Envelope{
		minX: e.minX,
		maxX: e.maxX,
		minY: e.minY,
		maxY: e.maxY,
	}
}
