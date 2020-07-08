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

	"github.com/angelsolaorbaiceta/inkmath/nums"
)

func TestMakeDenseMatrixInitializedWithZeroes(t *testing.T) {
	matrix := MakeSquareDense(3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if val := matrix.Value(i, j); !nums.IsCloseToZero(val) {
				t.Errorf("Expected value to be 0.0, but got %f", val)
			}
		}
	}
}

func TestDenseMatrixValues(t *testing.T) {
	matrix := MakeSquareDense(3)
	matrix.SetValue(1, 2, 34.0)

	if val := matrix.Value(1, 2); val != 34.0 {
		t.Errorf("Expected value to be 34.0, got %f", val)
	}
}

func TestDenseMatrixRowsAndCols(t *testing.T) {
	matrix := MakeDense(3, 5)

	if matrix.Rows() != 3 {
		t.Errorf("Expected 3 rows, got %d", matrix.Rows())
	}
	if matrix.Cols() != 5 {
		t.Errorf("Expected 5 rows, got %d", matrix.Cols())
	}
}

func TestDenseMatrixAddToValue(t *testing.T) {
	matrix := MakeDenseWithData(2, 3, []float64{1, 2, 3, 4, 5, 6})
	matrix.AddToValue(1, 2, 10.0)

	if val := matrix.Value(1, 2); val != 16.0 {
		t.Errorf("Expected 16.0, but got %f", val)
	}
}

func TestDenseSetZeroColumn(t *testing.T) {
	var (
		matrix   = MakeDenseWithData(2, 3, []float64{1, 2, 3, 4, 5, 6})
		wantData = []float64{
			1, 0, 3,
			4, 0, 6,
		}
	)
	matrix.SetZeroCol(1)

	if !matrixContainsData(matrix, wantData) {
		t.Errorf("Matrix contains the wrong data. Want %v, got %v", wantData, matrix)
	}
}

func TestDenseSetIdentityRow(t *testing.T) {
	var (
		matrix   = MakeDenseWithData(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
		wantData = []float64{
			1, 2, 3,
			0, 1, 0,
			7, 8, 9,
		}
	)
	matrix.SetIdentityRow(1)

	if !matrixContainsData(matrix, wantData) {
		t.Errorf("Matrix contains the wrong data. Want %v, got %v", wantData, matrix)
	}
}
