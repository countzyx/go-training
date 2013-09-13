package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	aiFib := []int{0, 0, 1}

	return func() int {
		aiFib[0] = aiFib[1]
		aiFib[1] = aiFib[2]
		aiFib[2] = aiFib[0] + aiFib[1]
		return aiFib[0]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
