package main

import "fmt"

func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	var numero int
	fmt.Print("Informe um número: ")
	fmt.Scan(&numero)

	array[0] += numero
	array[len(array)-1] += numero

	fmt.Println("Array resultante:", array)
}
