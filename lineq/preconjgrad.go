package lineq

import (
	"math"

	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/*
PreconditionedConjugateGradientSolver is an interative solver for linear equation resolution where
a preconditioner is used to speed up convergence.

The preconditioner should be a square matrix.
*/
type PreconditionedConjugateGradientSolver struct {
	MaxError float64
	MaxIter  int
}

/*
CanSolve returns whether Conjugate Gradient is suitable for solving the given system of equations.

The conditions required are:
	- System matrix is square
	- System matrix is symmetric
	- System matrix and vector have same size
*/
func (solver PreconditionedConjugateGradientSolver) CanSolve(
	coefficients mat.ReadOnlyMatrix,
	freeTerms vec.ReadOnlyVector,
) bool {
	return mat.IsSquare(coefficients) &&
		coefficients.Rows() == freeTerms.Length() &&
		mat.IsSymmetric(coefficients)
}

/*
Solve solves the system of equations iteratively until a sufficiently good
solution is found or the maximum number of iterations reached.
*/
func (solver PreconditionedConjugateGradientSolver) Solve(
	a mat.ReadOnlyMatrix,
	b vec.ReadOnlyVector,
) *Solution {
	var (
		size                      = b.Length()
		x                         = vec.MakeReadOnly(size)
		precond                   = computePreconditioner(a)
		r, oldr, p, precondTimesR vec.ReadOnlyVector
		alpha, beta, err          float64
		iter                      int
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
	p = precond.TimesVector(r)

	// Iteration loop
	for iter = 0; iter < solver.MaxIter; iter++ {
		if solutionGoodEnough() {
			return makeSolution(iter, solver.MaxError, x)
		}

		alpha = r.Times(precond.TimesVector(r)) / (p.Times(a.TimesVector(p)))
		x = x.Plus(p.Scaled(alpha))
		oldr = r.Clone()
		r = oldr.Minus(a.TimesVector(p).Scaled(alpha))
		precondTimesR = precond.TimesVector(r)
		beta = r.Times(precondTimesR) / oldr.Times(precond.TimesVector(oldr))
		p = precondTimesR.Plus(p.Scaled(beta))
	}

	return makeErrorSolution(iter, err, x)
}

func computePreconditioner(m mat.ReadOnlyMatrix) mat.ReadOnlyMatrix {
	precond := mat.MakeSparse(m.Rows(), m.Cols())
	for i := 0; i < m.Rows(); i++ {
		precond.SetValue(i, i, 1.0/m.Value(i, i))
	}

	return precond
}
