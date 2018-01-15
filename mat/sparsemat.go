package mat

import (
	"github.com/angelsolaorbaiceta/inkmath"
)

/*
SparseMat is an implementation of a sparse (does not store zeroes) Matrixable.
*/
type SparseMat struct {
	rows, cols int
	data       map[int]map[int]float64
}

/* ::::::::::::::: Construction ::::::::::::::: */

// MakeSparse creates a new sparse matrix with the indicated number of rows and columns.
func MakeSparse(rows, cols int) *SparseMat {
	return &SparseMat{rows, cols, make(map[int]map[int]float64)}
}

/* ::::::::::::::: Properties ::::::::::::::: */

// Rows returns the number of rows in the matrix.
func (m SparseMat) Rows() int { return m.Rows() }

// Cols returns the number of columns in the matrix.
func (m SparseMat) Cols() int { return m.Cols() }

/* ::::::::::::::: Methods ::::::::::::::: */

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
	if dataRow, ok := m.data[row]; ok {
		delete(dataRow, col)
	}
}

// AddToValue adds the given value to the existing value in the indicated row and column.
func (m *SparseMat) AddToValue(int, int, float64) {
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
