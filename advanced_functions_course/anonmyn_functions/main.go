package main

import "fmt"

// main - Função principal do pacote, responsável por executar
// todas as funções e rotinas do pacote.
//
// No exemplo abaixo, uma função anônima é executada.
// Essa função não tem nome e não pode ser chamada novamente.
// Ela é criada e executada em tempo de execução.
func main() {
	func() {
		fmt.Println("Anonima")
	}()
}
