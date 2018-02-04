package lineq

import (
	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/*
Solver is a linear equation system solver.
*/
type Solver interface {
	Solve(m mat.Matrixable, v *vec.Vector) *LineqSolution
}
