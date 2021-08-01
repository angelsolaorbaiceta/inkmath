package mat

import "github.com/angelsolaorbaiceta/inkmath/nums"

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
