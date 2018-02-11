package mat

import "math"

/*
CholeskyDecomposition computes the Cholesky lower matrix decomposition
for the given square and symmetric matrix.
*/
func CholeskyDecomposition(m Matrixable) Matrixable {
	if !IsSquare(m) {
		panic("Cannot use Cholesky factorization in non-square matrices")
	}

	var (
		systemSize                       = m.Rows()
		lowerMatrix                      = MakeSparse(systemSize, systemSize)
		nonDiagonalSum, nonDiagonalValue float64
		rowSquareSumBeforeDiagonal       float64
	)

	for i := 0; i < systemSize; i++ {
		rowSquareSumBeforeDiagonal = 0

		for j := 0; j <= i; j++ {
			if i == j {
				// Main Diagonal Value
				lowerMatrix.SetValue(i, j, math.Sqrt(m.Value(i, i)-rowSquareSumBeforeDiagonal))
			} else {
				// Value under Main Diagonal
				nonDiagonalSum = 0.0
				for k := 0; k < j; k++ {
					nonDiagonalSum += lowerMatrix.Value(i, k) * lowerMatrix.Value(j, k)
				}

				nonDiagonalValue = (m.Value(i, j) - nonDiagonalSum) / lowerMatrix.Value(j, j)
				lowerMatrix.SetValue(i, j, nonDiagonalValue)
				rowSquareSumBeforeDiagonal += nonDiagonalValue * nonDiagonalValue
			}
		}
	}

	return lowerMatrix
}
