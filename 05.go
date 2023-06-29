package main

import (
	"fmt"
)

func main() {
	matriz := [3][2]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Informe o valor do elemento [%d][%d]: ", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}

	fmt.Println("Matriz:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println()
	}
}
