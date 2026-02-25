package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup = &sync.WaitGroup{}
	waitGroup.Add(2)
	go escrever("Hello World!", waitGroup)
	go escrever("Aprendendo waitGroup", waitGroup)
	waitGroup.Wait()
}

// escrever - Imprime um texto 6 vezes com um intervalo de 1 segundo entre cada impressão.
// Recebe um texto e um ponteiro para um objeto do tipo WaitGroup.
// Chama o método Done() ao finalizar a execução da função.
func escrever(texto string, waitGroup *sync.WaitGroup) {
	for range 6 {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
	waitGroup.Done()
}
