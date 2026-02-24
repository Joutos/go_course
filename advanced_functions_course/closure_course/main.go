package main

import "fmt"

func closure() func(int, int) int {
	return func(a, b int) int {
		return a - b
	}
}

func main() {
	fmt.Println("Aprendendo Closures")

	soma := func(a, b int) int {
		return a + b
	}

	fmt.Println(soma(10, 5))

	subtracao := closure()
	fmt.Println(subtracao(10, 5))
}
