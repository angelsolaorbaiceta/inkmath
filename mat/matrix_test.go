package mat

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkmath"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

/* <--------------- Dense ---------------> */

func TestAddToValueDense(t *testing.T) {
	m := MakeDense(2, 2)
	m.AddToValue(0, 1, 5.0)
	m.AddToValue(0, 1, 6.0)

	if m.Value(0, 1) != 11.0 {
		t.Error("Value not as expected")
	}
}

func TestSetZeroColDense(t *testing.T) {
	m := MakeDense(3, 3)
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

func TestSetIdentityRowDense(t *testing.T) {
	m := MakeDense(3, 3)
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

func TestNonZeroIndicesInRowDense(t *testing.T) {
	m := MakeDense(3, 3)
	m.SetValue(1, 1, 4.0)
	indices := m.NonZeroIndicesAtRow(1)

	if len(indices) != 1 {
		t.Error("Non zero indices expected to have only one index")
	}
	if indices[0] != 1 {
		t.Error("Non zero index expected to be 1")
	}
}

/* <--------------- Sparse ---------------> */

/* Set & Get Values */
func TestSetNonZeroValueSparse(t *testing.T) {
	m := MakeSparse(5, 5)
	m.SetValue(1, 2, 7.5)

	if val := m.Value(1, 2); val != 7.5 {
		t.Errorf("Value not as expected. Got %f", val)
	}
}

func TestSetZeroValueRemovesValueSparse(t *testing.T) {
	m := MakeSparse(5, 5)
	m.SetValue(1, 2, 7.5)
	m.SetValue(1, 2, 0.0)

	if val := m.Value(1, 2); val != 0.0 {
		t.Errorf("Value not as expected. Got %f", val)
	}
}

func TestNonAssignedValueIsZeroSparse(t *testing.T) {
	m := MakeSparse(2, 4)
	if val := m.Value(1, 3); val != 0 {
		t.Errorf("Value not as expected. Got %f", val)
	}
}

func TestAddToValueSparse(t *testing.T) {
	m := MakeSparse(2, 2)
	m.AddToValue(0, 1, 5.0)
	m.AddToValue(0, 1, 6.0)

	if m.Value(0, 1) != 11.0 {
		t.Error("Value not as expected")
	}
}

func TestSetZeroColSparse(t *testing.T) {
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

func TestSetIdentityRowSparse(t *testing.T) {
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

func TestNonZeroIndicesInRowSparse(t *testing.T) {
	m := MakeSparse(3, 3)
	m.SetValue(1, 1, 4.0)
	indices := m.NonZeroIndicesAtRow(1)

	if len(indices) != 1 {
		t.Error("Non zero indices expected to have only one index")
	}
	if indices[0] != 1 {
		t.Error("Non zero index expected to be 1")
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
