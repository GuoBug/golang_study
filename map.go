package main

import (
	"fmt"

//	"sort"
)

func main() {
	m := map[int]string{3: "Love", 1: "Guo", 2: "Qiang", 4: "TT"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	//	sort.Ints(s)
	fmt.Println(s)
	fmt.Println(m)
}
