package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Declarando map
	var map1 map[string]int
	fmt.Println("Map vazio:", map1)

	// Populando maps
	map1 = make(map[string]int)
	map1["key1"] = 10
	map1["key2"] = 20
	fmt.Println("Populando maps:", map1)

	// Declarando e populando maps com inferência
	map2 := map[string]int{"key1": 10, "key2": 20}
	fmt.Println("Map Inferencia:", map2)

	fmt.Println("Valor key2:", map2["key2"])

	// Deletando maps
	delete(map2, "key2")
	fmt.Println("Map Deletado:", map2)

	usuario := map[string]map[string]string{
		"dev": {
			"nome":      "João Paulo",
			"sobrenome": "Leão",
		},
		"dev2": {
			"nome":      "Sofia Helena",
			"sobrenome": "Ibanez",
		},
	}
	fmt.Println(usuario)
}
