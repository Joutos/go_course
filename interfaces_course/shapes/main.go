package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type forma interface {
	area() float64
}

func main() {
	r := retangulo{10, 15}
	escrevarArea(r)
	circulo := circulo{20}
	escrevarArea(circulo)
}

func (r retangulo) area() float64 {
	return (r.altura * r.largura)
}

func (c circulo) area() float64 {
	return (math.Pi * (c.raio * c.raio))
}

func escrevarArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
	f.area()
}
