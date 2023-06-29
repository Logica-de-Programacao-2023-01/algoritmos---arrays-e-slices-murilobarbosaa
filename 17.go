package main

import (
	"fmt"
)

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := 0
	for i, value := range array {
		if i%2 == 0 {
			sum += value
		}
	}

	fmt.Println("Soma dos elementos nas posições pares:", sum)
}
