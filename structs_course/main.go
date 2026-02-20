package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint
	senha    string
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint
}

func main() {
	fmt.Println("Aprendendo structs")
	var u usuario
	u.nome = "João Paulo Leão"
	u.idade = 22
	u.senha = "1234"
	fmt.Println("Usuário 1:", u)

	usuario2 := usuario{
		"Sofia Helena Ibanez",
		24,
		"1234",
		endereco{"Avenida Nossa Senhora do Sabará", 1228},
	}

	fmt.Println("Usuário 2", usuario2)

	usuario3 := usuario{nome: "Lucas Pignagrande"}

	fmt.Println(usuario3)
}
