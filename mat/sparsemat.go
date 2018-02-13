package mat

import (
	"github.com/angelsolaorbaiceta/inkmath"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/*
SparseMat is an implementation of a sparse (does not store zeroes) Matrixable.
*/
type SparseMat struct {
	rows, cols int
	data       map[int]map[int]float64
}

/* ::::::::::::::: Construction ::::::::::::::: */

/*
MakeSparse creates a new sparse matrix with the indicated number of rows and columns.
*/
func MakeSparse(rows, cols int) *SparseMat {
	return &SparseMat{rows, cols, make(map[int]map[int]float64)}
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

/* ::::::::::::::: Properties ::::::::::::::: */

// Rows returns the number of rows in the matrix.
func (m SparseMat) Rows() int { return m.rows }

// Cols returns the number of columns in the matrix.
func (m SparseMat) Cols() int { return m.cols }

/* ::::::::::::::: Methods ::::::::::::::: */

// SetZeroCol sets all the values in the given column as zero.
func (m *SparseMat) SetZeroCol(col int) {
	for i := range m.data {
		m.SetValue(i, col, 0.0)
	}
}

/*
SetIdentityRow sets the given row as identity: one in the main diagonal value,
and zeroes in all other positions of the row.
*/
func (m *SparseMat) SetIdentityRow(row int) {
	if _, hasRow := m.data[row]; hasRow {
		delete(m.data, row)
	}

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
	if inkmath.IsCloseToZero(value) {
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

/*
NonZeroIndicesAtRow returns a slice with all non-zero elements indices for the given row.
*/
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

/* ::::::::::::::: Operations ::::::::::::::: */

/*
AddInPlace adds this matrix with other and sets the aresult in this matrix.
*/
func (m SparseMat) AddInPlace(other Matrixable) error {
	return nil
}

/*
TimesInPlace multiplies this matrix times other and sets the result in this matrix.
*/
func (m SparseMat) TimesInPlace(other Matrixable) error {
	return nil
}

/*
TimesVector multiplies this matrix and a vector.
*/
func (m SparseMat) TimesVector(v *vec.Vector) *vec.Vector {
	if m.Cols() != v.Length() {
		panic("Can't multiply matrix vs vector due to size mismatch")
	}

	var (
		result = vec.Make(m.Cols())
		sum    float64
	)

	for i, row := range m.data {
		sum = 0.0
		for j, val := range row {
			sum += val * v.Value(j)
		}
		result.SetValue(i, sum)
	}

	return result
}

/*
TimesMatrix multiplies this matrix with other.
*/
func (m SparseMat) TimesMatrix(other Matrixable) Matrixable {
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

/*
RowTimesVector returns the result of multiplying the row at the given index
times the given vector.
*/
func (m SparseMat) RowTimesVector(row int, v *vec.Vector) float64 {
	if m.Cols() != v.Length() {
		panic("Can't multiply matrix row with vector due to size mismatch")
	}

	if rowData, hasRow := m.data[row]; hasRow {
		result := 0.0
		for i, val := range rowData {
			result += v.Value(i) * val
		}

		return result
	}

	return 0.0
}
