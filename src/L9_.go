package main

import (
	"fmt"
)

func main() {
	fmt.Println(A(1, 2, 3))

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	B(s1)
	fmt.Println(s1)
	//通过指针传递改变值
	c := 100
	C(&c)
	fmt.Println(c)

	//Go语言中，函数也是一种变量
	fd := D //定义fd为D类型
	fd()

	//匿名函数
	fnoname := func() {
		fmt.Println("匿名函数运行成功！")

	}
	fnoname()

	//闭包函数调用
	bibao := closure(10)
	fmt.Println(bibao(1))
	fmt.Println(bibao(2))

	W()

	//defer操作,逆序执行
	//defer:即使发生严重错误也能执行，用于资源清理、文件关闭、解锁、记录时间等
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	for i := 0; i < 3; i++ {
		defer func() {
			defer fmt.Println(i)
		}()
	}

}

func A(a ...int) (max int) {
	//func A(a ...int)  不定长变参，必须作为最后一个输入参数
	lena := len(a)
	max = a[0]
	for i := 1; i < lena; i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}

func B(s []int) {
	s[0] = 11
	s[2] = 12
}

func C(p *int) {
	*p = *p + 1
}

func D() {
	fmt.Println("你好！")
}

//闭包，返回一个匿名函数
func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// 引发panic，调用recover
func W() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover，回复继续执行")
		}
	}()
	panic("运行panic，终止执行")
}
