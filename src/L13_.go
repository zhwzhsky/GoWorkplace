package main

import (
	"fmt"
	"reflect"
)

type User struct {
	//定义一个结构
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	//定义User结构类型的一个方法Hello（）
	fmt.Println("HelloWorld!")
}

/*反射：使用TypeOf和ValueOf函数从接口获取目标对象信息*/
func main() {
	u := User{1, "sky", 28} //定义并初始化一个结构
	Info(u)
}
func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())
	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v=%v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)

	}
}
