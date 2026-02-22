package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
}
func funcao2() {
	fmt.Println("Função 2")
}

// aprovado - Verifica se a média das notas n1 e n2 é maior ou igual a 6.
// Se a média for maior ou igual a 6, retorna true, caso contrário, false.
// A função imprime também uma mensagem indicando que a média foi calculada e qual foi o resultado.
func aprovado(n1, n2 float64) bool {
	defer fmt.Println("Media calculada. Resultado:")
	fmt.Println("Verificando aprovaçao")

	avg := (n1 + n2) / 2
	if avg >= 6 {
		return true
	}
	return false
}

// Main é a função principal do programa e é responsável por chamar as outras funções.
// Nesse exemplo, a função funcao1 é chamada com defer, o que significa que ela será executada no final do programa.
// Enquanto, a função funcao2 é chamada normalmente.
// O resultado da execução do programa é a impressão de "Função 2" e "Função 1".
func main() {
	fmt.Println("Estudando Defer")
	defer funcao1()
	funcao2()
	fmt.Println(aprovado(7, 8))
}
