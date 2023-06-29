package main

import "fmt"

func main() {
	slice := make([]int, 0, 5)

	var numero int
	fmt.Print("Informe um número inteiro: ")
	fmt.Scan(&numero)

	Presente := false
	for _, valor := range slice {
		if valor == numero {
			Presente = true
			break
		}
	}

	if !Presente {
		slice = append(slice, numero)
	}

	fmt.Println("Slice resultante:", slice)
}
