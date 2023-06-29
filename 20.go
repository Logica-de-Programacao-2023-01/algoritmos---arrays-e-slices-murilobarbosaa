package main

import (
	"fmt"
)

func main() {
	fmt.Println("Informe os elementos do array (separados por espaço):")
	array := readArray()

	// Verificando se o array está ordenado em ordem crescente
	if isSorted(array) {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente.")
	}
}

func readArray() []int {
	var n int
	fmt.Print("Tamanho do array: ")
	fmt.Scan(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&array[i])
	}

	return array
}

func isSorted(array []int) bool {
	n := len(array)
	for i := 1; i < n; i++ {
		if array[i] < array[i-1] {
			return false
		}
	}
	return true
}
