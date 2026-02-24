package main

import "fmt"

// inverterSinal inverte o sinal do valor apontado por n.
func inverterSinal(n *int) {
	*n *= -1
}

func main() {
	fmt.Println("Usando ponteiros em funções")
	n := 100
	inverterSinal(&n)
	fmt.Println(n)
}
