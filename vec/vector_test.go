package vec

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkmath"
)

func TestAdd(t *testing.T) {
	u := MakeWithValues([]float64{1, 2})
	v := MakeWithValues([]float64{3, 4})
	sum := u.Plus(v)
	expectedSum := MakeWithValues([]float64{4, 6})

	if !sum.Equals(expectedSum) {
		t.Errorf("Wrong vector sum. Expected %v, but got %v", expectedSum, sum)
	}
}

func TestSubtract(t *testing.T) {
	u := MakeWithValues([]float64{1, 2})
	v := MakeWithValues([]float64{5, 4})
	sub := u.Minus(v)
	expectedSub := MakeWithValues([]float64{-4, -2})

	if !sub.Equals(expectedSub) {
		t.Errorf("Wrong vector sum. Expected %v, but got %v", expectedSub, sub)
	}
}

func TestMultiply(t *testing.T) {
	u := MakeWithValues([]float64{1, 2})
	v := MakeWithValues([]float64{3, 4})
	prod := u.Times(v)
	expectedProd := 11.0

	if !inkmath.FuzzyEqual(prod, expectedProd) {
		t.Errorf("Wrong vector prod. Expected %f, but got %f", expectedProd, prod)
	}
}
