package main

import (
	"fmt"
	_ "fmt"
	_ "math"
)

func main() {
	var (
		e  int     = 1
		t  float64 = 0.001
		ta float64 = 1
		la float64 = 1
		a  float64 = 1
	)
	n := 1
	for a > t {
		n++
		ta++
		la *= float64(n)
		a = ta / la
		e += int(a)
		fmt.Printf("n: %d; ta: %f; la: %f; a: %f; e: %d\n", n, ta, la, a, e)
		// if a > t {
		// 	continue
		// }
	}
	fmt.Printf("\nEND n: %d; ta: %f; la: %f; a: %f; e: %d", n, ta, la, a, e)

}
