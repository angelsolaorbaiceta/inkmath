package vec

import "github.com/angelsolaorbaiceta/inkmath/nums"

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
