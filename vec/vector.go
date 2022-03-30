package vec

import (
	"math"

	"github.com/angelsolaorbaiceta/inkmath/nums"
)

// A ReadOnlyVector is a Vector whose operations never mutate the internal state.
type ReadOnlyVector interface {
	/* Properties */
	Length() int
	Norm() float64

	/* Methods */
	Value(i int) float64
	Opposite() ReadOnlyVector
	Scaled(factor float64) ReadOnlyVector
	Plus(other ReadOnlyVector) ReadOnlyVector
	Minus(other ReadOnlyVector) ReadOnlyVector
	Times(other ReadOnlyVector) float64

	Clone() ReadOnlyVector
	Equals(other ReadOnlyVector) bool
	AsMutable() MutableVector
}

// Vector is a one-dimension array of float64 values.
type Vector struct {
	length int
	data   []float64
}

// Length is the size of the vector.
func (v Vector) Length() int {
	return v.length
}

// Norm returns the L2-norm of the vector.
func (v Vector) Norm() float64 {
	norm := 0.0
	for _, val := range v.data {
		norm += val * val
	}

	return math.Sqrt(norm)
}

// Value returns the value at the given index.
func (v Vector) Value(i int) float64 {
	return v.data[i]
}

/*
Opposite creates a new vector which is the opposite of this one, that is,
points in the opposite direction.
*/
func (v Vector) Opposite() ReadOnlyVector {
	opposite := makeVector(v.length)
	for i, val := range v.data {
		opposite.data[i] = -val
	}

	return opposite
}

/*
Scaled creates a new vector consisting on the scaled projections of this vector.
*/
func (v Vector) Scaled(factor float64) ReadOnlyVector {
	scaled := makeVector(v.length)
	for i, val := range v.data {
		scaled.data[i] = val * factor
	}

	return scaled
}

// Plus adds two vectors.
func (v Vector) Plus(other ReadOnlyVector) ReadOnlyVector {
	return operateWithVectors(&v, other, func(a float64, b float64) float64 {
		return a + b
	})
}

// Minus subtracts two vectors.
func (v Vector) Minus(other ReadOnlyVector) ReadOnlyVector {
	return operateWithVectors(&v, other, func(a float64, b float64) float64 {
		return a - b
	})
}

func operateWithVectors(u, v ReadOnlyVector, operation func(float64, float64) float64) *Vector {
	if u.Length() != v.Length() {
		panic("Cannot operate with vectors of different sizes")
	}

	result := makeVector(u.Length())
	for i := 0; i < u.Length(); i++ {
		result.data[i] = operation(u.Value(i), v.Value(i))
	}

	return result
}

// Times multiplies two vectors as v' · other.
func (v Vector) Times(other ReadOnlyVector) float64 {
	if v.length != other.Length() {
		panic("Cannot operate with vectors of different sizes")
	}

	result := 0.0
	for i := 0; i < v.length; i++ {
		result += v.data[i] * other.Value(i)
	}

	return result
}

// Equals compares two vectors and returns true if they contain the same elements.
func (v Vector) Equals(other ReadOnlyVector) bool {
	if v.length != other.Length() {
		return false
	}

	for i := 0; i < v.length; i++ {
		if !nums.FloatsEqual(v.data[i], other.Value(i)) {
			return false
		}
	}

	return true
}

// Clone creates an exact copy of the vector.
func (v Vector) Clone() ReadOnlyVector {
	vec := makeVector(v.length)
	for i, val := range v.data {
		vec.data[i] = val
	}

	return vec
}

// As Mutable returns this vector typed as a mutable one.
func (v Vector) AsMutable() MutableVector {
	return &v
}
