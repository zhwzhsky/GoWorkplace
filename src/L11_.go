package main

// 变量大小写是区分私有或公有，所谓的私有是指整个包内的私有，包内的函数、方法都可以访问
import (
	"fmt"
)

type myInt int
type TZ int

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)

	var tz TZ
	tz.Print()
	(*TZ).Print(&tz) //两种实现方法

	var test myInt
	test = 500
	test.Increase()
	fmt.Println(test)

}

/*方法*/
//方法和类型绑定
func (a *A) Print() {
	//类型A的方法，a是变量；
	a.Name = "AA"
	fmt.Println("A")
}

func (a B) Print() {
	//类型A的方法，a是变量；
	a.Name = "BB"
	fmt.Println("B")
}

func (a *myInt) Increase() {
	*a = *a + 100
}
func (a *TZ) Print() {
	//TZ类的方法
	fmt.Println("TZ")
}
