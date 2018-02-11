package mat

import (
	"math"

	"github.com/angelsolaorbaiceta/inkmath"
)

/*
CholeskyDecomposition computes the Cholesky lower matrix decomposition
for the given square and symmetric matrix.
*/
func CholeskyDecomposition(m Matrixable) Matrixable {
	if !IsSquare(m) {
		panic("Cannot use Cholesky factorization in non-square matrices")
	}

	var (
		size                             = m.Rows()
		lowerMatrix                      = MakeSparse(size, size)
		nonDiagonalSum, nonDiagonalValue float64
		rowSquareSumBeforeDiagonal       float64
	)

	for i := 0; i < size; i++ {
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

func IncompleteCholeskyDecomposition(m Matrixable) Matrixable {
	if !IsSquare(m) {
		panic("Cannot use Cholesky factorization in non-square matrices")
	}

	// var (
	// 	size                       = m.Rows()
	// 	lowerMatrix                = MakeSparse(size, size)
	// 	val, nonDiagonalSum        float64
	// 	rowSquareSumBeforeDiagonal float64
	// )
	//
	// for i := 0; i < size; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		if i == j {
	// 			// Main Diagonal Value
	// 			rowSquareSumBeforeDiagonal = 0
	// 			for k := 0; k < i; k++ {
	// 				val = m.Value(i, k)
	// 				rowSquareSumBeforeDiagonal += val * val
	// 			}
	// 			fmt.Println(m.Value(i, i) - rowSquareSumBeforeDiagonal)
	// 			lowerMatrix.SetValue(i, i, math.Sqrt(m.Value(i, i)-rowSquareSumBeforeDiagonal))
	// 		} else if val = m.Value(i, j); !inkmath.IsCloseToZero(val) {
	// 			// Value under Main Diagonal
	// 			nonDiagonalSum = 0.0
	// 			for k := 0; k < j; k++ {
	// 				nonDiagonalSum += lowerMatrix.Value(i, k) * lowerMatrix.Value(j, k)
	// 			}
	//
	// 			lowerMatrix.SetValue(i, j, (m.Value(i, j)-nonDiagonalSum)/lowerMatrix.Value(j, j))
	// 		}
	// 	}
	// }

	var (
		size                             = m.Rows()
		lowerMatrix                      = MakeSparse(size, size)
		nonDiagonalSum, nonDiagonalValue float64
		rowSquareSumBeforeDiagonal       float64
	)

	for i := 0; i < size; i++ {
		rowSquareSumBeforeDiagonal = 0

		for j := 0; j <= i; j++ {
			if i == j {
				// Main Diagonal Value
				lowerMatrix.SetValue(i, j, math.Sqrt(m.Value(i, i)-rowSquareSumBeforeDiagonal))
			} else {
				// Value under Main Diagonal
				nonDiagonalSum = 0.0

				if !inkmath.IsCloseToZero(m.Value(i, j)) {
					for k := 0; k < j; k++ {
						nonDiagonalSum += lowerMatrix.Value(i, k) * lowerMatrix.Value(j, k)
					}

					nonDiagonalValue = (m.Value(i, j) - nonDiagonalSum) / lowerMatrix.Value(j, j)
					lowerMatrix.SetValue(i, j, nonDiagonalValue)
					rowSquareSumBeforeDiagonal += nonDiagonalValue * nonDiagonalValue
				}
			}
		}
	}

	return lowerMatrix
}
