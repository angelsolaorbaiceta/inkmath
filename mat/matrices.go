package mat

import (
	"math"

	"github.com/angelsolaorbaiceta/inkmath/nums"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/*
AreEqual returns true iff matrices have the same number of rows and columns with
exactly the same values in matching positions.
*/
func AreEqual(m1, m2 ReadOnlyMatrix) bool {
	if m1.Rows() != m2.Rows() || m1.Cols() != m2.Cols() {
		return false
	}

	for i := 0; i < m1.Rows(); i++ {
		for j := 0; j < m1.Cols(); j++ {
			if !nums.FloatsEqual(m1.Value(i, j), m2.Value(i, j)) {
				return false
			}
		}
	}

	return true
}

/*
IsSquare returns true if the given matrix has the same number of rows and
columns.
*/
func IsSquare(m ReadOnlyMatrix) bool {
	return m.Rows() == m.Cols()
}

/*
IsSymmetric returns true if the given matrix is square and equals to it's
traspose.
*/
func IsSymmetric(m ReadOnlyMatrix) bool {
	if !IsSquare(m) {
		panic("Matrix symmetry only applies to square matrices")
	}

	for i := 0; i < m.Rows(); i++ {
		for j := i + 1; j < m.Cols(); j++ {
			if !nums.FloatsEqual(m.Value(i, j), m.Value(j, i)) {
				return false
			}
		}
	}

	return true
}

/*
IsRowDominant returns true if for every row in the matrix, the element in the
main diagonal is greater than every other element.
*/
func IsRowDominant(m ReadOnlyMatrix) bool {
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

/*
HasZeroInMainDiagonal returns true if a zero is found in the matrix main diagonal.
*/
func HasZeroInMainDiagonal(m ReadOnlyMatrix) bool {
	if !IsSquare(m) {
		panic("Matrix main diagonal only applies to square matrices")
	}

	for i := 0; i < m.Rows(); i++ {
		if nums.IsCloseToZero(m.Value(i, i)) {
			return true
		}
	}

	return false
}

/*
MainDiagonal returns a vector containing the values of the main diagonal.
*/
func MainDiagonal(m ReadOnlyMatrix) vec.ReadOnlyVector {
	if !IsSquare(m) {
		panic("Matrix main diagonal only applies to square matrices")
	}

	diag := vec.Make(m.Rows())
	for i := 0; i < m.Rows(); i++ {
		diag.SetValue(i, m.Value(i, i))
	}

	return diag
}

/*
MatrixContainsData tests whether a given matrix has exactly the same data as the slice
of float64 numbers.

The number of items in the matrix and the slice need to be the same in order for this
test to return true.
*/
func MatrixContainsData(matrix ReadOnlyMatrix, data []float64) bool {
	var (
		offset    int
		got, want float64
	)

	if matrix.Rows()*matrix.Cols() != len(data) {
		return false
	}

	for rowIndex := 0; rowIndex < matrix.Rows(); rowIndex++ {
		offset = rowIndex * matrix.Cols()

		for colIndex := 0; colIndex < matrix.Cols(); colIndex++ {
			got = matrix.Value(rowIndex, colIndex)
			want = data[offset+colIndex]

			if !nums.FloatsEqual(got, want) {
				return false
			}
		}
	}

	return true
}

/*
FillMatrixWithData fills the matrix using the given slice of float64 numbers. This
function expects the matrix to have the same number of items as the slice.
*/
func FillMatrixWithData(matrix MutableMatrix, data []float64) {
	if matrix.Rows()*matrix.Cols() != len(data) {
		panic("Wrong number of items: can't initialize matrix")
	}

	var (
		offset int
		rows   = matrix.Rows()
		cols   = matrix.Cols()
	)

	for rowIndex := 0; rowIndex < rows; rowIndex++ {
		offset = rowIndex * cols

		for colIndex := 0; colIndex < cols; colIndex++ {
			matrix.SetValue(rowIndex, colIndex, data[offset+colIndex])
		}
	}
}
