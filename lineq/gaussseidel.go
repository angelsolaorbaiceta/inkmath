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

package lineq

import (
	"math"

	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/*
GaussSeidelSolver is an interative solver for linear equation resolution.
*/
type GaussSeidelSolver struct {
	MaxError float64
	MaxIter  int
}

/* <-- Methods : Solver --> */

/*
CanSolve returns whether Jacobi is suitable for solving the given system of equations.

The conditions required are:
    - System matrix is square
    - System matrix and vector have same size
	- System matrix has no zeroes in main diagonal
*/
func (solver GaussSeidelSolver) CanSolve(m mat.ReadOnlyMatrix, v *vec.Vector) bool {
	return mat.IsSquare(m) && m.Rows() == v.Length() && !mat.HasZeroInMainDiagonal(m)
}

/*
Solve solves the system of equations iteratively until a sufficiently good
solution is found or the maximum number of iterations reached.
*/
func (solver GaussSeidelSolver) Solve(m mat.ReadOnlyMatrix, v *vec.Vector) *Solution {
	var (
		size          = v.Length()
		solution      = vec.Make(size)
		iter          int
		solutionError float64
	)

	solutionGoodEnough := func() bool {
		for row := 0; row < size; row++ {
			solutionError = math.Abs(m.RowTimesVector(row, solution) - v.Value(row))
			if solutionError > solver.MaxError {
				return false
			}
		}

		return true
	}

	improveSolution := func() {
		var total float64

		for i := 0; i < size; i++ {
			total = 0.0
			for k := 0; k < size; k++ {
				if i == k {
					continue
				}

				total += m.Value(i, k) * solution.Value(k)
			}

			solution.SetValue(i, (v.Value(i)-total)/m.Value(i, i))
		}
	}

	for iter = 0; iter < solver.MaxIter; iter++ {
		if solutionGoodEnough() {
			return makeSolution(iter, solver.MaxError, solution)
		}

		improveSolution()
	}

	return makeErrorSolution(iter, solutionError, solution)
}
