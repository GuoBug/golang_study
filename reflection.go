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

func (u User) Hello() {
	fmt.Println("Hello ,world.")
}

func main() {
	me := User{1,"su",24}
	Info(me)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)

	fmt.Println("Type:",t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("field:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%s : %v = %v\n",f.Name, f.Type,val)
	}
}