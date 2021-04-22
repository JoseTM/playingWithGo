package main

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sidelength float64
}

type shape interface {
	getArea() float64
}

func main() {
	triangulo := triangle{height: 5, base: 8}
	cuadrado := square{sidelength: 2}

	printArea(triangulo)
	printArea(cuadrado)
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func printArea(s shape) {
	println(s.getArea())
}
