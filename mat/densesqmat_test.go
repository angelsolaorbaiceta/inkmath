package mat

import "testing"

func TestMultiplyInPlace(t *testing.T) {
	matA, matB := MakeSquareDense(2), MakeSquareDense(2)

	matA.SetValue(0, 0, 2.0)
	matA.SetValue(1, 0, 1.0)
	matA.SetValue(1, 1, 2.0)

	matB.SetValue(0, 0, 1.0)
	matB.SetValue(0, 1, 2.0)
	matB.SetValue(1, 0, 3.0)
	matB.SetValue(1, 1, 4.0)

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
