package inkmath

import "math"

const DEFAULT_EPSILON = 1E-10

// Compares two float64 values and returns true if the difference between the two is
// smaller than a given epsilon.
func FuzzyEqualEps(a, b, epsilon float64) bool {
    return math.Abs(a - b) < epsilon
}

// Compares two float64 values and returns true if the difference between the two is
// smaller than a default epsilon value (1E-10).
func FuzzyEqual(a, b float64) bool {
    return FuzzyEqualEps(a, b, DEFAULT_EPSILON)
}

// Returns true if the given number is close enough to zero.
// Comparation made with the default epsilon value.
func IsCloseToZero(a float64) bool {
    return FuzzyEqual(a, 0.0)
}

// Computes the linear interpolation for a given position given two points on
// the desired line: (startPos, startVal) and (endPos, endVal).
func LinInterpol(startPos, startVal, endPos, endVal, posToInterpolate float64) float64 {
    return startVal + (posToInterpolate - startPos) * (endVal - startVal) / (endPos - startPos)
}
