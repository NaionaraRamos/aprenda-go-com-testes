package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

type Circulo struct {
	Raio float64
}

//método
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

//função
func PerimetroRetangulo(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

// func Area(retangulo Retangulo) float64 {
// 	return retangulo.Largura * retangulo.Altura
// }

// func AreaCirculo(circulo Circulo) float64 {
// 	return circulo.Raio * circulo.Raio * 3
// }

func main() {

	circulo := Circulo{4}
	circ := circulo.Area()
	fmt.Println(circ)

	triangulo := Triangulo{3, 4}
	tri := triangulo.Area()
	fmt.Println(tri)

	retangulo := Retangulo{4, 5}
	ret := retangulo.Area()
	fmt.Println(ret)

}
