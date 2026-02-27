package main

import (
	"fmt"
	"time"
)

func main() {

	channel, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Canal 2"
		}
	}()

	for {
		select {
		case mensagem1 := <-channel:
			fmt.Println(mensagem1)
		case mensagem2 := <-channel2:
			fmt.Println(mensagem2)
		}
	}
}
