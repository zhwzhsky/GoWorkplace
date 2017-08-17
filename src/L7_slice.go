package main

import (
	"fmt"
)

func main() {
	//切片：slice，引用型
	//档超过cap能力后，slice不再指向原数组，而是复制一个新数组并指向
	var s1 []int = make([]int, 3) //注意区别：数组是[...]int
	fmt.Println("len(s1)=", len(s1), "   cap(s1)=", cap(s1))
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	s1 = a[1:6]
	s1 = append(s1, 100, 101, 102)
	fmt.Println("len(s1)=", len(s1), "   cap(s1)=", cap(s1))
	s2 := a[:2]
	s3 := s1[:2]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println()
}
