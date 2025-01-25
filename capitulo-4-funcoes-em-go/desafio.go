package main

import "fmt"

func saudacao(nome string, periodo string) string {
	return fmt.Sprintf("Boa %s, %s!", periodo, nome)
}

func main() {
	mensagem := saudacao("Maria", "tarde")
	fmt.Println(mensagem)
}
