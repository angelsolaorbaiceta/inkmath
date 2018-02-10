package mat

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/angelsolaorbaiceta/inkmath"
)

/*
IsSquare returns true if the given matrix has the same number of rows and columns.
*/
func IsSquare(m Matrixable) bool {
	return m.Rows() == m.Cols()
}

/*
IsSymmetric returns true if the given matrix is square and equals to it's traspose.
*/
func IsSymmetric(m Matrixable) bool {
	if !IsSquare(m) {
		panic("Matrix symmetry only applies to square matrices")
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

/*
IsRowDominant returns true if for every row in the matrix, the element in the main diagonal
is greater than every other element.
*/
func IsRowDominant(m Matrixable) bool {
	if !IsSquare(m) {
		panic("Matrix dominancy only applies to square matrices")
	}

	var diagonalValue float64
	for i := 0; i < m.Rows(); i++ {
		diagonalValue = math.Abs(m.Value(i, i))
		for j := 0; j < m.Cols(); j++ {
			if i != j && diagonalValue < math.Abs(m.Value(i, j)) {
				return false
			}
		}
	}

	return true
}

/*
HasZeroInMainDiagonal returns true if a zero is found in the matrix main diagonal.
*/
func HasZeroInMainDiagonal(m Matrixable) bool {
	if !IsSquare(m) {
		panic("Matrix main diagonal only applies to square matrices")
	}

	for i := 0; i < m.Rows(); i++ {
		if inkmath.IsCloseToZero(m.Value(i, i)) {
			return true
		}
	}

	return false
}

/* ::::::::::::::: Image ::::::::::::::: */

/*
ToImage creates an image with as many width pixels as columns has the matrix and
as many height pixels as rows. Each pixel will be coloured:
	- Gray if matrix value is zero
	- Red if matrix value is positive
	- Blue if matrix value is negative
*/
func ToImage(m Matrixable, path, fileName string) {
	var (
		width     = m.Cols()
		height    = m.Rows()
		img       = image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
		zeroColor = color.RGBA{230, 230, 230, 255}
		posColor  = color.RGBA{255, 0, 0, 255}
		negColor  = color.RGBA{0, 0, 255, 255}
		val       float64
	)

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			val = m.Value(row, col)
			if inkmath.IsCloseToZero(val) {
				img.Set(row, col, zeroColor)
			} else if val > 0.0 {
				img.Set(row, col, posColor)
			} else {
				img.Set(row, col, negColor)
			}
		}
	}

	f, _ := os.OpenFile(fmt.Sprintf("%s/%s.png", path, fileName), os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
