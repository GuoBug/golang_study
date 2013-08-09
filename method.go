package main

import(
	"fmt"
)

type intTest int8

type A struct{
	Name string
}

type B struct{
	Name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)

	var inta intTest
	inta.Print()
}

func (a *A) Print() {
	fmt.Println("A")
	a.Name = "AA"
}

func (b B) Print() {
	fmt.Println("B")
	b.Name = "BB"
}

func (a *intTest)Print() {
	fmt.Println("int")
}