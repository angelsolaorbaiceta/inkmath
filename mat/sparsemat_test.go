package mat

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkmath"
	"github.com/angelsolaorbaiceta/inkmath/vec"
)

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
