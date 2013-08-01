package main

import(
	"fmt"
)

func main() {
	for i := 0 ; i < 5 ;i++{
		fmt.Println(i)
		defer fmt.Println(i)
	}

	// for i := 0; i < 5; i++ {
	// 	defer func(){
	// 		fmt.Println(i)
	// 	}()
	// }
	fmt.Println("试试panic")
	testRun()
	test()
}

func test() {
	defer func () {
		if err := recover(); err!=nil{
			fmt.Println("recover .....")
		}
	}()
	panic ("This is panic")
}

func testRun() {
	defer recover()
	fmt.Println("panic in Run")
}