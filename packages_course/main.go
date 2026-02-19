package main

import (
	"fmt"
	"modulo/utils"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Estou estudando pacotes em Go!!")
	utils.Escrever()
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	if erro != nil {
		fmt.Println("Email inválido")
	} else {
		fmt.Println("Email válido")
	}
}
