package mat

import (
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

// A SparseMat is a matrix where the zeroes aren't stored.
type SparseMat struct {
	rows, cols int
	data       map[int]map[int]float64
}

// MakeSquareSparse creates a new square sparse matrix with the given number of rows and columns.
func MakeSquareSparse(size int) *SparseMat {
	return MakeSparse(size, size)
}

// MakeSparse creates a new sparse matrix with the given number of rows and columns.
func MakeSparse(rows, cols int) *SparseMat {
	return &SparseMat{rows, cols, make(map[int]map[int]float64)}
}

/*
MakeSparseWithData creates a new sparse matrix initialized with the given data.

This method is mainly used for testing purposes as it makes no sense to create a sparse
matrix with a given data slice. Most of the elements in a sparse matrix should be zero.
*/
func MakeSparseWithData(rows, cols int, data []float64) *SparseMat {
	matrix := MakeSparse(rows, cols)
	FillMatrixWithData(matrix, data)

	return matrix
}

/*
MakeIdentity creates a new sparse matrix with all zeroes except in the main diagonal,
which has ones.
*/
func MakeIdentity(size int) *SparseMat {
	identity := MakeSparse(size, size)
	for i := 0; i < size; i++ {
		identity.SetValue(i, i, 1.0)
	}

	return identity
}

// Rows returns the number of rows in the matrix.
func (m SparseMat) Rows() int { return m.rows }

// Cols returns the number of columns in the matrix.
func (m SparseMat) Cols() int { return m.cols }

// Value returns the value at a given row and column.
func (m SparseMat) Value(row, col int) float64 {
	if dataRow, hasRow := m.data[row]; hasRow {
		if value, hasValue := dataRow[col]; hasValue {
			return value
		}
	}

	return 0.0
}

// NonZeroIndicesAtRow returns a slice with all non-zero elements indices for the given row.
func (m SparseMat) NonZeroIndicesAtRow(row int) []int {
	if dataRow, hasRow := m.data[row]; hasRow {
		var (
			keys = make([]int, len(m.data[row]))
			i    = 0
		)

		for k := range dataRow {
			keys[i] = k
			i++
		}

		return keys
	}

	return []int{}
}

// TimesVector multiplies this matrix and a vector.
func (m SparseMat) TimesVector(vector vec.ReadOnlyVector) vec.ReadOnlyVector {
	if m.cols != vector.Length() {
		panic("Can't multiply matrix and vector due to size mismatch")
	}

	result := vec.Make(m.rows)

	for rowIndex := range m.data {
		result.SetValue(rowIndex, m.rowTimesVector(rowIndex, vector))
	}

	return result
}

// TimesMatrix multiplies this matrix times other.
func (m SparseMat) TimesMatrix(other ReadOnlyMatrix) ReadOnlyMatrix {
	if m.cols != other.Rows() {
		panic("Can't multiply matrices due to size mismatch")
	}

	var (
		rows   = m.rows
		cols   = other.Cols()
		sum    float64
		result = MakeSparse(rows, cols)
	)

	for i, row := range m.data {
		for j := 0; j < cols; j++ {

			sum = 0.0

			for k, val := range row {
				sum += val * other.Value(k, j)
			}

			result.SetValue(i, j, sum)
		}
	}

	return result
}

// RowTimesVector returns the result of multiplying the row at the given index times the given vector.
func (m SparseMat) RowTimesVector(row int, vector vec.ReadOnlyVector) float64 {
	if m.cols != vector.Length() {
		panic("Can't multiply matrix row with vector due to size mismatch")
	}

	return m.rowTimesVector(row, vector)
}

func (m SparseMat) rowTimesVector(row int, vector vec.ReadOnlyVector) float64 {
	if rowData, hasRow := m.data[row]; hasRow {
		result := 0.0

		for i, val := range rowData {
			result += vector.Value(i) * val
		}

		return result
	}

	return 0.0
}
