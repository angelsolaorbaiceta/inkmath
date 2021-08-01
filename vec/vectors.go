package vec

import "github.com/angelsolaorbaiceta/inkmath/nums"

func makeVector(size int) *Vector {
	return &Vector{size, make([]float64, size)}
}

// Make returns a vector with the given size all filled with zeroes.
func Make(size int) MutableVector {
	return makeVector(size)
}

// MakeReadOnly returns a read-only vector with the given size all filled with zeroes.
func MakeReadOnly(size int) ReadOnlyVector {
	return makeVector(size)
}

// MakeWithValues returns a vector with the given values.
func MakeWithValues(vals []float64) MutableVector {
	var (
		size = len(vals)
		v    = makeVector(size)
	)

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
func VectorContainsData(vector ReadOnlyVector, data []float64) bool {
	if vector.Length() != len(data) {
		return false
	}

	for i := 0; i < vector.Length(); i++ {
		if !nums.FuzzyEqual(vector.Value(i), data[i]) {
			return false
		}
	}

	return true
}
