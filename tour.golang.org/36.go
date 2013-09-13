package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	yPicMatrix := make([][]uint8, dy)
	for iY := range yPicMatrix {
		yPicMatrix[iY] = make([]uint8, dx)
		for iX := range yPicMatrix[iY] {
			//yPicMatrix[iY][iX] = uint8(iY * iX) // wavy fractal pattern
			//yPicMatrix[iY][iX] = uint8((iY + iX) / 2) // gradient
			yPicMatrix[iY][iX] = uint8(iX ^ iY) // cool square fractal pattern
		}
	}

	return yPicMatrix
}

func main() {
	pic.Show(Pic)
}
