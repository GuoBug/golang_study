package main

import(
	"fmt"
)

func main() {
	a,b := 1,2
	argc(a,b)
	fmt.Println(a,b)

	fmt.Println("next :\n")

	f := func (c int)int{
		fmt.Println("匿名函数")
		return c * c
	}

	fmt.Println(f(4))
	fmt.Println(f(9))

	funcClosure := closure(10)
	fmt.Println(&funcClosure)
	fmt.Println(funcClosure(44))
}

func argc(a ...int) {
	fmt.Println(a)
}

/* 闭包 */

func closure(iIncome int) func(int) int{
	fmt.Println("%p",&iIncome)
	fmt.Println("Function closure")
	return func (inner int) int{
		fmt.Println("%p",&iIncome)
		return iIncome * inner
		
	}
}