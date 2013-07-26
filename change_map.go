package main 

import (
	"fmt"
)

func main() {
	intMap := map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e",6:"f",7:"g"}
	var strMap map[string]int = make(map[string]int)

	/* 
	 * var intMap map[int]string
	 * intMap = make(map[int]string)
	 * var intMap map[int]string = make(map[int]string) 
	 */

	for  iCount,strTmp := range intMap{
		strMap[strTmp] = iCount
	}

	fmt.Println(intMap)
	fmt.Println(strMap)
}
