package main

import "fmt"

// Função que retorna a soma e o produto de dois números
func operacoes(a int, b int) (int, int) {
	soma := a + b
	produto := a * b
	return soma, produto
}

func main() {
	soma, produto := operacoes(3, 4)
	fmt.Println("Soma:", soma)
	fmt.Println("Produto:", produto)
}
