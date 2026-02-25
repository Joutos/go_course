package main

import "fmt"

func generica(i interface{}) {
	fmt.Println(i)
}

func main() {
	generica("Texto")
	generica(1298)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:       "String",
		"Texto": true,
	}
	generica(mapa)
}
