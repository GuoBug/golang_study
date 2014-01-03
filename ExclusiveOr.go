/*
	异域运算的运用
	说在一个整型数组里，所有的数都重复了两次，仅有两个数是不重复的，如何在时间 O(n) 和空间 O(1) 内找出这两个数？

	算法出自 : http://www.v2ex.com/t/94742
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{23, 23, 19, 19, 12, 88, 88, 3, 3, 5, 56, 56}

	//设两个不重复的数为a1和a2，x = a1 ^ a2，bits为a1和a2某个不一致的位
	var a1, a2, x, bits int

	//将所有的数字异或，得到的结果就为x，因为重复的数经过异或后都为0
	for _, v := range nums {
		x = x ^ v
	}

	//找出a1和a2某个不一致的位，换句话说，就是找出x为1的位（当然，x为1的位有很多，我们这找的是x从右往左第一个为1的位）
	bits = 1
	for i := 31; i >= 0; i++ {
		if x&bits != 0 {
			break
		}
		bits = bits << 1
	}

	//舍去所有bits位为0的位，将剩下的数字全部异或，这样就能得出两个不重复的数字其中的一个
	for _, v := range nums {
		fmt.Printf("V  [%v],[%b]\n", v, v)
		if v&bits != 0 {
			fmt.Printf("a1: %b,%v  v:%b,%v\n", a1, a1, v, v)
			a1 = a1 ^ v
		}
	}

	//根据x和a1可以很容易求出a2
	a2 = x ^ a1

	fmt.Printf("Result : %v,%v\n", a1, a2)

}
