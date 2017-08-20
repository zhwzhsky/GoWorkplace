package main

import (
	"fmt"
)

type USB interface {
	//USB接口定义两个方法如下：
	Name() string // Name方法返回string类型的返回值
	Connecter
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	//定义一个结构，将该结构实现USB的方法
	name string
}

func (pc PhoneConnecter) Name() string {
	// 返回名称
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	var a USB                    //定义a为一个USB接口类型的变量
	a = PhoneConnecter{"xiaomi"} //我们认为PhoneConnecter实现了UBS接口，所以直接赋值
	a.Connect()
	Disconnect(a)

}

func Disconnect(usb interface{}) {
	// fmt.Println("Disconnected:", usb.Name())
	// 类型判断，判断跑是否为PhoneConnecter类型
	// if pc, ok := usb.(PhoneConnecter); ok {
	// 	fmt.Println("Disconnected:", pc.name)
	// 	return //直接返回
	// }
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown device.")
	}
}
