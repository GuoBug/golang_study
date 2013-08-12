/*
三角和圆的面积和
*/

package main

import (
  "fmt";
	"math"
)

type niltype interface{}

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

	pointA := point{0,0}
	pointB := point{3,4}
	myR := rectangle{pointA,pointB}

	fmt.Println("求和为：",countSum(myR,c))
}

func (c *cricle) area() float64{
	return math.Pi*c.r*c.r
}

func (r *rectangle) rArea() float64{
	a := r.a.x - r.b.x
	b := r.a.y - r.b.y
	return  a*b
}

func countSum(sharpInput ...niltype) float64{

	fmt.Println("这个里面有函数")
	var sumNum float64
	sumNum = 0
	for _,sharp := range sharpInput{
		switch sharp := sharp.(type){
		case rectangle:
			sumNum += sharp.rArea()
			break
		case cricle:
			sumNum += sharp.area()
			break
		default:
			fmt.Println("Unknown",sharp)
			break
		}
	}
	return sumNum
}
