package mat

import "testing"

/* <--------------- Is Square ---------------> */
func TestIsSquare(t *testing.T) {
	m := MakeSquareDense(3)
	if !IsSquare(m) {
		t.Error("Expected matrix to be square")
	}
}

func TestIsNotSquare(t *testing.T) {
	m := MakeSparse(1, 3)
	if IsSquare(m) {
		t.Error("Expected matrix not to be square")
	}
}

/* <--------------- Is Symmetric ---------------> */
func TestIsSymmetric(t *testing.T) {
	m := MakeSparse(5, 5)
	m.SetValue(1, 4, 3.0)
	m.SetValue(4, 1, 3.0)
	m.SetValue(2, 3, 9.0)
	m.SetValue(3, 2, 9.0)

	if !IsSymmetric(m) {
		t.Error("Expected matrix to be symmetric")
	}
}

func TestNonSquareIsNotSymmetric(t *testing.T) {
	m := MakeSparse(1, 3)
	if IsSymmetric(m) {
		t.Error("Expected matrix not to be symmetric")
	}
}

func TestIsNotSymmetric(t *testing.T) {
	m := MakeSparse(5, 5)
	m.SetValue(3, 4, 9.0)

	if IsSymmetric(m) {
		t.Error("Expected matrix not to be symmetric")
	}
}
