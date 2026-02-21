package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)
	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiro - Referencia de memória
	/*
		& - Endereço de memória
		* - Dado armazenado no endereço de memória
	*/
	var variavel3 int = 100
	var ponteiro *int

	ponteiro = &variavel3
	fmt.Println(variavel3, *ponteiro)

	variavel3++
	fmt.Println(variavel3, *ponteiro)

	*ponteiro++
	fmt.Println(variavel3, *ponteiro)
}
