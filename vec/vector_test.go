/*
Copyright 2020 Angel Sola Orbaiceta

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vec

import (
	"math"
	"testing"

	"github.com/angelsolaorbaiceta/inkmath/nums"
)

/* <--------------- Properties ---------------> */
func TestNorm(t *testing.T) {
	var (
		v    = MakeWithValues([]float64{1, 2, 3})
		want = math.Sqrt(14.0)
	)

	if norm := v.Norm(); !nums.FuzzyEqual(norm, want) {
		t.Errorf("Wrong Vector norm. Expected %f, but got %f", want, norm)
	}
}

func TestLength(t *testing.T) {
	v := Make(3)
	if length := v.Length(); length != 3 {
		t.Errorf("Expected %d to be %d", length, 3)
	}
}

func TestSetGetValue(t *testing.T) {
	v := Make(3)
	v.SetValue(1, 25)

	if value := v.Value(1); value != 25 {
		t.Errorf("Expected %f to be 25", value)
	}
}

/* <--------------- Methods ---------------> */
func TestOpposite(t *testing.T) {
	var (
		v    = MakeWithValues([]float64{1, 2, 3})
		want = MakeWithValues([]float64{-1, -2, -3})
	)

	if !v.Opposite().Equals(want) {
		t.Error("Opposite vector not as expected")
	}
}

func TestEquals(t *testing.T) {
	t.Run("vectors with differnet sizes aren't equal", func(t *testing.T) {
		var (
			u = Make(2)
			v = Make(3)
		)

		if u.Equals(v) {
			t.Errorf("Expected %v and %v to be different", u, v)
		}
	})

	t.Run("vectors are't equal", func(t *testing.T) {
		var (
			u = MakeWithValues([]float64{1, 2, 3})
			v = MakeWithValues([]float64{4, 5, 6})
		)

		if u.Equals(v) {
			t.Errorf("Expected %v and %v to be different", u, v)
		}
	})

	t.Run("vectors are equal", func(t *testing.T) {
		var (
			u = MakeWithValues([]float64{1, 2, 3})
			v = MakeWithValues([]float64{1, 2, 3})
		)

		if !u.Equals(v) {
			t.Errorf("Expected %v and %v to be equal", u, v)
		}
	})
}

func TestScaled(t *testing.T) {
	var (
		u    = MakeWithValues([]float64{1, 2})
		want = MakeWithValues([]float64{3, 6})
		got  = u.Scaled(3)
	)

	if !got.Equals(want) {
		t.Errorf("Expected %v, but got %v", got, want)
	}
}

func TestClone(t *testing.T) {
	u := MakeWithValues([]float64{1, 2, 3})
	if got := u.Clone(); !got.Equals(u) {
		t.Errorf("Want %v, but got %v", u, got)
	}
}

/* <--------------- Operations ---------------> */
func TestAdd(t *testing.T) {
	var (
		u    = MakeWithValues([]float64{1, 2})
		v    = MakeWithValues([]float64{3, 4})
		got  = u.Plus(v)
		want = MakeWithValues([]float64{4, 6})
	)

	if !got.Equals(want) {
		t.Errorf("Wrong vector sum. Expected %v, but got %v", want, got)
	}
}

func TestSubtract(t *testing.T) {
	var (
		u    = MakeWithValues([]float64{1, 2})
		v    = MakeWithValues([]float64{5, 4})
		got  = u.Minus(v)
		want = MakeWithValues([]float64{-4, -2})
	)

	if !got.Equals(want) {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}

func TestMultiply(t *testing.T) {
	var (
		u    = MakeWithValues([]float64{1, 2})
		v    = MakeWithValues([]float64{3, 4})
		got  = u.Times(v)
		want = 11.0
	)

	if !nums.FuzzyEqual(got, want) {
		t.Errorf("Expected %f, but got %f", want, got)
	}
}
