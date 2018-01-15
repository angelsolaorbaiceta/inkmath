/*
Package mat defines matrices and operations with matrices.
*/
package mat

/*
Matrixable defines the generic contract for a matrix.
*/
type Matrixable interface {
	Rows() int
	Cols() int

	Value(int, int) float64
	SetValue(int, int, float64)
	AddToValue(int, int, float64)
}
