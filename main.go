package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		x float64
		n float64
	)

	fmt.Println("Введи x:")
	fmt.Scan(&x)
	fmt.Println("Введи n:")
	fmt.Scan(&n)
	stx := x    //{x в степени}
	fakt := 1.0 //{значение факториала в знаменателе}
	ex := 1.0   //{первый элемент в разложении}
	n = 1.0     //{счётчик}
	e := math.E //{число e}
	fmt.Println("e:", e)

	for (stx / fakt) >= e {
		ex += stx / fakt
		n++
		stx *= x
		fakt *= n

	}
	fmt.Println("Количество элементов в разложении = ", n+1)
	fmt.Println("Значение e^", x, "= ", ex)
	fmt.Println("Значение компьютера: ", math.Exp(x))

}
