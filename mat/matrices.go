package mat

import "github.com/angelsolaorbaiceta/inkmath"

func IsSquare(m Matrixable) bool {
	return m.Rows() == m.Cols()
}

func IsSymmetric(m Matrixable) bool {
	if !IsSquare(m) {
		return false
	}

	for i := 0; i < m.Rows(); i++ {
		for j := i + 1; j < m.Cols(); j++ {
			if !inkmath.FuzzyEqual(m.Value(i, j), m.Value(j, i)) {
				return false
			}
		}
	}

	return true
}
