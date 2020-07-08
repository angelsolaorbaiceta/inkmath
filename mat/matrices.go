/*
Copyright 2020 Angel Sola Orbaiceta

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
			if !nums.FuzzyEqual(m1.Value(i, j), m2.Value(i, j)) {
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
			if !nums.FuzzyEqual(m.Value(i, j), m.Value(j, i)) {
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
func MainDiagonal(m ReadOnlyMatrix) *vec.Vector {
	if !IsSquare(m) {
		panic("Matrix main diagonal only applies to square matrices")
	}

	diag := vec.Make(m.Rows())
	for i := 0; i < m.Rows(); i++ {
		diag.SetValue(i, m.Value(i, i))
	}

	return diag
}

func matrixContainsData(matrix ReadOnlyMatrix, data []float64) bool {
	var (
		offset    int
		got, want float64
	)

	for rowIndex := 0; rowIndex < matrix.Rows(); rowIndex++ {
		offset = rowIndex * matrix.Cols()

		for colIndex := 0; colIndex < matrix.Cols(); colIndex++ {
			got = matrix.Value(rowIndex, colIndex)
			want = data[offset+colIndex]

			if !nums.FuzzyEqual(got, want) {
				return false
			}
		}
	}

	return true
}
