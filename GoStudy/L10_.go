package main

import (
	"fmt"
)

type human struct {
	Sex int
}
type student struct {
	human   //嵌入human结构
	Name    string
	Age     int
	Contact struct {
		phone, city string
	}
}

func main() {
	xiaoming := student{}
	xiaoming.Name = "小明"
	xiaoming.Age = 20
	xiaoming.Contact.city = "shanghai"
	xiaoming.Contact.phone = "15121061996"
	xiaoming.Sex = 1
	hong := student{
		Age:   19,
		Name:  "小红",
		human: human{Sex: 1},
	}
	fmt.Println(xiaoming)
	fmt.Println(hong)
	A(xiaoming)
	fmt.Println(xiaoming)

	B(&xiaoming)
	fmt.Println(xiaoming)

	//yang就变成了指针
	yang := &student{
		Name: "Yang",
		Age:  27,
	}
	fmt.Println(yang, *yang)
	B(yang)
	fmt.Println(yang)

	/*匿名结构*/
	sky := &struct {
		Name string
		Age  int
	}{
		Name: "sky",
		Age:  28,
	}
	fmt.Println(sky)
}

//采用值传递，只是值得拷贝
func A(stu student) {
	stu.Age = 25
	fmt.Println("A", stu)
}

//采用指针，直接修改原有的值
func B(pStu *student) {
	pStu.Age = 25
	fmt.Println("B", *pStu)

}
