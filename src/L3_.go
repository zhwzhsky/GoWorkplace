package main

import (
	"fmt"
)

func main() {
	/*类型与变量*/
	// 声明多个变量
	var (
		a int    = 230
		b string = "大家好"
		c bool
		d []int
	)
	// 采用别名声明变量
	type (
		Bytesize int64
	)
	var e Bytesize

	// 省略关键字的声明和定义
	f := 899 //全局变量不可使用这种声明方式,:相当于var关键字
	//
	f = f + 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	aa, _, cc, dd := 11, 22, 33, 44
	fmt.Println(aa)
	fmt.Println(cc)
	fmt.Println(dd)

	/*类型转换*/
	// Go不存在隐式转换
	var a3 int = 65
	b3 := string(a3)
	fmt.Println(b3)

}
