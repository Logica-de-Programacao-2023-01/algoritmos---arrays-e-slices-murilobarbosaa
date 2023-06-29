package main

import "fmt"

func main() {
	slice := []string{"valor1", "valor2", "valor3", "valor2", "valor4", "valor5", "valor2", "valor6"}

	var valor string
	fmt.Print("Informe um valor: ")
	fmt.Scan(&valor)

	novoSlice := make([]string, 0, len(slice))
	for _, elemento := range slice {
		if elemento != valor {
			novoSlice = append(novoSlice, elemento)
		}
	}

	fmt.Println("Slice resultante:", novoSlice)
}
