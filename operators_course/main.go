package main

import "fmt"

func main() {
	// Aritiméticos

	soma := 2 + 2
	subtracao := 10 - 5
	divisao := 100 / 20
	resto := 10 % 2
	multiplicacao := 5 * 5

	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)

	// Atribuição (=|:=)

	var variavel string = "Texto!"
	variavel2 := "Texto 2!"

	fmt.Println(variavel, variavel2)

	// Relacionais (>|<|==|!=)
	fmt.Println(2 < 1, 2 > 1, 1 <= 1, 4 >= 1)

	// Lógicos (!|&&|"||")
	fmt.Println(true && false, false || true, !true)

	// Unários (++|--|+=|-=|*=|/=|%=)
	num := 10
	num++
	fmt.Println(num)

	
}
