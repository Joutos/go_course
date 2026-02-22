package main

import (
	"fmt"
)

func main() {

	// i := 0

	// for i < 10 {
	// 	i++
	// 	time.Sleep(time.Second)
	// 	fmt.Println("I:", i)
	// }

	// for j := 0; j < 10; j++ {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("J:", j)
	// }

	nomes := [3]string{"Sofia", "João", "Alfredo"}
	for _, nome := range nomes {
		fmt.Println("Nome:", nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println("Indice:", indice, "Letra:", string(letra))
	}

	usuario := map[int]map[string]string{}

	usuario[0] = map[string]string{
		"nome":      "João Paulo",
		"sobrenome": "Leão",
	}

	usuario[1] = map[string]string{
		"nome":      "Sofia Helena",
		"sobrenome": "Ibanez",
	}

	usuario[len(usuario)] = map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Pignagrande",
	}

	for _, v := range usuario {
		fmt.Println(v["nome"], "\n", v["sobrenome"])
	}

	type pessoa struct {
		nome      string
		sobrenome string
	}

	p1 := pessoa{"João Paulo", "Leão"}
	p2 := pessoa{"Sofia Helena", "Ibanez"}
	p3 := pessoa{"Lucas", "Pignagrande"}

	usuarios := []pessoa{p1, p2, p3}
	for _, v := range usuarios {
		fmt.Println(v.nome, "\n", v.sobrenome)
	}

	for {
		fmt.Println("Executando")
	}
}
