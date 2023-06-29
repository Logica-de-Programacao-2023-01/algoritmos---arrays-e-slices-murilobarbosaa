package main

import "fmt"

func main() {
	var tamanho int
	fmt.Print("Informe o tamanho do slice: ")
	fmt.Scan(&tamanho)

	slice := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Printf("Informe o valor do elemento %d: ", i+1)
		fmt.Scan(&slice[i])
	}

	fmt.Println("Slice:", slice)

	soma := 0
	for _, valor := range slice {
		soma += valor
	}

	fmt.Println("A soma dos valores é:", soma)
}
