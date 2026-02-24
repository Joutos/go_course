package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recuperado com sucesso", r)
	}
}
// aprovado - Verifica se a média das notas n1 e n2 é maior ou igual a 6.
// Se a média for maior ou igual a 6, retorna true, caso contrário, false.
// A função imprime também uma mensagem indicando que a média foi calculada e qual foi o resultado.
// Caso a média seja exatamente igual a 6, a função lança um panic.
func aprovado(n1, n2 float64) bool {
	fmt.Println("Verificando aprovaçao")
	defer recuperarExecucao()
	defer fmt.Println("Resultado calculado")

	avg := (n1 + n2) / 2

	if avg > 6 {
		return true
	} else if avg < 6 {
		return false
	}

	panic("A média é exatamente 6")
}

// main - Função principal do pacote, responsável por executar
// todas as funções e rotinas do pacote.
//
// No exemplo abaixo, a função aprovado é executada com um valor
// de média exatamente igual a 6, o que gera um panic.
// A função recuperarExecucao é chamada com defer, o que significa que
// ela será executada no final do programa.
// Essa função imprime uma mensagem indicando que a média foi calculada e
// qual foi o resultado.
func main() {
	fmt.Println("Aprendendo Panic e Recover")
	fmt.Println(aprovado(6, 6))
	fmt.Println("Pós")
}
