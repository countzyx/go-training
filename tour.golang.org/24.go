package main

import (
	"fmt"
)

const kfDeltaThreshold = 0.00000000001

func Sqrt(x float64) float64 {
	z := x
	oldz := x + 1.0
	//    fmt.Println(z)

	for i := 1; i <= 100; i++ {
		fmt.Println("iteration #", i)
		if oldz-z < kfDeltaThreshold {
			break
		}
		oldz = z
		z = z - ((z*z - x) / (2 * z))
		//        fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(10000))
}
