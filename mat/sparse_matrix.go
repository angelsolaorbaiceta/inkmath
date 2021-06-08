package mat

import (
	"github.com/angelsolaorbaiceta/inkmath/nums"
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

// SetZeroCol sets all the values in the given column as zero.
func (m *SparseMat) SetZeroCol(col int) {
	for i := range m.data {
		m.SetValue(i, col, 0.0)
	}
}

/*
SetIdentityRow sets the given row as identity: one in the main diagonal value, and zeroes in all
other positions of the row.
*/
func (m *SparseMat) SetIdentityRow(row int) {
	delete(m.data, row)
	m.SetValue(row, row, 1.0)
}

// Value returns the value at a given row and column.
func (m SparseMat) Value(row, col int) float64 {
	if dataRow, hasRow := m.data[row]; hasRow {
		if value, hasValue := dataRow[col]; hasValue {
			return value
		}
	}

	return 0.0
}

// SetValue sets a value for a given row and column.
func (m *SparseMat) SetValue(row, col int, value float64) {
	if nums.IsCloseToZero(value) {
		m.removeValueAt(row, col)
	} else {
		m.setValueToAdd(row, col, value)
	}
}

func (m *SparseMat) setValueToAdd(row, col int, value float64) {
	if _, ok := m.data[row]; !ok {
		m.data[row] = make(map[int]float64)
	}
	m.data[row][col] = value
}

func (m *SparseMat) removeValueAt(row, col int) {
	if dataRow, hasRow := m.data[row]; hasRow {
		delete(dataRow, col)
	}
}

// AddToValue adds the given value to the existing value in the indicated row and column.
func (m *SparseMat) AddToValue(row, col int, value float64) {
	if dataRow, hasRow := m.data[row]; hasRow {
		dataRow[col] += value
	} else {
		m.SetValue(row, col, value)
	}
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
func (m SparseMat) TimesVector(vector *vec.Vector) *vec.Vector {
	if m.cols != vector.Length() {
		panic("Can't multiply matrix and vector due to size mismatch")
	}

	var (
		result = vec.Make(m.rows)
		ch     = make(chan vectMultResult)
	)

	for rowIndex := range m.data {
		go m.rowTimesVectorRoutine(rowIndex, vector, ch)
	}

	for j := 0; j < m.rows; j++ {
		multResult := <-ch
		result.SetValue(multResult.index, multResult.value)
	}
	close(ch)

	return result
}

// TimesMatrix multiplies this matrix times other.
func (m SparseMat) TimesMatrix(other ReadOnlyMatrix) ReadOnlyMatrix {
	if m.Cols() != other.Rows() {
		panic("Can't multiply matrices due to size mismatch")
	}

	var (
		rows   = m.Rows()
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
func (m SparseMat) RowTimesVector(row int, vector *vec.Vector) float64 {
	if m.Cols() != vector.Length() {
		panic("Can't multiply matrix row with vector due to size mismatch")
	}

	return m.rowTimesVector(row, vector)
}

func (m SparseMat) rowTimesVector(row int, vector *vec.Vector) float64 {
	if rowData, hasRow := m.data[row]; hasRow {
		result := 0.0

		for i, val := range rowData {
			result += vector.Value(i) * val
		}

		return result
	}

	return 0.0
}

type vectMultResult struct {
	index int
	value float64
}

func (m SparseMat) rowTimesVectorRoutine(row int, vector *vec.Vector, ch chan<- vectMultResult) {
	result := vectMultResult{index: row, value: m.rowTimesVector(row, vector)}
	ch <- result
}
