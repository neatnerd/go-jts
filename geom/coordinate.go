package geom

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Coordinate struct {
	x          float64
	y          float64
	z          float64
	dimensions int
}

const (
	// Standard ordinate index value for, where X is 0
	X = 0
	// Standard ordinate index value for, where Y is 1
	Y = 1
	// Standard ordinate index value for, where Z is 2
	Z = 2
)

// The value used to indicate a null or missing ordinate value.
// In particular, used for the value of ordinates for dimensions
// greater than the defined dimension of a coordinate.
var NullOrdinate = math.NaN()

func NewCoordinate(x, y, z float64) Coordinate {
	return Coordinate{
		x: x,
		y: y,
		z: z,
	}
}

// Constructs Coordinate at (0,0,NaN)
func NewEmptyCoordinate() Coordinate {
	return NewXYEmptyCoordinate()
}

// Constructs a Coordinate at (x,y,NaN).
func NewXYCoordinate(x, y float64) Coordinate {
	result := NewCoordinate(x, y, NullOrdinate)
	result.dimensions = 2
	return result
}

// Constructs a Coordinate at (0,0,NaN)
func NewXYEmptyCoordinate() Coordinate {
	return NewXYCoordinate(0, 0)
}

func (c *Coordinate) setCoordinate(other *Coordinate) {
	c.x = other.x
	c.y = other.y
	c.z = other.z
}

// Gets the ordinate value for the given index.
func (c *Coordinate) getOrdinate(ordinateIndex int) (error, float64) {
	switch ordinateIndex {
	case X:
		return nil, c.x
	case Y:
		return nil, c.y
	case Z:
		return nil, c.z
	}
	return errors.New("Invalid ordinate index: " + strconv.Itoa(ordinateIndex)), 0
}

// Sets the ordinate for the given index to a given value.
func (c *Coordinate) setOrdinate(ordinateIndex int, value float64) error {
	switch ordinateIndex {
	case X:
		c.x = value
		return nil
	case Y:
		c.y = value
		return nil
	case Z:
		c.z = value
		return nil
	}
	return errors.New("Invalid ordinate index: " + strconv.Itoa(ordinateIndex))
}

// Returns whether the planar projections of the two Coordinates
func (c *Coordinate) equals2D(other *Coordinate) bool {
	if c.x != other.x {
		return false
	}
	if c.y != other.y {
		return false
	}
	return true
}

// Tests if another Coordinate has the same values for the X and Y ordinates
// within a specified tolerance value.
// The Z ordinate is ignored.
func (c *Coordinate) equals2DWithTolerance(other *Coordinate, tolerance float64) bool {
	if !EqualsWithTolerance(c.x, other.x, tolerance) {
		return false
	}
	if !EqualsWithTolerance(c.y, other.y, tolerance) {
		return false
	}
	return true
}

// Tests if another coordinate has the same values for the X, Y and Z ordinates.
func (c *Coordinate) equals3D(other *Coordinate) bool {
	return (c.x == other.x) && (c.y == other.y) &&
		((c.z == other.z) || (math.IsNaN(c.z) && math.IsNaN(other.z)))
}

// Tests if another Coordinate has the same value for Z, within a tolerance.
func (c *Coordinate) equalInZ(other *Coordinate, tolerance float64) bool {
	return EqualsWithTolerance(c.z, other.z, tolerance)
}

// Returns true if other has the same values for
// the x and y ordinates.
// Since Coordinates are 2.5D, this routine ignores the z value when making the comparison.
func (c *Coordinate) equals(other *Coordinate) bool {
	return c.equals2D(other)
}

// Compares this Coordinate with the another specified Coordinate for order.
// This method ignores the z value when making the comparison.
// Note: This method assumes that ordinate values
// are valid numbers.  NaN values are not handled correctly.
func (c *Coordinate) compareTo(other Coordinate) int {
	if c.x < other.x {
		return -1
	}
	if c.x > other.x {
		return 1
	}
	if c.y < other.y {
		return -1
	}
	if c.y > other.y {
		return 1
	}
	return 0
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%g, %g, %g)", c.x, c.y, c.z)
}

// Computes the 2-dimensional Euclidean distance to another location.
// The Z-ordinate is ignored.
func (c *Coordinate) distance(other *Coordinate) float64 {
	dx := c.x - other.x
	dy := c.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

// Computes the 3-dimensional Euclidean distance to another location.
func (c *Coordinate) distance3D(other *Coordinate) float64 {
	dx := c.x - other.x
	dy := c.y - other.y
	dz := c.z - other.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (c *Coordinate) measure() int {
	return 0 // TODO: add support for M dimension
}

func (c *Coordinate) clone() Coordinate {
	return Coordinate{
		x:          c.x,
		y:          c.y,
		z:          c.z,
		dimensions: c.dimensions,
	}
}
