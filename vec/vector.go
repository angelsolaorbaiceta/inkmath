package vec

import (
	"errors"

	"github.com/angelsolaorbaiceta/inkmath"
)

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
func (v Vector) Plus(other *Vector) (*Vector, error) {
	return operateWithVectors(&v, other, func(a float64, b float64) float64 {
		return a + b
	})
}

// Minus subtracts two vectors.
func (v Vector) Minus(other *Vector) (*Vector, error) {
	return operateWithVectors(&v, other, func(a float64, b float64) float64 {
		return a - b
	})
}

func operateWithVectors(u, v *Vector, operation func(float64, float64) float64) (*Vector, error) {
	if u.length != v.length {
		return nil, errors.New("Cannot operate with vectors of different sizes")
	}

	result := Make(u.length)
	for i := 0; i < u.length; i++ {
		result.data[i] = operation(u.data[i], v.data[i])
	}

	return result, nil
}

// Times multiplies two vectors as v' · other.
func (v Vector) Times(other *Vector) (float64, error) {
	if v.length != other.length {
		return 0.0, errors.New("Cannot operate with vectors of different sizes")
	}

	result := 0.0
	for i := 0; i < v.length; i++ {
		result += v.data[i] * other.data[i]
	}

	return result, nil
}
