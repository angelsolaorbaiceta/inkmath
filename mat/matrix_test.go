package mat

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkmath"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/* <--------------- Dense ---------------> */

func TestAddDenseMatricesInPlace(t *testing.T) {
	matA, matB := makeDenseTestMatrices()
	matA.AddInPlace(matB)

	assertMatrixAddition(matA, t)
}

func TestMultiplyDenseMatricesInPlace(t *testing.T) {
	matA, matB := makeDenseTestMatrices()
	matA.TimesInPlace(matB)

	assertMatrixMultiplication(matA, t)
}

func TestMultiplyDenseMatrices(t *testing.T) {
	matA, matB := makeDenseTestMatrices()
	matC := matA.TimesMatrix(matB)

	assertMatrixMultiplication(matC, t)
}

func TestDenseSquareMatrixTimesVector(t *testing.T) {
	mat := MakeSquareDense(2)
	mat.SetValue(0, 0, 1.0)
	mat.SetValue(0, 1, 2.0)
	mat.SetValue(1, 0, 3.0)
	mat.SetValue(1, 1, 4.0)
	v := vec.MakeWithValues([]float64{5.0, 6.0})
	prod := mat.TimesVector(v)
	expectedVec := vec.MakeWithValues([]float64{17.0, 39.0})

	if !prod.Equals(expectedVec) {
		t.Errorf("Wrong multiplication. Expected %v but got %v", expectedVec, prod)
	}
}

/* <--------------- Sparse ---------------> */

/* Set & Get Values */
func TestSetNonZeroValue(t *testing.T) {
	m := MakeSparse(5, 5)
	m.SetValue(1, 2, 7.5)

	if val := m.Value(1, 2); val != 7.5 {
		t.Errorf("Value not as expected. Got %f", val)
	}
}

func TestSetZeroValueRemovesValue(t *testing.T) {
	m := MakeSparse(5, 5)
	m.SetValue(1, 2, 7.5)
	m.SetValue(1, 2, 0.0)

	if val := m.Value(1, 2); val != 0.0 {
		t.Errorf("Value not as expected. Got %f", val)
	}
}

func TestNonAssignedValueIsZero(t *testing.T) {
	m := MakeSparse(2, 4)
	if val := m.Value(1, 3); val != 0 {
		t.Errorf("Value not as expected. Got %f", val)
	}
}

func TestSetZeroCol(t *testing.T) {
	m := MakeSparse(3, 3)
	m.SetValue(0, 1, 1.0)
	m.SetValue(1, 1, 2.0)
	m.SetValue(2, 1, 3.0)
	m.SetZeroCol(1)

	if !inkmath.IsCloseToZero(m.Value(0, 1)) ||
		!inkmath.IsCloseToZero(m.Value(1, 1)) ||
		!inkmath.IsCloseToZero(m.Value(2, 1)) {
		t.Error("Column expected to be zero")
	}
}

func TestSetIdentityRow(t *testing.T) {
	m := MakeSparse(3, 3)
	m.SetValue(1, 0, 4.0)
	m.SetValue(1, 1, 4.0)
	m.SetValue(1, 2, 4.0)
	m.SetIdentityRow(1)

	if !inkmath.IsCloseToZero(m.Value(1, 0)) ||
		!inkmath.FuzzyEqual(m.Value(1, 1), 1.0) ||
		!inkmath.IsCloseToZero(m.Value(1, 2)) {
		t.Error("Row expected to be identity")
	}
}

/* Operations */
func TestSparseMatrixTimesVector(t *testing.T) {
	mat := MakeSparse(2, 2)
	mat.SetValue(0, 0, 1.0)
	mat.SetValue(0, 1, 2.0)
	mat.SetValue(1, 0, 3.0)
	mat.SetValue(1, 1, 4.0)
	v := vec.MakeWithValues([]float64{5.0, 6.0})
	prod := mat.TimesVector(v)
	expectedVec := vec.MakeWithValues([]float64{17.0, 39.0})

	if !prod.Equals(expectedVec) {
		t.Errorf("Wrong multiplication. Expected %v but got %v", expectedVec, prod)
	}
}

func TestMultiplySparseMatrices(t *testing.T) {
	matA, matB := makeSparseTestMatrices()
	result := matA.TimesMatrix(matB)

	assertMatrixMultiplication(result, t)
}

/* <--------------- Test Data ---------------> */
func makeDenseTestMatrices() (*DenseMat, *DenseMat) {
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

func makeSparseTestMatrices() (*SparseMat, *SparseMat) {
	matA, matB := MakeSparse(2, 2), MakeSparse(2, 2)

	matA.SetValue(0, 0, 2.0)
	matA.SetValue(1, 0, 1.0)
	matA.SetValue(1, 1, 2.0)

	matB.SetValue(0, 0, 1.0)
	matB.SetValue(0, 1, 2.0)
	matB.SetValue(1, 0, 3.0)
	matB.SetValue(1, 1, 4.0)

	return matA, matB
}

func assertMatrixMultiplication(result Matrixable, t *testing.T) {
	if val := result.Value(0, 0); val != 2.0 {
		t.Errorf("Wrong multiplication. Expected 2.0, got %f", val)
	}
	if val := result.Value(0, 1); val != 4.0 {
		t.Errorf("Wrong multiplication. Expected 4.0, got %f", val)
	}
	if val := result.Value(1, 0); val != 7.0 {
		t.Errorf("Wrong multiplication. Expected 7.0, got %f", val)
	}
	if val := result.Value(1, 1); val != 10.0 {
		t.Errorf("Wrong multiplication. Expected 10.0, got %f", val)
	}
}

func assertMatrixAddition(result Matrixable, t *testing.T) {
	if val := result.Value(0, 0); val != 3.0 {
		t.Errorf("Wrong addition. Expected 3.0, got %f", val)
	}
	if val := result.Value(0, 1); val != 2.0 {
		t.Errorf("Wrong addition. Expected 2.0, got %f", val)
	}
	if val := result.Value(1, 0); val != 4.0 {
		t.Errorf("Wrong addition. Expected 4.0, got %f", val)
	}
	if val := result.Value(1, 1); val != 6.0 {
		t.Errorf("Wrong addition. Expected 6.0, got %f", val)
	}
}
