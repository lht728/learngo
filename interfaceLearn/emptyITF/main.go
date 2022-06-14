package main

import "fmt"

// 定义一个不包含任何方法的空接口类型 Any
type Any interface{}

type Dog struct{}

// 空接口可以作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	var x Any
	x = "你好"
	fmt.Printf("type:%T value:%v\n", x, x)
	x = 100 // int型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = true // 布尔型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = Dog{} // 结构体类型
	fmt.Printf("type:%T value:%v\n", x, x)

	fmt.Printf("——————————————————\n")
	// 空接口可以作为函数参数
	var y Any
	show(y)

	fmt.Printf("——————————————————\n")
	// 空接口作为map值 可以实现保存任意类型值的map
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "鱼蛋"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

}
