/*
Copyright 2020 Angel Sola Orbaiceta

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	RowTimesVector(row int, v *vec.Vector) float64
	TimesVector(v *vec.Vector) *vec.Vector
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
