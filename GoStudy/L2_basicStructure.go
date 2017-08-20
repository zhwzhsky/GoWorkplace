// 当前程序的包名
package main

// 导入其他的包
// import "fmt"
import sky "fmt" //定义别名

// 常量的定义
const pi = 3.14
const (
	A = 1
	B = 2
)

//全局变量的定义与赋值
var name = "gopher"

//一般类型声明
type myType int

//结构声明
type gopher struct{}

// 接口声明
type golang interface{}

// 函数
func main() {
	sky.Println("第二堂课")
	//fmt.Println("Hello World! 你好，世界！")
}
