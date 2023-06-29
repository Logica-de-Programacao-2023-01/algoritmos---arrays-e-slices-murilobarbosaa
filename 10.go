package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []int{9, 2, 6, 1, 5, 8, 3, 7, 4, 10}

	min := math.MaxInt64
	max := math.MinInt64
	for _, valor := range slice {
		if valor < min {
			min = valor
		}
		if valor > max {
			max = valor
		}
	}

	fmt.Println("Valor mínimo:", min)
	fmt.Println("Valor máximo:", max)
}
