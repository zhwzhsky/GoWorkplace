package main

import (
	"fmt"
)

func main() {
	//数组在Go中是值类型，切片是引用类型
	var c [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	b := [2]int{1, 2}

	a := [...]int{1: 2, 19: 19}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	var p *[20]int
	p2 := &a[19]
	p = &a
	fmt.Println(*p)
	fmt.Println(*p2)

	//指针数组
	// var p3 [3]*int
	p3 := [...]*int{&b[0], &b[1]}
	fmt.Println(p3)
	for i := 0; i < len(p3); i++ {
		fmt.Println(*p3[i])
	}
}
