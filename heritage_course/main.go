package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint
	altura    uint
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

/*
Go não possui herança de verdade
*/
func main() {
	fmt.Println("'Herança' em Go")

	p1 := pessoa{"João Paulo", "Leão", 22, 180}
	fmt.Println("Pessoa:", p1)

	e1 := estudante{p1, "Sistemas de Informação", "Una"}
	fmt.Println("Estudante:", e1)
}
