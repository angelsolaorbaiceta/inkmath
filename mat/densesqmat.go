package mat

import (
	"errors"

	"github.com/angelsolaorbaiceta/inkmath/vec"
)

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

// Rows returns the number of rows in the matrix.
func (m DenseSqMat) Rows() int { return m.size }

// Cols returns the number of columns in the matrix.
func (m DenseSqMat) Cols() int { return m.size }

/* ::::::::::::::: Methods ::::::::::::::: */

// Value returns the value at a given row and column.
func (m DenseSqMat) Value(row, col int) float64 {
	return m.data[row][col]
}

// SetValue sets a value for a given row and column.
func (m DenseSqMat) SetValue(row, col int, value float64) {
	m.data[row][col] = value
}

// AddToValue adds the given value to the existing value in the indicated row and column.
func (m DenseSqMat) AddToValue(row, col int, value float64) {
	m.data[row][col] += value
}

/* ::::::::::::::: Operations ::::::::::::::: */

/*
AddInPlace adds this matrix with other and sets the aresult in this matrix.
As this is a square matrix, it is required that the other matrix is square as well.
*/
func (m *DenseSqMat) AddInPlace(other Matrixable) error {
	if m.Rows() != other.Rows() || m.Cols() != other.Cols() {
		return errors.New("Can't add matrices due to size mismatch")
	}

	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			m.data[i][j] += other.Value(i, j)
		}
	}

	return nil
}

/*
TimesInPlace multiplies this matrix times other and sets the result in this matrix.
As this is a square matrix, it is required that the other matrix is square as well.
*/
func (m *DenseSqMat) TimesInPlace(other Matrixable) error {
	if m.Rows() != other.Rows() || m.Cols() != other.Cols() {
		return errors.New("Can't multiply matrices due to size mismatch")
	}

	var (
		sum float64
		row = make([]float64, m.Rows())
	)

	for i := 0; i < m.Rows(); i++ {
		// cummulative sum of this.row x other.column
		for j := 0; j < other.Cols(); j++ {
			sum = 0.0
			for k := 0; k < m.Cols(); k++ {
				sum += m.data[i][k] * other.Value(k, j)
			}
			row[j] = sum
		}

		// set new values in this matrix
		for j := 0; j < m.Cols(); j++ {
			m.data[i][j] = row[j]
		}
	}

	return nil
}

/*
TimesVector multiplies this matrix and a vector.
*/
func (m DenseSqMat) TimesVector(v *vec.Vector) *vec.Vector {
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
