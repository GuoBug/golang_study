/* a的地址给 Increate 函数操作指针
 * 之后函数退出，指针指向内容被修改
 * 也就是修改成功
 */

package main

import(
	"fmt"
)

type small int

func main() {
	var a small
	a = 0
	(&a).Increate()
	fmt.Println(a)
}

func (a *small) Increate() {
	for i:= 0;i < 100; i ++ {
		*a++
		fmt.Println(*a)
	}
}