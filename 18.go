package main

import (
	"fmt"
)

func main() {
	fmt.Print("Informe um número inteiro positivo: ")
	var n int
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Número inválido. O número deve ser inteiro positivo.")
		return
	}

	primes := make([]int, 0)
	num := 2
	for len(primes) < n {
		if isPrime(num) {
			primes = append(primes, num)
		}
		num++
	}

	fmt.Println("Números primos:", primes)
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
