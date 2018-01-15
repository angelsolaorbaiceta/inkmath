package mat

import "testing"

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
