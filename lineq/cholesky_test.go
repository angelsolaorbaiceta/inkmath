package lineq

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkmath/mat"
)

func TestCholeskyDecomposition(t *testing.T) {
	var (
		m                     = makeCholeskyMatrix()
		expectedDecomposition = makeCholeskyDecomposition()
		cholesky              = CholeskyDecomposition(m)
	)

	if !mat.AreEqual(cholesky, expectedDecomposition) {
		t.Error("Wrong Cholesky factorization")
	}
}

func TestIncompleteCholeskyDecomposition(t *testing.T) {
	var (
		m                     = makeCholeskyMatrix()
		expectedDecomposition = makeCholeskyDecomposition()
		cholesky              = IncompleteCholeskyDecomposition(m)
	)

	if !mat.AreEqual(cholesky, expectedDecomposition) {
		t.Error("Wrong Cholesky factorization")
	}
}

func makeCholeskyMatrix() mat.ReadOnlyMatrix {
	m := mat.MakeSparse(4, 4)

	m.SetValue(0, 0, 4.0)
	m.SetValue(0, 1, -2.0)
	m.SetValue(0, 2, 4.0)
	m.SetValue(0, 3, 2.0)

	m.SetValue(1, 0, -2.0)
	m.SetValue(1, 1, 10.0)
	m.SetValue(1, 2, -2.0)
	m.SetValue(1, 3, -7.0)

	m.SetValue(2, 0, 4.0)
	m.SetValue(2, 1, -2.0)
	m.SetValue(2, 2, 8.0)
	m.SetValue(2, 3, 4.0)

	m.SetValue(3, 0, 2.0)
	m.SetValue(3, 1, -7.0)
	m.SetValue(3, 2, 4.0)
	m.SetValue(3, 3, 7.0)

	return m
}

func makeCholeskyDecomposition() mat.ReadOnlyMatrix {
	m := mat.MakeSquareDense(4)

	m.SetValue(0, 0, 2.0)

	m.SetValue(1, 0, -1.0)
	m.SetValue(1, 1, 3.0)

	m.SetValue(2, 0, 2.0)
	m.SetValue(2, 2, 2.0)

	m.SetValue(3, 0, 1.0)
	m.SetValue(3, 1, -2.0)
	m.SetValue(3, 2, 1.0)
	m.SetValue(3, 3, 1.0)

	return m
}
