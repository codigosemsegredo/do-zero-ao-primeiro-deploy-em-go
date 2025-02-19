package main

import "fmt"

func main() {
	var numeros []int
	var numero int

	for i := 1; i <= 5; i++ {
		fmt.Printf("Digite o número %d: ", i)
		fmt.Scanln(&numero)
		numeros = append(numeros, numero)
	}

	soma := 0
	for _, valor := range numeros {
		soma += valor
	}

	fmt.Println("A soma dos números é:", soma)
}
