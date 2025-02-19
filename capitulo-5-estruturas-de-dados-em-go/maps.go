package main

import "fmt"

func main() {
	telefones := make(map[string]string) // Cria um map com chave e valor do tipo string
	telefones["João"] = "1234-5678"
	telefones["Maria"] = "9876-5432"

	fmt.Println("Telefone do João:", telefones["João"])
}
