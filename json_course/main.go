package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type dog struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Breed string `json:"breed"`
}

func main() {
	d := dog{
		Name:  "Fido",
		Age:   3,
		Breed: "Golden Retriever",
	}
	fmt.Println(d)
	JSON, erro := json.Marshal(d)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(string(JSON))

	d2 := map[string]string{
		"name":  "Alfredo",
		"age":   "4",
		"breed": "Pequines",
	}

	JSON2, erro := json.Marshal(d2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(string(JSON2))

	json_retorno, erro := os.ReadFile("./example.json")

	if erro != nil {
		log.Fatal(erro)
	}
	var d3 dog
	json.Unmarshal(json_retorno, &d3)
	fmt.Println(d3)
}
