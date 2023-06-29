package main

import (
	"fmt"
)

func main() {
	array := [10]float64{2.5, 6.7, 3.2, 8.9, 4.1, 9.3, 5.7, 1.2, 7.6, 0.9}

	slice := make([]float64, 0)
	for _, value := range array {
		if value > 5 {
			slice = append(slice, value)
		}
	}

	fmt.Println("Novo slice:", slice)
}
