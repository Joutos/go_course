package main

import (
	"fmt"
)

// Calcula o n-ésimo termo da sequência de Fibonacci.
//
// A sequência de Fibonacci é uma sequência de números naturais em que cada número subsequente é a soma dos dois anteriores, começando de 0 e 1.
//
// A função utiliza recursividade para calcular os termos da sequência de Fibonacci.
//
// Exemplo:
//
//	fibonnacci(10) // 55
func fibonnacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonnacci(posicao-2) + fibonnacci(posicao-1)
}

// Imprime a sequência de Fibonacci com 11 termos.
//
// A sequência de Fibonacci é uma sequência de números naturais em que cada número subsequente é a soma dos dois anteriores, começando de 0 e 1.
//
// Essa função imprime os primeiros 11 termos da sequência de Fibonacci, de 0 a 55.
func main() {
	fmt.Println("Funções Recursivas")
	for i := range 500 {
		fmt.Println(fibonnacci(uint(i)))
	}
}
