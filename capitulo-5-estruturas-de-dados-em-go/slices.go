package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5} // Criando um slice com 5 números
	fmt.Println(numeros)

	numeros = append(numeros, 6) // Adicionando um novo elemento
	fmt.Println(numeros)
}
