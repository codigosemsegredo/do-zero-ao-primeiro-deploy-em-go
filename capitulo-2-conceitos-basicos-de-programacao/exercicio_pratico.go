package main

import "fmt"

func main() {
	var nome string
	var idade int

	fmt.Println("Digite seu nome:")
	fmt.Scanln(&nome)

	fmt.Println("Digite sua idade:")
	fmt.Scanln(&idade)

	if idade >= 18 {
		fmt.Printf("Olá, %s. Você pode dirigir.\n", nome)
	} else {
		fmt.Printf("Olá, %s. Você ainda não pode dirigir.\n", nome)
	}
}
