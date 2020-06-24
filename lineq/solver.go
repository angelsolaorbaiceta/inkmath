package lineq

import (
	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/*
A Solver is a solver for linear systems of equations.
*/
type Solver interface {
	CanSolve(m mat.ReadOnlyMatrix, v *vec.Vector) bool
	Solve(m mat.ReadOnlyMatrix, v *vec.Vector) *Solution
}
