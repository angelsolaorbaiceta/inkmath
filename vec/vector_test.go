package vec

import "testing"

func TestAdd(t *testing.T) {
	u := MakeWithValues([]float64{1, 2})
	v := MakeWithValues([]float64{3, 4})
	sum := u.Plus(v)
	expectedSum := MakeWithValues([]float64{4, 6})

	if !sum.Equals(expectedSum) {
		t.Errorf("Wrong vector sum. Expected %v, but got %v", expectedSum, sum)
	}
}
