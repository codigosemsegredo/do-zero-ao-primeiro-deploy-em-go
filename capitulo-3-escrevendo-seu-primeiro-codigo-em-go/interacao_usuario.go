package main

import "fmt"

func main() {
	var nome string
	fmt.Println("Digite seu nome:")
	fmt.Scanln(&nome)
	fmt.Printf("Olá, %s! Seja bem-vindo ao mundo da programação!\n", nome)
}
