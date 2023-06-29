package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	indice := 2
	slice = append(slice[:indice], slice[indice+1:]...)

	fmt.Println("Slice resultante: ", slice)
}
