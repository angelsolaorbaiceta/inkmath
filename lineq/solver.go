package lineq

import (
	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/*
Solver is a linear equation system solver.
*/
type Solver interface {
	CanSolve(m mat.Matrixable, v *vec.Vector) bool
	Solve(m mat.Matrixable, v *vec.Vector) *LineqSolution
}
