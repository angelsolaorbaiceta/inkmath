package mat

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkmath/vec"
)

func TestAddInPlace(t *testing.T) {
	matA, matB := makeTestMatrices()
	matA.AddInPlace(matB)

	if val := matA.Value(0, 0); val != 3.0 {
		t.Errorf("Wrong addition. Expected 3.0, got %f", val)
	}
	if val := matA.Value(0, 1); val != 2.0 {
		t.Errorf("Wrong addition. Expected 2.0, got %f", val)
	}
	if val := matA.Value(1, 0); val != 4.0 {
		t.Errorf("Wrong addition. Expected 4.0, got %f", val)
	}
	if val := matA.Value(1, 1); val != 6.0 {
		t.Errorf("Wrong addition. Expected 6.0, got %f", val)
	}
}

func TestMultiplyInPlace(t *testing.T) {
	matA, matB := makeTestMatrices()
	matA.TimesInPlace(matB)

	if val := matA.Value(0, 0); val != 2.0 {
		t.Errorf("Wrong multiplication. Expected 2.0, got %f", val)
	}
	if val := matA.Value(0, 1); val != 4.0 {
		t.Errorf("Wrong multiplication. Expected 4.0, got %f", val)
	}
	if val := matA.Value(1, 0); val != 7.0 {
		t.Errorf("Wrong multiplication. Expected 7.0, got %f", val)
	}
	if val := matA.Value(1, 1); val != 10.0 {
		t.Errorf("Wrong multiplication. Expected 10.0, got %f", val)
	}
}

func TestDenseSquareMatrixTimesVector(t *testing.T) {
	mat := MakeSquareDense(2)
	mat.SetValue(0, 0, 1.0)
	mat.SetValue(0, 1, 2.0)
	mat.SetValue(1, 0, 3.0)
	mat.SetValue(1, 1, 4.0)
	v := vec.MakeWithValues([]float64{5.0, 6.0})
	_, prod := mat.TimesVector(v)
	expectedVec := vec.MakeWithValues([]float64{17.0, 39.0})

	if !prod.Equals(expectedVec) {
		t.Errorf("Wrong multiplication. Expected %v but got %v", expectedVec, prod)
	}
}

/* Test Values */
func makeTestMatrices() (*DenseSqMat, *DenseSqMat) {
	matA, matB := MakeSquareDense(2), MakeSquareDense(2)

	matA.SetValue(0, 0, 2.0)
	matA.SetValue(1, 0, 1.0)
	matA.SetValue(1, 1, 2.0)

	matB.SetValue(0, 0, 1.0)
	matB.SetValue(0, 1, 2.0)
	matB.SetValue(1, 0, 3.0)
	matB.SetValue(1, 1, 4.0)

	return matA, matB
}
