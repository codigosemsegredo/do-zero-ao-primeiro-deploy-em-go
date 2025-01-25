package main

import "fmt"

// Declaração da função com parâmetro
func boasVindas(nome string) {
	fmt.Printf("Olá, %s! Seja bem-vindo ao mundo da programação com Go!\n", nome)
}

func main() {
	boasVindas("João") // Passando "João" como argumento
}
