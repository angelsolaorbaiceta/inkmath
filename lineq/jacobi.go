package lineq

import (
	"fmt"
	"math"

	"github.com/angelsolaorbaiceta/inkmath"
	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/*
JacobiSolver is an interative solver for linear equation resolution.
*/
type JacobiSolver struct {
	MaxError float64
	MaxIter  int
}

/* ::::::::::::::: Methods : Solver ::::::::::::::: */

/*
Solve solves the system of equations iteratively until a sufficiently good solution is found
or the maximum number of iterations reached.
*/
func (solver JacobiSolver) Solve(m mat.Matrixable, v *vec.Vector) *LineqSolution {
	var (
		size          = v.Length()
		solution      = vec.Make(size)
		iter          int
		solutionError float64
	)

	solGoodEnough := func() bool {
		vec, _ := m.TimesVector(solution)
		residual, _ := vec.Minus(v)

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
		newSolutionTmp, _ := m.TimesVector(solution)
		newSolution, _ := v.Minus(newSolutionTmp)

		for i := 0; i < size; i++ {
			diagonalValue = m.Value(i, i)

			if inkmath.IsCloseToZero(diagonalValue) {
				fmt.Printf("Row %d\n", i)
				panic("Found a main diagonal value of zero")
			}

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
	return makeSolution(iter, solution)
}
