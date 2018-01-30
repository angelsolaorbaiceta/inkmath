package vec

import "github.com/angelsolaorbaiceta/inkmath"

// Vector is an array of values.
type Vector struct {
	length int
	data   []float64
}

/* ::::::::::::::: Construction ::::::::::::::: */

// Make returns a vector with the given size all filled with zeroes.
func Make(size int) *Vector {
	return &Vector{size, make([]float64, size)}
}

// MakeWithValues returns a vector with the given values.
func MakeWithValues(vals []float64) *Vector {
	v := Make(len(vals))
	for i := 0; i < len(vals); i++ {
		v.data[i] = vals[i]
	}

	return v
}

/* ::::::::::::::: Properties ::::::::::::::: */

// Length is the size of the vector.
func (v Vector) Length() int {
	return v.length
}

/* ::::::::::::::: Methods ::::::::::::::: */

// SetValue sets the given value at the given index.
func (v *Vector) SetValue(i int, value float64) {
	v.data[i] = value
}

// Value returns the value at the given index.
func (v Vector) Value(i int) float64 {
	return v.data[i]
}

// Equals compares two vectors and returns true if they contain the same elements.
func (v Vector) Equals(other *Vector) bool {
	if v.length != other.length {
		return false
	}

	for i := 0; i < v.length; i++ {
		if !inkmath.FuzzyEqual(v.data[i], other.data[i]) {
			return false
		}
	}

	return true
}

/* ::::::::::::::: Operations ::::::::::::::: */

// Plus adds two vectors.
func (v Vector) Plus(other *Vector) *Vector {
	if v.length != other.length {
		panic("Cannot sum vectors of different sizes")
	}

	sum := Make(v.length)
	for i := 0; i < v.length; i++ {
		sum.data[i] = v.data[i] + other.data[i]
	}

	return sum
}
