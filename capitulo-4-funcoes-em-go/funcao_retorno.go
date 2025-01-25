package main

import "fmt"

// Função que retorna um valor
func dobrar(numero int) int {
	return numero * 2
}

func main() {
	resultado := dobrar(5) // Armazena o retorno da função
	fmt.Println("O dobro de 5 é:", resultado)
}
