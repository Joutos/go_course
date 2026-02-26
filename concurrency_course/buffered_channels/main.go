package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "Olá mundo"

	mensagem := <-channel
	fmt.Println(mensagem)
}
