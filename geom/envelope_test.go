package geom

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestEverything(t *testing.T) {
	assert := assert2.New(t)
	e1 := NewEmptyEnvelope()
	assert.True(e1.isNull())
	assert.Equal(0.0, e1.width())
	assert.Equal(0.0, e1.height())
	e1.expandToInclude(100, 101)
	e1.expandToInclude(200, 202)
	e1.expandToInclude(150, 151)
	assert.Equal(200.0, e1.maxX)
	assert.Equal(202.0, e1.maxY)
	assert.Equal(100.0, e1.minX)
	assert.Equal(101.0, e1.minY)
	assert.True(e1.contains(120, 120))
	assert.True(e1.contains(120, 101))
	assert.False(e1.contains(120, 100))
	assert.Equal(101.0, e1.height())
	assert.Equal(100.0, e1.width())
	assert.False(e1.isNull())

	e2 := NewEnvelope(499, 500, 500, 501)
	assert.False(e1.containsEnvelope(e2))
	assert.False(e1.intersectsEnvelope(e2))
	e1.expandToIncludeEnvelope(e2)
	assert.True(e1.containsEnvelope(e2))
	assert.True(e1.intersectsEnvelope(e2))
	assert.Equal(500.0, e1.maxX)
	assert.Equal(501.0, e1.maxY)
	assert.Equal(100.0, e1.minX)
	assert.Equal(101.0, e1.minY)

	e3 := NewEnvelope(300, 700, 300, 700)
	assert.False(e1.containsEnvelope(e3))
	assert.True(e1.intersectsEnvelope(e3))

	e4 := NewEnvelope(300, 301, 300, 301)
	assert.True(e1.containsEnvelope(e4))
	assert.True(e1.intersectsEnvelope(e4))
}

func TestIntersect(t *testing.T) {
	checkIntersectsPermuted(1, 1, 2, 2, 2, 2, 3, 3, true, t)
	checkIntersectsPermuted(1, 1, 2, 2, 3, 3, 4, 4, false, t)
}

func TestIntersectsEmpty(t *testing.T) {
	assert := assert2.New(t)
	a := NewEnvelope(-5, 5, -5, 5)
	b := NewEnvelope(100, 101, 100, 101)
	empty := NewEmptyEnvelope()
	assert.False(a.intersectsEnvelope(empty))
	assert.False(empty.intersectsEnvelope(a))
	assert.False(empty.intersectsEnvelope(b))
	assert.False(b.intersectsEnvelope(empty))
}

func TestDisjointEmpty(t *testing.T) {
	assert := assert2.New(t)
	a := NewEnvelope(-5, 5, -5, 5)
	b := NewEnvelope(100, 101, 100, 101)
	empty := NewEmptyEnvelope()
	assert.True(a.disjoint(empty))
	assert.True(empty.disjoint(a))
	assert.True(empty.disjoint(b))
	assert.True(b.disjoint(empty))
}

func TestContainsEmpty(t *testing.T) {
	assert := assert2.New(t)
	a := NewEnvelope(-5, 5, -5, 5)
	b := NewEnvelope(100, 101, 100, 101)
	empty := NewEmptyEnvelope()
	assert.False(a.containsEnvelope(empty))
	assert.False(empty.containsEnvelope(a))
	assert.False(empty.containsEnvelope(b))
	assert.False(b.containsEnvelope(empty))
}

func TestExpandToIncludeEmpty(t *testing.T) {
	assert := assert2.New(t)
	a := NewEnvelope(-5, 5, -5, 5)
	b := NewEnvelope(-5, 5, -5, 5)
	c := NewEnvelope(100, 101, 100, 101)
	d := NewEnvelope(100, 101, 100, 101)
	empty := NewEmptyEnvelope()
	a.expandToIncludeEnvelope(empty)
	assert.True(b == a)
	empty.expandToIncludeEnvelope(a)
	assert.True(b == a)
	empty = NewEmptyEnvelope()
	c.expandToIncludeEnvelope(empty)
	assert.True(d == c)
	empty.expandToIncludeEnvelope(c)
	assert.True(d == c)
}

func TestEmpty(t *testing.T) {
	assert := assert2.New(t)
	empty := NewEmptyEnvelope()
	assert.Equal(0.0, empty.height())
	assert.Equal(0.0, empty.width())
	assert.Equal(NewEmptyEnvelope(), NewEmptyEnvelope())
	e := NewEnvelope(100, 101, 100, 101)
	e.init(NewEmptyEnvelope())
	assert.Equal(NewEmptyEnvelope(), e)
}

//TODO: add geometry test

func TestSetToNull(t *testing.T) {
	assert := assert2.New(t)
	e1 := NewEmptyEnvelope()
	assert.True(e1.isNull())
	e1.expandToInclude(5, 5)
	assert.False(e1.isNull())
	e1.setToNull()
	assert.True(e1.isNull())
}

