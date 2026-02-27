package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// main - Funçao principal do pacote, responsavel por executar
// todas as funções e rotinas do pacote.
//
// A função main chama a fun o multiplex, que multiplexa dois
func main() {
	channel := multiplex(escrever("Hello World"), escrever("Learning Multiplex"))

	for range 45 {
		fmt.Println(<-channel)
	}
}

// multiplex - multiplexa dois canais de strings em um único canal de saída.
// Recebe dois canais de strings e retorna um canal de saída que
// multiplexa as mensagens recebidas dos canais de entrada.
// A função multiplex é útil para quando você precisa lidar com
// dois ou mais fluxos de dados e precisa lidar com eles como se fossem
// um único fluxo de dados.
func multiplex(chanell1, channel2 <-chan string) <-chan string {
	outputChannel := make(chan string)
	go func() {
		for {
			select {
			case message := <-chanell1:
				outputChannel <- message
			case message := <-channel2:
				outputChannel <- message
			}
		}
	}()
	return outputChannel
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.IntN(2000)))
		}
	}()
	return canal
}
