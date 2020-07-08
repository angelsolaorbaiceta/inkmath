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
	v := MakeWithValues([]float64{1, 2, 3})
	expectedNorm := math.Sqrt(14.0)

	if norm := v.Norm(); !nums.FuzzyEqual(norm, expectedNorm) {
		t.Errorf("Wrong Vector norm. Expected %f, but got %f", expectedNorm, norm)
	}
}

/* <--------------- Methods ---------------> */
func TestOpposite(t *testing.T) {
	v := MakeWithValues([]float64{1, 2, 3})
	expectedOpposite := MakeWithValues([]float64{-1, -2, -3})

	if !v.Opposite().Equals(expectedOpposite) {
		t.Error("Opposite vector not as expected")
	}
}

/* <--------------- Operations ---------------> */
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

	if !nums.FuzzyEqual(prod, expectedProd) {
		t.Errorf("Wrong vector prod. Expected %f, but got %f", expectedProd, prod)
	}
}
