package main

import "fmt"

func main() {
	var nome string
	var idade int

	fmt.Println("Digite seu nome:")
	fmt.Scanln(&nome)

	fmt.Println("Digite sua idade:")
	fmt.Scanln(&idade)

	fmt.Printf("Olá, %s! Você tem %d anos e está começando a programar com Go.\n", nome, idade)
}
