package main

import (
	"fmt"
	"strconv"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: " + strconv.FormatFloat(float64(e), 'f', -1, 64)
}

const kfDeltaThreshold = 0.00000000001

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0.0, ErrNegativeSqrt(x)
	}

	z := x
	oldz := x + 1.0
	//fmt.Println(z)

	for i := 1; i <= 100; i++ {
		//fmt.Println("iteration #", i)
		if oldz-z < kfDeltaThreshold {
			break
		}
		oldz = z
		z = z - ((z*z - x) / (2 * z))
		//fmt.Println(z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
