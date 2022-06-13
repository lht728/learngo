package calc

//自定义包，由于实现代码复用

import "fmt"

//Add1 求x+y+1的值
func Add1(x, y int) int {
	return x + y + 1
}

//Add2 求x+y-1的值
func Add2(x, y int) int {
	return x + y - 1
}

//init函数在导入import时自动执行 没有参数也没有返回值
//常用于做一些初始化操作
func init() {
	fmt.Println("正在执行package calc的init函数...")
}
