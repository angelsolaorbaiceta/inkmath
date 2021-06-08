package mat

import (
	"math"

	"github.com/angelsolaorbaiceta/inkmath/nums"
)

// CholeskyDecomposition computes the Cholesky lower matrix for a square, symmetric matrix.
func CholeskyDecomposition(m ReadOnlyMatrix) ReadOnlyMatrix {
	if !IsSquare(m) {
		panic("Cannot use Cholesky factorization in non-square matrices")
	}

	var (
		size        = m.Rows()
		lowerMatrix = MakeSparse(size, size)
		sqSum, sum  float64
	)

	for i := 0; i < size; i++ {
		sqSum = 0.0
		for j := 0; j <= i; j++ {
			if i == j {
				lowerMatrix.SetValue(i, i, math.Sqrt(m.Value(i, i)-sqSum))
			} else {
				sum = 0.0
				for k := 0; k < j; k++ {
					sum += lowerMatrix.Value(i, k) * lowerMatrix.Value(j, k)
				}

				lowerMatrix.SetValue(i, j, (m.Value(i, j)-sum)/lowerMatrix.Value(j, j))
				sqSum += lowerMatrix.Value(i, j) * lowerMatrix.Value(i, j)
			}
		}
	}

	return lowerMatrix
}

/*
IncompleteCholeskyDecomposition computes the Incomplete Cholesky lower matrix
decomposition for the given square and symmetric matrix.
*/
func IncompleteCholeskyDecomposition(m ReadOnlyMatrix) ReadOnlyMatrix {
	if !IsSquare(m) {
		panic("Cannot use Cholesky factorization in non-square matrices")
	}

	var (
		size        = m.Rows()
		lowerMatrix = MakeSparse(size, size)
		sqSum, sum  float64
	)

	for i := 0; i < size; i++ {
		sqSum = 0.0
		for j := 0; j <= i; j++ {
			if i == j {
				lowerMatrix.SetValue(i, i, math.Sqrt(m.Value(i, i)-sqSum))
			} else {
				if !nums.IsCloseToZero(m.Value(i, j)) {
					sum = 0.0
					for k := 0; k < j; k++ {
						sum += lowerMatrix.Value(i, k) * lowerMatrix.Value(j, k)
					}

					lowerMatrix.SetValue(i, j, (m.Value(i, j)-sum)/lowerMatrix.Value(j, j))
				}

				sqSum += lowerMatrix.Value(i, j) * lowerMatrix.Value(i, j)
			}
		}
	}

	return lowerMatrix
}
