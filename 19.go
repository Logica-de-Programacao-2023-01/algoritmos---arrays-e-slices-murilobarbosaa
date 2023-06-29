package main

import (
	"fmt"
)

func main() {
	fmt.Print("Informe o tamanho dos arrays: ")
	var n int
	fmt.Scan(&n)

	array1 := make([]int, n)
	array2 := make([]int, n)

	fmt.Println("Informe os elementos do primeiro array:")
	readArray(array1)

	fmt.Println("Informe os elementos do segundo array:")
	readArray(array2)

	array3 := sumArrays(array1, array2)

	fmt.Println("Terceiro array (soma):", array3)
}

func readArray(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&array[i])
	}
}

func sumArrays(array1, array2 []int) []int {
	sum := make([]int, len(array1))
	for i := 0; i < len(array1); i++ {
		sum[i] = array1[i] + array2[i]
	}
	return sum
}
