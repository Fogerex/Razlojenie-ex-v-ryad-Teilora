package main

import (
	"fmt"
	"math"
)

func main() {

	// Разложение ex в ряд Тейлора
	// Задача:
	// Используя разложение в ряд Тейлора найти значение ex с заданной точностью e.
	// Описание:
	// Код:

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
	i := 1.0    //{счётчик}
	//e := math.E //{число e}
	//fmt.Println("e:", e)

	for (stx / fakt) >= n {
		ex += stx / fakt
		i++
		stx *= x
		fakt *= i
		fmt.Println("i:", i)

	}
	fmt.Println("Количество элементов в разложении = ", i+1)
	fmt.Println("Значение e^", x, "= ", ex)
	fmt.Println("Значение компьютера: ", math.Exp(x))

}
