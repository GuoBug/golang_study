package main 

import (
	"fmt"
)

func main() {
	var a [2]int
	var b [1]int
	c := [6]int{1,2,3,4,5,6}
	d := [20]int{19:1}
	e := [...]int{2,1,3,4}

	f := new([10]int)  //指向数组的指针

	var p *[2]int = &a

	c[2] = 10
	f[2] = 10

	fmt.Println("a:",a)
	fmt.Println("b:",b)
	fmt.Println("c"+c)
	fmt.Println("d"+d)
	fmt.Println("e"+e)
	fmt.Println("f"+f)

	fmt.Println(p)
}