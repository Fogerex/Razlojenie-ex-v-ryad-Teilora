package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		e  float64     //= 1 //Степень
		t  float64     //= 5 //Точность
		ta float64 = 1 //ta верхняя часть уравнения для а
		la float64 = 1 //la нижняя часть уравнения для а
		a  float64 = 1 //Отношение ta / la
		n  int     = 1 //Факториал
	)
	fmt.Println("Введи Степень:")
	fmt.Scan(&e)
	fmt.Println("Введи Точность:")
	fmt.Scan(&t)
	t = math.Pow(0.1, t)
	if e == 0 {
		e = math.Pow(math.E, 0)
		fmt.Println("Ты зочем пидр вводишь нули???")
		fmt.Printf("e=%v при точности = %f и n=%d. Проверочная e вычесленная ПК: %f\n", e, t, n, math.Pow(math.E, 0))
	} else {
		for a > t {
			n++
			ta++
			la *= float64(n)
			a = ta / la
			e += a
			//fmt.Printf("n: %d; ta: %f; la: %f; a: %f; e: %v\n", n, ta, la, a, e)
		}
		fmt.Printf("e=%v при точности = %f и n=%d. Проверочная e вычесленная ПК: %f\n", e, t, n, math.Pow(math.E, t))
	}
}
