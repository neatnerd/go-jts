package geom

import "math"

type ModelType string

const (
	FIXED           ModelType = "FIXED"
	FLOATING                  = "FLOATING"
	FLOATING_SINGLE           = "FLOATING_SINGLE"
)

// Specifies the precision model of the Coordinate(s) in a Geometry.
// In other words, specifies the grid of allowable
// points for all Geometry(s)
// The makePrecise method allows rounding a coordinate to
// a "precise" value; that is, one whose
//  precision is known exactly.
//
// Coordinates are assumed to be precise in geometries.
// That is, the coordinates are assumed to be rounded to the
// precision model given for the geometry.
// JTS input routines automatically round coordinates to the precision model
// before creating Geometries.
// All internal operations
// assume that coordinates are rounded to the precision model.
// Constructive methods (such as boolean operations) always round computed
// coordinates to the appropriate precision model.
//
// Currently three types of precision model are supported:
//
//
// 		- FLOATING - represents full double precision floating point. This is the default precision model used in JTS
// 		- FLOATING_SINGLE - represents single precision floating point.
// 		- FIXED - represents a model with a fixed number of decimal places.
//
// 	A Fixed Precision Model is specified by a scale factor.
//  The scale factor specifies the size of the grid which numbers are rounded to.
//  Input coordinates are mapped to fixed coordinates according to the following
//  equations:
//
//      jtsPt.x = round( (inputPt.x // scale ) / scale
//      jtsPt.y = round( (inputPt.y // scale ) / scale
//
//
// For example, to specify 3 decimal places of precision, use a scale factor
// of 1000. To specify -3 decimal places of precision (i.e. rounding to
// the nearest 1000), use a scale factor of 0.001.
//
// Coordinates are represented internally as Golang float64 values.
// Since Java uses the IEEE 754 floating point standard, this
// provides 53 bits of precision. (Thus the maximum precisely representable
// <i>integer</i> is 9,007,199,254,740,992 - or almost 16 decimal digits of precision).
// <p>
// JTS binary methods currently do not handle inputs which have different precision models.
// The precision model of any constructed geometric value is undefined.
type PrecisionModel struct {
	modelType ModelType
	scale     float64
}

// Determines which of two PrecisionModel is the most precise
// (allows the greatest number of significant digits).
func mostPrecise(pm1, pm2 *PrecisionModel) *PrecisionModel {
	if pm1.compareTo(pm2) >= 0 {
		return pm1
	}
	return pm2
}

// Creates a PrecisionModel with a default precision
// of FLOATING.
func NewDefaultPrecisionModel() PrecisionModel {
	return PrecisionModel{
		modelType: FLOATING,
	}
}

// Creates a PrecisionModel that specifies
// an explicit precision model type.
// If the model type is FIXED the scale factor will default to 1.
func NewPrecisionModel(modelType ModelType) PrecisionModel {
	result := PrecisionModel{
		modelType: modelType,
	}
	if modelType == FIXED {
		result.setScale(1.0)
	}
	return result
}

//  Creates a PrecisionModel that specifies Fixed precision.
//  Fixed-precision coordinates are represented as precise internal coordinates,
//  which are rounded to the grid defined by the scale factor.
func NewFixedPrecisionModel(scale float64) PrecisionModel {
	result := PrecisionModel{
		modelType: FIXED,
	}
	result.setScale(scale)
	return result
}

//  Sets the multiplying factor used to obtain a precise coordinate.
func (p *PrecisionModel) setScale(scale float64) {
	p.scale = math.Abs(scale)
}

// Tests whether the precision model supports floating point
// true if the precision model supports floating point
func (p *PrecisionModel) isFloating() bool {
	return p.modelType == FLOATING || p.modelType == FLOATING_SINGLE
}

// Returns the maximum number of significant digits provided by this
// precision model.
// Intended for use by routines which need to print out
// decimal representations of precise values (such as WKTWriter).
//
// This method would be more correctly called
// getMinimumDecimalPlaces,
// since it actually computes the number of decimal places
// that is required to correctly display the full
// precision of an ordinate value.
//
// Since it is difficult to compute the required number of
// decimal places for scale factors which are not powers of 10,
// the algorithm uses a very rough approximation in this case.
// This has the side effect that for scale factors which are
// powers of 10 the value returned is 1 greater than the true value.
func (p *PrecisionModel) getMaximumSignificantDigits() int {
	maxSigDigits := 16
	if p.modelType == FLOATING {
		maxSigDigits = 16
	} else if p.modelType == FLOATING_SINGLE {
		maxSigDigits = 6
	} else if p.modelType == FIXED {
		maxSigDigits = 1 + int(math.Ceil(math.Log(p.scale)/math.Log(10)))
	}
	return maxSigDigits
}

// Rounds a numeric value to the PrecisionModel grid.
// Asymmetric Arithmetic Rounding is used, to provide
// uniform rounding behaviour no matter where the number is
// on the number line.
//
// This method has no effect on NaN values.
func (p *PrecisionModel) makePrecise(value float64) float64 {
	if math.IsNaN(value) {
		return value
	}
	if p.modelType == FLOATING_SINGLE {
		floatSingleValue := float32(value)
		return float64(floatSingleValue)
	}
	if p.modelType == FIXED {
		return math.Round(value*p.scale) / p.scale
	}
	// modelType == FLOATING - no rounding necessary
	return value
}

// Rounds a Coordinate to the PrecisionModel grid.
// Modifies the Coordinate
func (p *PrecisionModel) makePreciseCoordinate(coordinate *Coordinate) {
	if p.modelType == FLOATING {
		return
	}

	coordinate.x = p.makePrecise(coordinate.x)
	coordinate.y = p.makePrecise(coordinate.y)

	//MD says it's OK that we're not makePrecise'ing the z [Jon Aquino]
}

// Compares this PrecisionModel object with the specified object for order.
// A PrecisionModel is greater than another if it provides greater precision.
// The comparison is based on the value returned by the
// getMaximumSignificantDigits method.
// This comparison is not strictly accurate when comparing floating precision models
// to fixed models; however, it is correct when both models are either floating or fixed.
func (p *PrecisionModel) compareTo(other *PrecisionModel) int {
	sigDigits := p.getMaximumSignificantDigits()
	otherSigDigits := other.getMaximumSignificantDigits()
	if sigDigits > otherSigDigits {
		return 1
	}
	if sigDigits < otherSigDigits {
		return -1
	}
	return 0
}
