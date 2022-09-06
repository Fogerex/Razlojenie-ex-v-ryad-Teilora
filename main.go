package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		e  float64 = 1 //Степень
		t  float64 = 5 //Точность
		ta float64 = 1 //ta верхняя часть уравнения для а
		la float64 = 1 //la нижняя часть уравнения для а
		a  float64 = 1 //Необходимое зло
	)
	fmt.Println("Введи Степень:")
	fmt.Scan(&e)
	fmt.Println("Введи Точность:")
	fmt.Scan(&t)
	t = math.Pow(0.1, t)
	n := 1.0
	for a > t {
		n++
		ta++
		la *= n
		a = ta / la
		e += a
		fmt.Printf("n: %f; ta: %f; la: %f; a: %f; e: %f\n", n, ta, la, a, e)

	}
	fmt.Printf("e=%f при точности = %f и n=%f", e, t, n)
	fmt.Printf("\nEND n: %f; ta: %f; la: %f; a: %f; e: %f", n, ta, la, a, e)

}