func TestEquals(t *testing.T) {
	assert := assert2.New(t)
	e1 := NewEnvelope(1, 2, 3, 4)
	e2 := NewEnvelope(1, 2, 3, 4)
	assert.Equal(e1, e2)
	// TODO: add hashcode comparison
	e3 := NewEnvelope(1, 2, 3, 5)
	assert.False(e1 == e3)
	// TODO: add hashcode comparison
	e1.setToNull()
	assert.False(e1 == e2)
	e2.setToNull()
	assert.Equal(e1, e2)
	// TODO: add hashcode comparison
}

func TestEquals2(t *testing.T) {
	assert := assert2.New(t)
	e1 := NewEmptyEnvelope()
	e2 := NewEmptyEnvelope()
	assert.Equal(e1, e2)
	// skipped one test as it is covered by previous method,
	// because we cannot do NewEmptyEnvelope as all methods are
	// on reference, so this would be not idiomatic
	e1 = NewEnvelope(1, 2, 1.5, 2)
	e2 = NewEnvelope(1, 2, 2, 2)
	assert.NotEqual(e1, e2)
}

func TestCopyConstructor(t *testing.T) {
	assert := assert2.New(t)
	e1 := NewEnvelope(1, 2, 3, 4)
	e2 := CopyEnvelope(e1)
	assert.Equal(1.0, e2.minX)
	assert.Equal(2.0, e2.maxX)
	assert.Equal(3.0, e2.minY)
	assert.Equal(4.0, e2.maxY)
}

func TestCopy(t *testing.T) {

	assert := assert2.New(t)
	e1 := NewEnvelope(1, 2, 3, 4)
	e2 := e1.copy()
	assert.Equal(1.0, e2.minX)
	assert.Equal(2.0, e2.maxX)
	assert.Equal(3.0, e2.minY)
	assert.Equal(4.0, e2.maxY)
	eNull := NewEmptyEnvelope()
	eNullCopy := eNull.copy()
	assert.True(eNullCopy.isNull())
}

func TestMetrics(t *testing.T) {
	assert := assert2.New(t)
	env := NewEnvelope(0, 4, 0, 3)
	assert.Equal(env.width(), 4.0)
	assert.Equal(env.height(), 3.0)
	assert.Equal(env.diameter(), 5.0)
}

func TestEmptyMetrics(t *testing.T) {
	assert := assert2.New(t)
	env := NewEmptyEnvelope()
	assert.Equal(env.width(), 0.0)
	assert.Equal(env.height(), 0.0)
	assert.Equal(env.diameter(), 0.0)
}

func TestCompareTo(t *testing.T) {
	a := assert2.New(t)
	checkCompareTo(a, 0, NewEmptyEnvelope(), NewEmptyEnvelope())
	checkCompareTo(a, 0, NewEnvelope(1, 2, 1, 2), NewEnvelope(1, 2, 1, 2))
	checkCompareTo(a, 1, NewEnvelope(2, 3, 1, 2), NewEnvelope(1, 2, 1, 2))
	checkCompareTo(a, -1, NewEnvelope(1, 2, 1, 2), NewEnvelope(2, 3, 1, 2))
	checkCompareTo(a, 1, NewEnvelope(1, 2, 1, 3), NewEnvelope(1, 2, 1, 2))
	checkCompareTo(a, 1, NewEnvelope(2, 3, 1, 3), NewEnvelope(1, 3, 1, 2))
}

func checkCompareTo(assert *assert2.Assertions, expected int, env1, env2 Envelope) {
	assert.True(expected == env1.compareTo(env2))
	assert.True(-expected == env2.compareTo(env1))
}

func checkIntersectsPermuted(a1x, a1y, a2x, a2y, b1x, b1y, b2x, b2y float64, expected bool, t *testing.T) {
	checkIntersects(a1x, a1y, a2x, a2y, b1x, b1y, b2x, b2y, expected, t)
	checkIntersects(a1x, a2y, a2x, a1y, b1x, b1y, b2x, b2y, expected, t)
	checkIntersects(a1x, a1y, a2x, a2y, b1x, b2y, b2x, b1y, expected, t)
	checkIntersects(a1x, a2y, a2x, a1y, b1x, b2y, b2x, b1y, expected, t)
}

func checkIntersects(a1x, a1y, a2x, a2y, b1x, b1y, b2x, b2y float64, expected bool, t *testing.T) {
	assert := assert2.New(t)
	a := NewEnvelope(a1x, a2x, a1y, a2y)
	b := NewEnvelope(b1x, b2x, b1y, b2y)
	assert.Equal(expected, a.intersectsEnvelope(b))
	assert.NotEqual(expected, a.disjoint(b))

	a1 := NewXYCoordinate(a1x, a1y)
	a2 := NewXYCoordinate(a2x, a2y)
	b1 := NewXYCoordinate(b1x, b1y)
	b2 := NewXYCoordinate(b2x, b2y)
	assert.Equal(expected, intersect(a1, a2, b1, b2))
	assert.Equal(expected, a.intersectsExtent(b1, b2))
}
