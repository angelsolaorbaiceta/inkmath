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

import (
	"math"

	"github.com/angelsolaorbaiceta/inkmath/nums"
)

/*
CholeskyDecomposition computes the Cholesky lower matrix decomposition
for the given square and symmetric matrix.
*/
// TODO: should return pointer
func CholeskyDecomposition(m ReadOnlyMatrix) ReadOnlyMatrix {
	if !IsSquare(m) {
		panic("Cannot use Cholesky factorization in non-square matrices")
	}

	var (
		size        = m.Rows()
		lowerMatrix = MakeSparse(size, size)
		sqSum, sum  float64
	)

	for i := 0; i < size; i++ {
		sqSum = 0.0
		for j := 0; j <= i; j++ {
			if i == j {
				lowerMatrix.SetValue(i, i, math.Sqrt(m.Value(i, i)-sqSum))
			} else {
				sum = 0.0
				for k := 0; k < j; k++ {
					sum += lowerMatrix.Value(i, k) * lowerMatrix.Value(j, k)
				}

				lowerMatrix.SetValue(i, j, (m.Value(i, j)-sum)/lowerMatrix.Value(j, j))
				sqSum += lowerMatrix.Value(i, j) * lowerMatrix.Value(i, j)
			}
		}
	}

	return lowerMatrix
}

/*
IncompleteCholeskyDecomposition computes the Incomplete Cholesky lower matrix
decomposition for the given square and symmetric matrix.
*/
// TODO: should return pointer
func IncompleteCholeskyDecomposition(m ReadOnlyMatrix) ReadOnlyMatrix {
	if !IsSquare(m) {
		panic("Cannot use Cholesky factorization in non-square matrices")
	}

	var (
		size        = m.Rows()
		lowerMatrix = MakeSparse(size, size)
		sqSum, sum  float64
	)

	for i := 0; i < size; i++ {
		sqSum = 0.0
		for j := 0; j <= i; j++ {
			if i == j {
				lowerMatrix.SetValue(i, i, math.Sqrt(m.Value(i, i)-sqSum))
			} else {
				if !nums.IsCloseToZero(m.Value(i, j)) {
					sum = 0.0
					for k := 0; k < j; k++ {
						sum += lowerMatrix.Value(i, k) * lowerMatrix.Value(j, k)
					}

					lowerMatrix.SetValue(i, j, (m.Value(i, j)-sum)/lowerMatrix.Value(j, j))
				}

				sqSum += lowerMatrix.Value(i, j) * lowerMatrix.Value(i, j)
			}
		}
	}

	return lowerMatrix
}
