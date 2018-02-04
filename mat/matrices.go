package mat

import "github.com/angelsolaorbaiceta/inkmath"

/*
IsSquare returns true if the given matrix has the same number of rows and columns.
*/
func IsSquare(m Matrixable) bool {
	return m.Rows() == m.Cols()
}

/*
IsSymmetric returns true if the given matrix is square and equals to it's traspose.
*/
func IsSymmetric(m Matrixable) bool {
	if !IsSquare(m) {
		return false
	}

	for i := 0; i < m.Rows(); i++ {
		for j := i + 1; j < m.Cols(); j++ {
			if !inkmath.FuzzyEqual(m.Value(i, j), m.Value(j, i)) {
				return false
			}
		}
	}

	return true
}
