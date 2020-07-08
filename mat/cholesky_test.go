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

import (
	"testing"
)

func TestCholeskyDecomposition(t *testing.T) {
	m, expectedDecomposition := makeCholeskyMatrix(), makeCholeskyDecomposition()
	cholesky := CholeskyDecomposition(m)

	if !AreEqual(cholesky, expectedDecomposition) {
		t.Error("Wrong Cholesky factorization")
	}
}

func TestIncompleteCholeskyDecomposition(t *testing.T) {
	m, expectedDecomposition := makeCholeskyMatrix(), makeCholeskyDecomposition()
	cholesky := IncompleteCholeskyDecomposition(m)

	if !AreEqual(cholesky, expectedDecomposition) {
		t.Error("Wrong Cholesky factorization")
	}
}

func makeCholeskyMatrix() ReadOnlyMatrix {
	m := MakeSparse(4, 4)

	m.SetValue(0, 0, 4.0)
	m.SetValue(0, 1, -2.0)
	m.SetValue(0, 2, 4.0)
	m.SetValue(0, 3, 2.0)

	m.SetValue(1, 0, -2.0)
	m.SetValue(1, 1, 10.0)
	m.SetValue(1, 2, -2.0)
	m.SetValue(1, 3, -7.0)

	m.SetValue(2, 0, 4.0)
	m.SetValue(2, 1, -2.0)
	m.SetValue(2, 2, 8.0)
	m.SetValue(2, 3, 4.0)

	m.SetValue(3, 0, 2.0)
	m.SetValue(3, 1, -7.0)
	m.SetValue(3, 2, 4.0)
	m.SetValue(3, 3, 7.0)

	return m
}

func makeCholeskyDecomposition() ReadOnlyMatrix {
	m := MakeSquareDense(4)

	m.SetValue(0, 0, 2.0)

	m.SetValue(1, 0, -1.0)
	m.SetValue(1, 1, 3.0)

	m.SetValue(2, 0, 2.0)
	m.SetValue(2, 2, 2.0)

	m.SetValue(3, 0, 1.0)
	m.SetValue(3, 1, -2.0)
	m.SetValue(3, 2, 1.0)
	m.SetValue(3, 3, 1.0)

	return m
}
