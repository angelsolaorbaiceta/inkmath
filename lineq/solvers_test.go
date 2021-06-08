package lineq

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/* <--------------- System 2x2 ---------------> */
func TestCGSolveSystem2x2(t *testing.T) {
	var (
		m, v   = makeSystem2x2()
		solver = ConjugateGradientSolver{1e-10, 2}
	)

	if sol := solver.Solve(m, v); !sol.Solution.Equals(expectedSol2x2) {
		t.Errorf("Wrong solution, Expected %v, but got %v", expectedSol2x2, sol)
	}
}

func TestPreconditionedWithIdentityCGSolveSystem2x2(t *testing.T) {
	var (
		m, v   = makeSystem2x2()
		solver = PreconditionedConjugateGradientSolver{1e-10, 2}
	)

	if sol := solver.Solve(m, v); !sol.Solution.Equals(expectedSol2x2) {
		t.Errorf("Wrong solution, Expected %v, but got %v", expectedSol2x2, sol)
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

func makeSystem2x2() (mat.MutableMatrix, *vec.Vector) {
	m := mat.MakeSquareDense(2)
	m.SetValue(0, 0, 4.0)
	m.SetValue(0, 1, 1.0)
	m.SetValue(1, 0, 1.0)
	m.SetValue(1, 1, 3.0)

	v := vec.MakeWithValues([]float64{1.0, 2.0})

	return m, v
}
