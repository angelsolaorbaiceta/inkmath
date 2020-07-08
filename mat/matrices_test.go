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

func TestIsNotSymmetric(t *testing.T) {
	m := MakeSparse(5, 5)
	m.SetValue(3, 4, 9.0)

	if IsSymmetric(m) {
		t.Error("Expected matrix not to be symmetric")
	}
}

/* <--------------- Dominant ---------------> */
func TestIsRowDominant(t *testing.T) {
	m := MakeSparse(2, 2)
	m.SetValue(0, 0, 50.0)
	m.SetValue(0, 1, 10.0)
	m.SetValue(1, 1, -50.0)
	m.SetValue(1, 0, 10.0)

	if !IsRowDominant(m) {
		t.Error("Expected matrix to be row dominant")
	}
}

func TestIsNotRowDominant(t *testing.T) {
	m := MakeSparse(2, 2)
	m.SetValue(0, 0, 50.0)
	m.SetValue(0, 1, 10.0)
	m.SetValue(1, 1, 10.0)
	m.SetValue(1, 0, 50.0)

	if IsRowDominant(m) {
		t.Error("Expected matrix not to be row dominant")
	}
}
