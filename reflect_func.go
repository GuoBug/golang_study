package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello(name string) {
	fmt.Println("Hello",name)
}

func main() {
	u := User{1,"Test",12}
	v := reflect.ValueOf(u)

	mv := v.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("asdf")}
	mv.Call(args)
}