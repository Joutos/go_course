package main

import (
	"errors"
	"fmt"
)

func main() {
	var float float32 = 3.14
	fmt.Println("Esse é meu float:", float)

	var integer int32 = -17
	fmt.Println("Esse é meu inteiro:", integer)

	// uint não pode ser negativo, int pode
	var unsigned_integer uint = 1000
	fmt.Println("Esse valor não pode ser negativo:", unsigned_integer)

	var texto string = "Esse é meu texto numa variável"
	fmt.Println("Esse é um texto simples", texto)

	char := 'A'

	fmt.Println("Esse é o valor de A em ASCII:", char)

	booleano := true

	fmt.Println("Esse é um booleano:", booleano)

	var erro error
	fmt.Println("O valor nulo de 'error' é:", erro)

	var erro_preenchido error = errors.New("Erro Interno para Teste")
	fmt.Println("Esse é o erro preenchido:", erro_preenchido)
}
