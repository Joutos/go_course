package main

import "fmt"

func main() {
	fmt.Println(somar(100, 83))

	var f = func(n1, n2 int) int {
		return n1 + n2
	}

	fmt.Println(f(1, 256))

	fmt.Println(soma_subtrai(10, 5))

	soma, _ := soma_subtrai(10, 5)
	fmt.Println(soma)
}

func somar(n1 int, n2 int) int {
	return n1 + n2
}

func soma_subtrai(n1, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}
