package main

import "fmt"

func main() {
	cep := "59190000"
	ini := "01000-000"
	fim := "05000-000"

	if ini <= cep && fim >= cep {
		fmt.Println("O CEP", cep, "está entre", ini, "e", fim)
	} else {
		fmt.Println("O CEP", cep, "não está entre", ini, "e", fim)
	}
}
