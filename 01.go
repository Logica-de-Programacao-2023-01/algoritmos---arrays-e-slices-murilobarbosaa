package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}

	soma := 0

	for _, valor := range array {
		soma += valor
	}

	fmt.Println("A soma dos valores é: ", soma)
}
