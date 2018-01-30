package lineq

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkmath/mat"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

func TestSolveSystem2x2(t *testing.T) {
	m := mat.MakeSquareDense(2)
	m.SetValue(0, 0, 4.0)
	m.SetValue(0, 1, 1.0)
	m.SetValue(1, 0, 1.0)
	m.SetValue(1, 1, 3.0)
	v := vec.MakeWithValues([]float64{1.0, 2.0})
	expectedSol := vec.MakeWithValues([]float64{1.0 / 11.0, 7.0 / 11.0})
	solver := ConjugateGradientSolver{1e-10, 50}

	if sol := solver.Solve(m, v); !sol.Equals(expectedSol) {
		t.Errorf("Wrong solution, Expected %v, but got %v", expectedSol, sol)
	}
}
