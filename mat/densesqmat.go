package mat

/*
DenseSqMat is an implementation of a dense Matrixable with the same number of rows and columns,
that is, square.
*/
type DenseSqMat struct {
	size int
	data [][]float64
}

/* ::::::::::::::: Construction ::::::::::::::: */

// MakeSquareDense creates a new dense matrix (strores zeroes) with the given dimension all filled with zeroes.
func MakeSquareDense(size int) *DenseSqMat {
	data := make([][]float64, size)
	for i := 0; i < size; i++ {
		data[i] = make([]float64, size)
	}
	return &DenseSqMat{size, data}
}

/* ::::::::::::::: Properties ::::::::::::::: */

// Rows returns the number of rows in the matrix. That is, 6.
func (m DenseSqMat) Rows() int {
	return m.size
}

// Cols returns the number of columns in the matrix. That is, 6.
func (m DenseSqMat) Cols() int {
	return m.size
}

/* ::::::::::::::: Methods ::::::::::::::: */

// Value returns the matrix value at a given row and column.
func (m DenseSqMat) Value(row, col int) float64 {
	return m.data[row][col]
}

// SetValue sets the matrix value at a given row and column.
func (m DenseSqMat) SetValue(row, col int, value float64) {
	m.data[row][col] = value
}

// AddToValue adds the given value to the existing value in the indicated row and column.
func (m DenseSqMat) AddToValue(row, col int, value float64) {
	m.data[row][col] += value
}

/* ::::::::::::::: Operations ::::::::::::::: */

// TimesInPlace multiplies this matrix times other and sets the result in this matrix.
func (m *DenseSqMat) TimesInPlace(other *DenseSqMat) {
	var (
		sum float64
		row = make([]float64, m.size)
	)

	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			sum = 0.0
			for k := 0; k < m.size; k++ {
				sum += m.data[i][k] * other.data[k][j]
			}
			row[j] = sum
		}
		for j := 0; j < m.size; j++ {
			m.data[i][j] = row[j]
		}
	}
}
