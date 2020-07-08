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
