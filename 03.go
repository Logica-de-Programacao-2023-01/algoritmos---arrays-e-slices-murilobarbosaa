package main

import "fmt"

func main() {
	array := [4]float64{2.5, 1.5, 3.0, 2.0}

	produto := array[0]

	for i := 1; i < len(array); i++ {
		produto *= array[i]
	}

	fmt.Println("O produto dos valores é:", produto)
}
