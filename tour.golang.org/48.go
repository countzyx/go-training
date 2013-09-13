package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := x

	for i := 1; i <= 100; i++ {
		z = z - ((cmplx.Pow(z, 3.0) - x) / (3 * cmplx.Pow(z, 2.0)))
	}

	return z

}

func main() {
	fmt.Println(Cbrt(8))
}
