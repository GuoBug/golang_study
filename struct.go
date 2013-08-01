package main

import (
	"fmt"
)

type person struct{
	name string
	age int
}

type superMan struct{
	person
	power string
}

func main() {
	me := &person{
		name:"GuoBug",
		age:0,
	}

	su := &superMan{
		person:person{
			name:"GuoBug!",
			age:100,
			},
		power:"eat",
		}

	fmt.Println(me)
	fmt.Println(su)

	A(me)

	fmt.Println(me)
}

func A(persons *person){
	persons.name = "Mr.A"
	persons.age = 99
}