/* The Fibonacci sequence is defined as: 
 * fib(0) = 0, fib(1) = 1, 
 * fib(n) = fib(n-1) + fib(n-2). 
 * Write a recursive function which can find fib(n).
 */

package main

import "fmt"

func main() {
    testNum1 := 5
    testNum2 := 3
    testNum3 := 4
    
    fmt.Println(fib(testNum1))
    fmt.Println(fib(testNum2))
    fmt.Println(fib(testNum3))
}

func fib(input int) int{
    if (input < 2 ){
        return input
    }else{
        return fib(input-1)+fib(input-2)
    }
}
