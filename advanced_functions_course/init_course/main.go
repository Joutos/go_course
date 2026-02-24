package main

import "fmt"

var mensagem string = "Mensagem global"

func init() {
	fmt.Println("Executando função init")
	mensagem = "Mensagem alterada"
}
func main() {
	fmt.Println("Aprendendo Init Func")
	fmt.Println(mensagem)

}
