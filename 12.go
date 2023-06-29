package main

import "fmt"

func main() {
	array := [5]int{10, 15, 8, 9, 12}

	slice := make([]int, 0)
	for _, elemento := range array {
		if elemento%3 == 0 {
			slice = append(slice, elemento)
		}
	}

	fmt.Println("Novo slice:", slice)
}
