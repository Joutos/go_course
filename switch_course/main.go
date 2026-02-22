package main

import "fmt"

func main() {
	fmt.Println("Switch")

	var Dia int = 2

	switch Dia {
	case 0:
		fmt.Println("Domingo")
	case 1:
		fmt.Println("Segunda-feira")
	case 2:
		fmt.Println("Terça-feira")
	case 3:
		fmt.Println("Quarta-feira")
	case 4:
		fmt.Println("Quinta-feira")
	case 5:
		fmt.Println("Sexta-feira")
	case 6:
		fmt.Println("Sábado")
	default:
		fmt.Println("Dia inválido")
	}
}
