package main

import "fmt"

func main() {
	matriz := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	var linha, coluna int
	fmt.Print("Informe o índice da linha: ")
	fmt.Scan(&linha)
	fmt.Print("Informe o índice da coluna: ")
	fmt.Scan(&coluna)

	if linha >= 0 && linha < len(matriz) && coluna >= 0 && coluna < len(matriz[0]) {
		valor := matriz[linha][coluna]
		fmt.Printf("Valor armazenado na posição [%d][%d]: %d\n", linha, coluna, valor)
	} else {
		fmt.Println("Índices inválidos!")
	}
}
