package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go escrever("Hello world", channel, wg)
	go escrever("Learning channels", channel, wg)

	go func() {
		wg.Wait()
		close(channel)
	}()
	// for {
	// 	mensagem, open := <-channel
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }
	for mensagem := range channel {
		fmt.Println(mensagem)
	}
}

func escrever(texto string, channel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 5 {
		channel <- texto
		time.Sleep(time.Second)
	}
	close(channel)
}
