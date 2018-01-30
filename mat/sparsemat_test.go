package mat

import (
	"testing"

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

/* Operations */
func TestSparseMatrixTimesVector(t *testing.T) {
	mat := MakeSparse(2, 2)
	mat.SetValue(0, 0, 1.0)
	mat.SetValue(0, 1, 2.0)
	mat.SetValue(1, 0, 3.0)
	mat.SetValue(1, 1, 4.0)
	v := vec.MakeWithValues([]float64{5.0, 6.0})
	prod, _ := mat.TimesVector(v)
	expectedVec := vec.MakeWithValues([]float64{17.0, 39.0})

	if !prod.Equals(expectedVec) {
		t.Errorf("Wrong multiplication. Expected %v but got %v", expectedVec, prod)
	}
}
