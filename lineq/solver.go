package lineq

import (
	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

// A Solver is an implementation of a method for solving linear systems of equations.
type Solver interface {
	CanSolve(coefficients mat.ReadOnlyMatrix, freeTerms vec.ReadOnlyVector) bool
	Solve(coefficients mat.ReadOnlyMatrix, freeTerms vec.ReadOnlyVector) *Solution
}
