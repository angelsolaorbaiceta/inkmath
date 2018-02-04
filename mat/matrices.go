package mat

import (
	"math"

	"github.com/angelsolaorbaiceta/inkmath"
)

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
		panic("Matrix symeetry only applies to square matrices")
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

/*
IsRowDominant returns true if for every row in the matrix, the element in the main diagonal
is greater than every other element.
*/
func IsRowDominant(m Matrixable) bool {
	if !IsSquare(m) {
		panic("Matrix dominancy only applies to square matrices")
	}

	var diagonalValue float64
	for i := 0; i < m.Rows(); i++ {
		diagonalValue = math.Abs(m.Value(i, i))
		for j := 0; j < m.Cols(); j++ {
			if i != j && diagonalValue < math.Abs(m.Value(i, j)) {
				return false
			}
		}
	}

	return true
}
