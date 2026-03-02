package main

import (
	"fmt"
	"introducao-de-testes/adresses"
)

func main() {
	adressType := adresses.AdressType("Street A")
	fmt.Println(adressType)
}
