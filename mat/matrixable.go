/*
Package mat defines matrices and operations with matrices.
*/
package mat

import "github.com/angelsolaorbaiceta/inkmath/vec"

/*
Matrixable defines the generic contract for a matrix.
*/
type Matrixable interface {
	/* Properties */
	Rows() int
	Cols() int

	/* Methods */
	Value(int, int) float64
	SetValue(int, int, float64)
	AddToValue(int, int, float64)

	/* Operations */
	AddInPlace(other Matrixable) error
	TimesInPlace(other Matrixable) error

	TimesVector(v *vec.Vector) (*vec.Vector, error)
}
