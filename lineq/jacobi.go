package lineq

import (
	"math"

	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

// JacobiSolver is an interative solver for linear equation resolution.
type JacobiSolver struct {
	MaxError float64
	MaxIter  int
}

/*
CanSolve returns whether Jacobi is suitable for solving the given system of equations.

The conditions required are:
  - System matrix is square
  - System matrix and vector have same size
	- System matrix has no zeroes in main diagonal
*/
func (solver JacobiSolver) CanSolve(m mat.ReadOnlyMatrix, v *vec.Vector) bool {
	return mat.IsSquare(m) && m.Rows() == v.Length() && !mat.HasZeroInMainDiagonal(m)
}

/*
Solve solves the system of equations iteratively until a sufficiently good solution is found or the
maximum number of iterations reached.
*/
func (solver JacobiSolver) Solve(m mat.ReadOnlyMatrix, v *vec.Vector) *Solution {
	var (
		size          = v.Length()
		solution      = vec.Make(size)
		iter          int
		solutionError float64
	)

	solGoodEnough := func() bool {
		residual := m.TimesVector(solution).Minus(v)

		for i := 0; i < size; i++ {
			if err := math.Abs(residual.Value(i)); err > solver.MaxError {
				solutionError = err
				return false
			}
		}

		return true
	}

	improveSol := func() {
		var diagonalValue float64
		newSolution := v.Minus(m.TimesVector(solution))

		for i := 0; i < size; i++ {
			diagonalValue = m.Value(i, i)

			newSolution.SetValue(
				i,
				(newSolution.Value(i)+diagonalValue*solution.Value(i))/diagonalValue,
			)
		}

		solution = newSolution
	}

	for iter = 1; iter <= solver.MaxIter && !solGoodEnough(); iter++ {
		improveSol()
	}

	if iter >= solver.MaxIter {
		return makeErrorSolution(iter, solutionError, solution)
	}
	return makeSolution(iter, solver.MaxError, solution)
}
