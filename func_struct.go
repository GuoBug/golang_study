package main

import (
  "fmt";
	"math"
)

type cricle struct {
	r float64
}

type point struct{
	x, y float64
}

type rectangle struct {
	a point
	b point
}

func main() {
	c := cricle{2}
	c.area()

	pointA := point{0,0}
	pointB := point{3,4}
	myR := rectangle{pointA,pointB}

	fmt.Println(myR)
	myR.rArea()
}

func (c *cricle) area() {
	fmt.Println(math.Pi*c.r*c.r)
}

func (r *rectangle) rArea() {
	a := r.a.x - r.b.x
	b := r.a.y - r.b.y
	fmt.Println("area:", a*b)
}
