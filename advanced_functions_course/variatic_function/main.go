package main

import "fmt"

// soma - Soma os números dados e imprime o texto.
//
// Args:
// texto - O texto a ser impresso.
// numeros - Os números a serem somados.
//
// Returns:
// total - A soma dos números dados.
func soma(texto string, numeros ...int) (total int) {
	total = 0
	for _, num := range numeros {
		total += num
	}
	fmt.Println(texto)
	return
}
// main - Função principal do programa.
//
// Imprime o resultado da soma de uma lista de números dados.
func main() {
	fmt.Println(soma("Teste", 2, 6, 7, 9, 105, 202))
}
