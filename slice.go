package main 

import (
	"fmt"
)

func main() {
	array := []byte{'a','b','c','d','e','f','g'}
	slice := array[2:4]
	fmt.Println(string(slice))
}
