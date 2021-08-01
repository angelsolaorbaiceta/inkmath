package mat

import "github.com/angelsolaorbaiceta/inkmath/vec"

/*
A ReadOnlyMatrix defines the contract for a matrix whose methods can't
(and shouldn't) mutate the matrix.

The operations defined in a ReadOnlyMatrix should always return a new instance,
never mutate the matrix.
*/
type ReadOnlyMatrix interface {
	/* Properties */
	Rows() int
	Cols() int
	NonZeroIndicesAtRow(int) []int

	/* Methods */
	Value(int, int) float64

	/* Operations */
	RowTimesVector(row int, v vec.ReadOnlyVector) float64
	TimesVector(v vec.ReadOnlyVector) vec.ReadOnlyVector
	TimesMatrix(other ReadOnlyMatrix) ReadOnlyMatrix
}

/*
A MutableMatrix defines the contract for a matrix which implements the
ReadOnlyMatrix and also provides methods that allow the mutation of its data.
*/
type MutableMatrix interface {
	ReadOnlyMatrix

	/* Methods */
	SetValue(int, int, float64)
	AddToValue(int, int, float64)

	SetZeroCol(int)
	SetIdentityRow(int)
}
