package main

import (
	"fmt"
	"sort"
)

func main() {
	// Map比现行查找快很多，但是比索（如数组、slice）引慢100倍
	var m map[int]string
	m = map[int]string{}
	var m2 map[int]string = make(map[int]string)
	// m3 := make(map[int]string)

	m[204782] = "赵伟志"
	m[204781] = "李"
	fmt.Println(m)
	fmt.Println(m2)
	a := m[204782]
	fmt.Println(a)

	m2 = map[int]string{4781: "li", 4782: "zhao", 3324: "chun"}
	s := make([]int, len(m2))
	i := 0
	for k, _ := range m2 {
		s[i] = k
		i++
	}
	fmt.Println(s)
	sort.Ints(s) //排序
	fmt.Println(s)
	m3 := map[string]int{}
	for k, v := range m2 {
		m3[v] = k
	}
	fmt.Println(m3)
}
