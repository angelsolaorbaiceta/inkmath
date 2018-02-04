package lineq

import (
	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/*
ConjugateGradientSolver is an interative solver for linear equation resolution.
*/
type ConjugateGradientSolver struct {
	MaxError float64
	MaxIter  int
}

/* ::::::::::::::: Methods : Solver ::::::::::::::: */

/*
Solve solves the system of equations iteratively until a sufficiently good solution is found
or the maximum number of iterations reached.
*/
func (solver ConjugateGradientSolver) Solve(m mat.Matrixable, v *vec.Vector) *LineqSolution {
	// var (
	// 	size     = v.Length()
	// 	sol      = vec.Make(size)
	// 	r        vec.Vector
	// 	residual float64
	// )
	//
	// solGoodEnough := func(sol *vec.Vector) bool {
	// 	return residual < solver.MaxError
	// }

	return &LineqSolution{false, 0, 1.0, v}
}
