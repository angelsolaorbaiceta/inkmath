package lineq

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

func TestCGSolveSystem2x2(t *testing.T) {
	var (
		m, v   = makeSystem2x2()
		solver = ConjugateGradientSolver{1e-10, 2}
	)

	sol := solver.Solve(m, v)

	if !sol.Solution.Equals(expectedSol2x2) {
		t.Errorf("Wrong solution, Expected %v, but got %v", expectedSol2x2, sol)
	}
	if sol.MinError > 1e-10 {
		t.Errorf("Wrong error, Expected < 1e-10, but got %f", sol.MinError)
	}
}

func TestPreconditionedCGSolveSystem2x2(t *testing.T) {
	var (
		m, v         = makeSystem2x2()
		progressChan = make(chan IterativeSolverProgress, 3)
		solver       = PreconditionedConjugateGradientSolver{
			MaxError:       1e-10,
			MaxIter:        2,
			Preconditioner: mat.MakeDenseWithData(2, 2, []float64{1.0 / 4.0, 0, 0, 1.0 / 3.0}),
			ProgressChan:   progressChan,
		}
		sol = solver.Solve(m, v)
	)

	if !sol.Solution.Equals(expectedSol2x2) {
		t.Errorf("Wrong solution, Expected %v, but got %v", expectedSol2x2, sol)
	}
	if sol.MinError > 1e-10 {
		t.Errorf("Wrong error, Expected < 1e-10, but got %f", sol.MinError)
	}

	p1, p2, p3 := <-progressChan, <-progressChan, <-progressChan
	if got := p1.ProgressPercentage; got > 2 {
		t.Errorf("Want 0-2%% progress, got %d%%", got)
	}
	if got := p1.IterCount; got != 0 {
		t.Errorf("Want 0 iterations, got %d", got)
	}

	if got := p2.ProgressPercentage; got < 2 || got > 4 {
		t.Errorf("Want 2-4%% progress, got %d%%", got)
	}
	if got := p2.IterCount; got != 1 {
		t.Errorf("Want 1 iterations, got %d", got)
	}

	if got := p3.ProgressPercentage; got != 100.0 {
		t.Errorf("Want 100%% progress, got %d%%", got)
	}
	if got := p3.IterCount; got != 2 {
		t.Errorf("Want 2 iterations, got %d", got)
	}
}

func TestJacobiSolveSystem2x2(t *testing.T) {
	var (
		m, v   = makeSystem2x2()
		solver = JacobiSolver{1e-10, 50}
	)

	if sol := solver.Solve(m, v); !sol.Solution.Equals(expectedSol2x2) {
		t.Errorf("Wrong solution, Expected %v, but got %v", expectedSol2x2, sol)
	}
}

func TestGaussSeidelSolveSystem2x2(t *testing.T) {
	var (
		m, v   = makeSystem2x2()
		solver = GaussSeidelSolver{1e-10, 50}
	)

	if sol := solver.Solve(m, v); !sol.Solution.Equals(expectedSol2x2) {
		t.Errorf("Wrong solution, Expected %v, but got %v", expectedSol2x2, sol)
	}
}

var expectedSol2x2 = vec.MakeWithValues([]float64{1.0 / 11.0, 7.0 / 11.0})

func makeSystem2x2() (mat.MutableMatrix, vec.ReadOnlyVector) {
	m := mat.MakeSquareDense(2)
	m.SetValue(0, 0, 4.0)
	m.SetValue(0, 1, 1.0)
	m.SetValue(1, 0, 1.0)
	m.SetValue(1, 1, 3.0)

	v := vec.MakeWithValues([]float64{1.0, 2.0})

	return m, v
}
