package main

import "fmt"

func main() {
	contatos := make(map[string]string)
	var opcao int

	for {
		fmt.Println("1 - Adicionar contato")
		fmt.Println("2 - Buscar contato")
		fmt.Println("3 - Remover contato")
		fmt.Println("4 - Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&opcao)

		if opcao == 1 {
			var nome, telefone string
			fmt.Print("Digite o nome: ")
			fmt.Scanln(&nome)
			fmt.Print("Digite o telefone: ")
			fmt.Scanln(&telefone)
			contatos[nome] = telefone
		} else if opcao == 2 {
			var nome string
			fmt.Print("Digite o nome: ")
			fmt.Scanln(&nome)
			telefone, existe := contatos[nome]
			if existe {
				fmt.Println("Telefone:", telefone)
			} else {
				fmt.Println("Contato não encontrado.")
			}
		} else if opcao == 3 {
			var nome string
			fmt.Print("Digite o nome: ")
			fmt.Scanln(&nome)
			delete(contatos, nome)
			fmt.Println("Contato removido!")
		} else if opcao == 4 {
			break
		} else {
			fmt.Println("Opção inválida!")
		}
	}
}
