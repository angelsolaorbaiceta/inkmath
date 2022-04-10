package lineq

import (
	"math"

	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

// PreconditionedConjugateGradientSolver is an interative solver for systems of linear
// equations where a preconditioner is used to speed up convergence.
//
// The preconditioner should be a square matrix.
//
// A channel can be added to the solver to receive the progress and the current error.
type PreconditionedConjugateGradientSolver struct {
	MaxError       float64
	MaxIter        int
	Preconditioner mat.ReadOnlyMatrix
	ProgressChan   chan<- IterativeSolverProgress
}

// CanSolve returns whether Conjugate Gradient is suitable for solving the given system
// of equations.
//
// The conditions required are:
// - System matrix is square
// - System matrix is symmetric
// - System matrix and vector have same size
func (solver PreconditionedConjugateGradientSolver) CanSolve(
	coefficients mat.ReadOnlyMatrix,
	freeTerms vec.ReadOnlyVector,
) bool {
	return mat.IsSquare(coefficients) &&
		coefficients.Rows() == freeTerms.Length() &&
		mat.IsSymmetric(coefficients)
}

// Solve solves the system of equations iteratively until a sufficiently good
// solution is found or the maximum number of iterations reached.
func (solver PreconditionedConjugateGradientSolver) Solve(
	a mat.ReadOnlyMatrix,
	b vec.ReadOnlyVector,
) *Solution {
	var (
		size                      = b.Length()
		x                         = vec.MakeReadOnly(size)
		precond                   = solver.Preconditioner
		r, oldr, p, precondTimesR vec.ReadOnlyVector
		alpha, beta, err          float64
		iter                      int
		computeProgressChan       chan computeProgressRequest
	)

	solutionGoodEnough := func() bool {
		for i := 0; i < size; i++ {
			if err = math.Abs(r.Value(i)); err > solver.MaxError {
				return false
			}
		}

		return true
	}

	notifyProgress := func() {
		if computeProgressChan != nil {
			errVec, ok := r.(*vec.Vector)
			if !ok {
				panic("r is not a vec.Vector")
			}

			computeProgressChan <- computeProgressRequest{
				currentErrorFn: func() float64 {
					return computeMaxError(errVec)
				},
				iterCount: iter,
				maxError:  solver.MaxError,
			}
		}
	}

	if solver.ProgressChan != nil {
		computeProgressChan = make(chan computeProgressRequest)
		go computeProgress(computeProgressChan, solver.ProgressChan)

		defer close(computeProgressChan)
	}

	// Initial values
	r = b.Minus(a.TimesVector(x))
	p = precond.TimesVector(r)

	// Iteration loop
	for iter = 0; iter < solver.MaxIter; iter++ {
		if solutionGoodEnough() {
			break
		}

		notifyProgress()

		alpha = r.Times(precond.TimesVector(r)) / (p.Times(a.TimesVector(p)))
		x = x.Plus(p.Scaled(alpha))
		oldr = r.Clone()
		r = oldr.Minus(a.TimesVector(p).Scaled(alpha))
		precondTimesR = precond.TimesVector(r)
		beta = r.Times(precondTimesR) / oldr.Times(precond.TimesVector(oldr))
		p = precondTimesR.Plus(p.Scaled(beta))
	}

	notifyProgress()
	return makeErrorSolution(iter, computeMaxError(r), x)
}

func computeMaxError(errVec vec.ReadOnlyVector) float64 {
	var err float64

	for i := 0; i < errVec.Length(); i++ {
		err = math.Max(err, math.Abs(errVec.Value(i)))
	}

	return err
}
