package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 8)

	fmt.Println("Informe dois índices de elementos para troca:")
	var index1, index2 int
	fmt.Print("Índice 1: ")
	fmt.Scan(&index1)
	fmt.Print("Índice 2: ")
	fmt.Scan(&index2)

	if index1 < 0 || index1 >= len(slice) || index2 < 0 || index2 >= len(slice) {
		fmt.Println("Índices inválidos!")
		return
	}

	slice[index1], slice[index2] = slice[index2], slice[index1]

	fmt.Println("Slice resultante:", slice)
}
