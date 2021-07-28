package vec

import "github.com/angelsolaorbaiceta/inkmath/nums"

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

/*
VectorContainsData tests whether a given vector contains exactly the same data as the
slice of float64 numbers.

Both need to be of the same size in order for this test to be true.
*/
func VectorContainsData(vector *Vector, data []float64) bool {
	if vector.length != len(data) {
		return false
	}

	for i := 0; i < vector.length; i++ {
		if !nums.FuzzyEqual(vector.Value(i), data[i]) {
			return false
		}
	}

	return true
}
