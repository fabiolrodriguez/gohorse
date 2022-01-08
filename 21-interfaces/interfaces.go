package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreveArea(f forma) {
	fmt.Printf("A area de forma é %0.2f\n", f.area())
}

func main() {
	r := retangulo{10, 15}
	escreveArea(r)

	c := circulo{10}
	escreveArea(c)
}
