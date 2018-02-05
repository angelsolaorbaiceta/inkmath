package lineq

import (
	"math"

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
CanSolve returns whether Conjugate Gradient is suitable for solving the given system of equations.

The conditions required are:
    - System matrix is square
	- System matrix is symmetric
    - System matrix and vector have same size
*/
func (solver ConjugateGradientSolver) CanSolve(m mat.Matrixable, v *vec.Vector) bool {
	return mat.IsSquare(m) && m.Rows() == v.Length() && mat.IsSymmetric(m)
}

/*
Solve solves the system of equations iteratively until a sufficiently good solution is found
or the maximum number of iterations reached.
*/
func (solver ConjugateGradientSolver) Solve(a mat.Matrixable, b *vec.Vector) *LineqSolution {
	var (
		size             = b.Length()
		x                = vec.Make(size)
		r, oldr, p       *vec.Vector
		alpha, beta, err float64
		iter             int
	)

	solutionGoodEnough := func() bool {
		for i := 0; i < size; i++ {
			if err = math.Abs(r.Value(i)); err > solver.MaxError {
				return false
			}
		}

		return true
	}

	// Initial values
	r = b.Minus(a.TimesVector(x))
	p = r.Clone()

	// Iteration loop
	for iter = 0; iter < solver.MaxIter; iter++ {
		alpha = r.Times(r) / p.Times(a.TimesVector(p))
		x = x.Plus(p.Scaled(alpha))
		oldr = r.Clone()
		r = r.Minus(a.TimesVector(p).Scaled(alpha))

		if solutionGoodEnough() {
			return makeSolution(iter, solver.MaxError, x)
		}

		beta = r.Times(r) / oldr.Times(oldr)
		p = r.Plus(p.Scaled(beta))
	}

	return makeErrorSolution(iter, err, x)
}
