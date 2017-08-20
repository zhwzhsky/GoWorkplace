package main

import (
	"fmt"
	// "strconv"
)

// const A int = 1
// const B = 'A'
// const (
// 	C = A
// 	D
// )
// const E = 88

// const (
// 	MAX_COUNT = 'A'
// 	B
// 	C  = iota
// 	D
// )

func main() {
	/*L3作业：string转换*/
	// var a int = 65
	// b := string(a)
	// fmt.Println(b)

	// c := strconv.Itoa(a)
	// fmt.Println(c)

	/*  常量  */

	// fmt.Println(A)
	// fmt.Println(B)
	// fmt.Println(C)
	// fmt.Println(D)
	// fmt.Println(E)

	/* 枚举iota*/
	/* iota从0开始，遇到const重置为0*/
	// fmt.Println(MAX_COUNT)
	// fmt.Println(B)
	// fmt.Println(C)
	// fmt.Println(D)
	fmt.Println(^1)
	fmt.Println(!true)
	fmt.Println(1 << 10 << 10 >> 10)
	/*
		6:  0110
		11: 1011
		--------------
		&   0010 = 2
		|   1111 = 15
		^   1101 = 13
		&^  0100 = 4 如果第二个数的这一位是1，则将第一个数的这位改为0；

	*/
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)

	a := 0
	if (a > 0) && (10/a > 1) { //&&前的公式不成立的时候，后面一个表达式就不会执行了
		fmt.Println("ok")
	}
}
