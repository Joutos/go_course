package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// if init
	if numero2 := 14; numero2 > 15 {
		fmt.Println("Maior que 15")
	} else if numero2 == 14 {
		fmt.Println("Catorze")
	} else {
		fmt.Println("Menor doque 14")
	}
}
