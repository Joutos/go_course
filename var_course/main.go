package main

import "fmt"

func main() {
	var text string = "Minha variável tipada!"
	fmt.Println(text)
	text2 := "Minha variável implícita!"
	fmt.Println(text2)

	var (
		variavel1 string
		variavel2 int
	)

	variavel1 = "Bom dia"
	variavel2 = 12651

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	const constante string = "Essa é minha constante"
	fmt.Println(constante, "Não pode mudar")
}
