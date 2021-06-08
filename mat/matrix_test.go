package mat

import (
	"testing"

	"github.com/angelsolaorbaiceta/inkmath/vec"
)

func TestMakeMatrixInitializedWithZeroes(t *testing.T) {
	wantData := []float64{
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
	}

	t.Run("dense matrix", func(t *testing.T) {
		matrix := MakeSquareDense(3)
		assertMatrixContainsData(t, matrix, wantData)
	})

	t.Run("sparse matrix", func(t *testing.T) {
		matrix := MakeSquareSparse(3)
		assertMatrixContainsData(t, matrix, wantData)
	})
}

func TestMatrixValues(t *testing.T) {
	testSetGetValue := func(matrix MutableMatrix) {
		matrix.SetValue(1, 2, 34.0)

		if val := matrix.Value(1, 2); val != 34.0 {
			t.Errorf("Expected value to be 34.0, got %f", val)
		}
	}

	t.Run("dense matrix", func(t *testing.T) {
		matrix := MakeSquareDense(3)
		testSetGetValue(matrix)
	})

	t.Run("sparse matrix", func(t *testing.T) {
		matrix := MakeSquareSparse(3)
		testSetGetValue(matrix)
	})
}

func TestMatrixRowsAndCols(t *testing.T) {
	testRowsAndColumns := func(matrix ReadOnlyMatrix) {
		if matrix.Rows() != 3 {
			t.Errorf("Expected 3 rows, got %d", matrix.Rows())
		}
		if matrix.Cols() != 5 {
			t.Errorf("Expected 5 rows, got %d", matrix.Cols())
		}
	}

	t.Run("dense matrix", func(t *testing.T) {
		matrix := MakeDense(3, 5)
		testRowsAndColumns(matrix)
	})

	t.Run("sparse matrix", func(t *testing.T) {
		matrix := MakeSparse(3, 5)
		testRowsAndColumns(matrix)
	})
}

func TestMatrixAddToValue(t *testing.T) {
	testAddToValue := func(matrix MutableMatrix) {
		matrix.AddToValue(1, 2, 10.0)

		if val := matrix.Value(1, 2); val != 16.0 {
			t.Errorf("Expected 16.0, but got %f", val)
		}
	}

	t.Run("dense matrix", func(t *testing.T) {
		matrix := MakeDenseWithData(2, 3, []float64{1, 2, 3, 4, 5, 6})
		testAddToValue(matrix)
	})

	t.Run("sparse matrix", func(t *testing.T) {
		matrix := MakeSparseWithData(2, 3, []float64{1, 2, 3, 4, 5, 6})
		testAddToValue(matrix)
	})
}

func TestMatrixSetZeroColumn(t *testing.T) {
	wantData := []float64{
		1, 0, 3,
		4, 0, 6,
	}
	testSetZeroCol := func(matrix MutableMatrix) {
		matrix.SetZeroCol(1)
		assertMatrixContainsData(t, matrix, wantData)
	}

	t.Run("dense matrix", func(t *testing.T) {
		matrix := MakeDenseWithData(2, 3, []float64{1, 2, 3, 4, 5, 6})
		testSetZeroCol(matrix)
	})

	t.Run("sparse matrix", func(t *testing.T) {
		matrix := MakeSparseWithData(2, 3, []float64{1, 2, 3, 4, 5, 6})
		testSetZeroCol(matrix)
	})
}

func TestMatrixSetIdentityRow(t *testing.T) {
	wantData := []float64{
		1, 2, 3,
		0, 1, 0,
		7, 8, 9,
	}
	testSetIdentity := func(matrix MutableMatrix) {
		matrix.SetIdentityRow(1)
		assertMatrixContainsData(t, matrix, wantData)
	}

	t.Run("dense matrix", func(t *testing.T) {
		matrix := MakeDenseWithData(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
		testSetIdentity(matrix)
	})

	t.Run("sparse matrix", func(t *testing.T) {
		matrix := MakeSparseWithData(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
		testSetIdentity(matrix)
	})
}

func TestMatrixNonZeroIndicesInRow(t *testing.T) {
	testNonZeroIndicesAtRow := func(matrix MutableMatrix) {
		matrix.SetValue(1, 1, 4.0)
		indices := matrix.NonZeroIndicesAtRow(1)

		if len(indices) != 1 {
			t.Error("Non zero indices expected to have only one index")
		}
		if indices[0] != 1 {
			t.Error("Non zero index expected to be 1")
		}
	}

	t.Run("dense matrix", func(t *testing.T) {
		matrix := MakeDense(2, 3)
		testNonZeroIndicesAtRow(matrix)
	})

	t.Run("sparse matrix", func(t *testing.T) {
		matrix := MakeSparse(2, 3)
		testNonZeroIndicesAtRow(matrix)
	})
}

func TestMultiplyMatrices(t *testing.T) {
	wantData := []float64{-62, 24, -52, 56}
	testMultiplyMatrices := func(matA, matB ReadOnlyMatrix) {
		got := matA.TimesMatrix(matB)
		assertMatrixContainsData(t, got, wantData)
	}

	t.Run("dense matrix", func(t *testing.T) {
		var (
			matA = MakeDenseWithData(2, 3, []float64{8, 1, 2, -5, 6, 7})
			matB = MakeDenseWithData(3, 2, []float64{-5, 1, 0, 2, -11, 7})
		)
		testMultiplyMatrices(matA, matB)
	})

	t.Run("sparse matrix", func(t *testing.T) {
		var (
			matA = MakeSparseWithData(2, 3, []float64{8, 1, 2, -5, 6, 7})
			matB = MakeSparseWithData(3, 2, []float64{-5, 1, 0, 2, -11, 7})
		)
		testMultiplyMatrices(matA, matB)
	})
}

func TestMatrixTimesVector(t *testing.T) {
	var (
		wantData = []float64{16, 28}
		vect     = vec.MakeWithValues([]float64{1, 2, 3})
	)
	testMatrixTimesVector := func(mat MutableMatrix) {
		if got := mat.TimesVector(vect); !vec.VectorContainsData(got, wantData) {
			t.Errorf("Wrong vector. Want %v, got %v", wantData, got)
		}
	}

	t.Run("dense matrix", func(t *testing.T) {
		mat := MakeDenseWithData(2, 3, []float64{8, 1, 2, -5, 6, 7})
		testMatrixTimesVector(mat)
	})

	t.Run("sparse matrix", func(t *testing.T) {
		mat := MakeSparseWithData(2, 3, []float64{8, 1, 2, -5, 6, 7})
		testMatrixTimesVector(mat)
	})
}

/* <-- UTILS --> */

func assertMatrixContainsData(t *testing.T, got ReadOnlyMatrix, wantData []float64) {
	if !MatrixContainsData(got, wantData) {
		t.Errorf("Matrix contains wrong data. Want %v, got %v", wantData, got)
	}
}
