package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Salvando dados do usuário: %s no banco\n", u.nome)
}

func (u usuario) maiorDeIdade() {
	if u.idade >= 18 {
		fmt.Println("Maior de idade")
	} else {
		fmt.Println("Menor de idade")
	}
}

func (u *usuario) aniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"João Paulo", 22}
	usuario1.salvar()
	usuario1.maiorDeIdade()
	usuario1.aniversario()
	fmt.Println("Nova idade:", usuario1.idade)
}
