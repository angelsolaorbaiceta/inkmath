package nums

import "testing"

func TestTwoNumbersEqual(t *testing.T) {
	if equal := FuzzyEqualEps(1.001, 1.002, 0.01); !equal {
		t.Error("Expected float64 values to be 'fuzzy' equal")
	}
}

func TestTwoNumbersNotEqual(t *testing.T) {
	if equal := FuzzyEqualEps(1.001, 1.002, 0.0001); equal {
		t.Error("Expected float64 values to be 'fuzzy' not equal")
	}
}

func TestLinearInterpolation(t *testing.T) {
	val := LinInterpol(1.0, 1.0, 3.0, 3.0, 2.0)
	if !FuzzyEqual(val, 2.0) {
		t.Error("Wrong linear interpolation value")
	}
}
