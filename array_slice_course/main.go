package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array")

	// Declarando array
	var array1 [5]int
	fmt.Println(array1)

	// Populando arrays
	array1[0] = 1
	array1[1] = 2
	array1[4] = 6
	fmt.Println(array1)

	// Declarando e populando arrays com inferência
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	//Array ... - Tem posição definida na inferencia
	array_ilimitado := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array_ilimitado)

	// Slice - Array dinâmico, apontam para array, ponteiro
	slice := []int{}
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)

	slice2 := array_ilimitado[5:]
	fmt.Println(slice2)

	// Arrays Internos
	fmt.Println("Arrays internos")
	slice3 := make([]int, 10, 15)
	slice3 = append(slice3, 6)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade
}
