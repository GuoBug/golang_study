package main

import(
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func main() {
	u := User{1, "OK", 12}
	Set(&u)
	fmt.Println(u)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("Can't set")
		return
	}else{
		v = v.Elem()
	}

	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("Changed")
	}
}