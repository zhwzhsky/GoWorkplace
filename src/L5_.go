package main

import (
	"fmt"
)

func main() {
	a := 1
	var p *int = &a
	fmt.Println(*p)
	c := 55
	if b := 10; b > a {
		//条件表达式后面没有括号，大括号必须与if同行
		c := b - a
		fmt.Println(c)
	}
	fmt.Println(c)
	// Go里只有for循环，功能强大

	sum := 0
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

	// switch不需要写break，要继续执行需要加fallthrough
	sw := 1
	switch sw {
	case 0:
		fmt.Println("sw=0")
	case 1:
		fmt.Println("sw=1")
	default:
		fmt.Println("None")
	}

	switch {
	case sw >= 0:
		fmt.Println("sw=0")
		fallthrough
	case sw >= 1:
		fmt.Println("sw=1")
		fallthrough
	default:
		fmt.Println("None")
	}

	// 跳转语句，goto,break,continue
	// 可以跳转到标签处,break跳出后不再执行该循环，goto则调到指定位置开始执行
	// continue继续执行下次循环,配合标签需要跳至有限循环,break跳出后该循环不执行了
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("OK")

LABEL2:
	for i := 0; i < 3; i++ {
		for {
			fmt.Println(i)
			// continue LABEL2
			break LABEL2
		}
	}
	fmt.Println("OKOKOK")
}
