package main

import "fmt"

// Função que soma dois números
func somar(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("A soma de 3 e 4 é:", somar(3, 4))
}
