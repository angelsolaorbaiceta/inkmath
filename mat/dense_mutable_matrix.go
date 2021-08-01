package mat

// SetValue sets a value for a given row and column.
func (m *DenseMat) SetValue(row, col int, value float64) {
	m.data[row][col] = value
}

// AddToValue adds the given value to the existing value in the indicated row and column.
func (m *DenseMat) AddToValue(row, col int, value float64) {
	m.data[row][col] += value
}

// SetZeroCol sets all the values in the given column as zero.
func (m *DenseMat) SetZeroCol(col int) {
	for row := 0; row < m.rows; row++ {
		m.data[row][col] = 0.0
	}
}

/*
SetIdentityRow sets the given row as identity: one in the main diagonal value, and zeroes in all
other positions of the row.
*/
func (m *DenseMat) SetIdentityRow(row int) {
	for col := 0; col < m.cols; col++ {
		m.data[row][col] = 0.0
	}
	m.data[row][row] = 1.0
}
