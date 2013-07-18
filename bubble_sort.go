package main

import (
	"fmt"
)

func main() {
	/* 冒泡排序
	 * 循环运用
	 */
	array := [10]int{14, 24, 35, 46, 2, 61, 73, 25, 725, 43}

	fmt.Print("数组为")
	fmt.Println(array)

	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			if array[i] < array[j] {
				array[i] = array[i] + array[j]
				array[j] = array[i] - array[j]
				array[i] = array[i] - array[j]
			}
		}
	}
	fmt.Println(array)
}
