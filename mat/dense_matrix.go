package mat

import (
	"github.com/angelsolaorbaiceta/inkmath/nums"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/*
A DenseMat is an implementation of a dense Matrix.

Dense matrices allocate all the memory required to store every value.
Every value which hasn't been explecitly set is zero.
*/
type DenseMat struct {
	rows, cols int
	data       [][]float64
}

/* <-- Construction --> */

// MakeSquareDense creates a new dense matrix (strores zeroes) with the given dimension all filled with zeroes.
func MakeSquareDense(size int) *DenseMat {
	data := make([][]float64, size)
	for i := 0; i < size; i++ {
		data[i] = make([]float64, size)
	}

	return &DenseMat{size, size, data}
}

// MakeDense creates a new dense matrix (stores zeroes) with the given rows and columns filled with zeroes.
func MakeDense(rows, cols int) *DenseMat {
	data := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]float64, cols)
	}

	return &DenseMat{rows, cols, data}
}

/* <-- Properties --> */

// Rows returns the number of rows in the matrix.
func (m DenseMat) Rows() int { return m.rows }

// Cols returns the number of columns in the matrix.
func (m DenseMat) Cols() int { return m.cols }

/* <-- Methods --> */

// Value returns the value at a given row and column.
func (m DenseMat) Value(row, col int) float64 {
	return m.data[row][col]
}

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
SetIdentityRow sets the given row as identity: one in the main diagonal value,
and zeroes in all other positions of the row.
*/
func (m *DenseMat) SetIdentityRow(row int) {
	for col := 0; col < m.cols; col++ {
		m.data[row][col] = 0.0
	}
	m.data[row][row] = 1.0
}

/*
NonZeroIndicesAtRow returns a slice with all non-zero elements indices for the given row.
*/
func (m DenseMat) NonZeroIndicesAtRow(row int) []int {
	indices := make([]int, 0)
	for i, val := range m.data[row] {
		if !nums.IsCloseToZero(val) {
			indices = append(indices, i)
		}
	}

	return indices
}

/* <-- Operations --> */

/*
TimesVector multiplies this matrix and a vector.
*/
func (m DenseMat) TimesVector(v *vec.Vector) *vec.Vector {
	if m.Cols() != v.Length() {
		panic("Can't multiply matrix vs vector due to size mismatch")
	}

	var (
		result = vec.Make(m.Cols())
		sum    float64
	)

	for i := 0; i < m.Rows(); i++ {
		sum = 0.0
		for j := 0; j < m.Cols(); j++ {
			sum += m.data[i][j] * v.Value(j)
		}
		result.SetValue(i, sum)
	}

	return result
}

/*
TimesMatrix multiplies this matrix with other.
*/
func (m DenseMat) TimesMatrix(other ReadOnlyMatrix) ReadOnlyMatrix {
	if m.Cols() != other.Rows() {
		panic("Can't multiply matrices due to size mismatch")
	}

	var (
		rows   = m.Rows()
		cols   = other.Cols()
		sum    float64
		result = MakeDense(rows, cols)
	)

	for i := 0; i < rows; i++ {
		// cummulative sum of this.row x other.column
		for j := 0; j < cols; j++ {
			sum = 0.0
			for k := 0; k < other.Rows(); k++ {
				sum += m.data[i][k] * other.Value(k, j)
			}

			result.data[i][j] = sum
		}
	}

	return result
}

/*
RowTimesVector returns the result of multiplying the row at the given index
times the given vector.
*/
func (m DenseMat) RowTimesVector(row int, v *vec.Vector) float64 {
	if m.Cols() != v.Length() {
		panic("Can't multiply matrix row with vector due to size mismatch")
	}

	var (
		rowData = m.data[row]
		result  = 0.0
	)

	for i := 0; i < m.Cols(); i++ {
		result += rowData[i] * v.Value(i)
	}

	return result
}
