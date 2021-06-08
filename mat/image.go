package mat

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/angelsolaorbaiceta/inkmath/nums"
)

var (
	gray = color.RGBA{230, 230, 230, 255}
	red  = color.RGBA{255, 0, 0, 255}
	blue = color.RGBA{0, 0, 255, 255}
)

/*
ToImage creates an image with as many width pixels as columns has the matrix and as many
height pixels as rows.

Each pixel will be colored:
	- Gray if the value is zero
	- Red if the value is positive
	- Blue if the value is negative
*/
func ToImage(m ReadOnlyMatrix, filePath string) {
	var (
		width  = m.Cols()
		height = m.Rows()
		img    = image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
		val    float64
	)

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			val = m.Value(row, col)
			if nums.IsCloseToZero(val) {
				img.Set(row, col, gray)
			} else if val > 0.0 {
				img.Set(row, col, red)
			} else {
				img.Set(row, col, blue)
			}
		}
	}

	f, _ := os.OpenFile(filePath+"_sysmat.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
