package main

import (
	"fmt"
	"time"
)

// main - Função principal do pacote, responsável por executar
// todas as funções e rotinas do pacote.
//
// A função main é a primeira a ser executada em um programa Go,
// e é a responsável por chamar todas as funções e rotinas do pacote.
// No exemplo abaixo, a função main executa a goroutine (go escrever("Hello World"))
// e uma função normal (escrever("Aprendendo concorrência")).
func main() {
	go escrever("Hello World")
	escrever("Aprendendo concorrência")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
