package calc

import "fmt"

func minus_func1(x, y int) int {
	return x - y - 1
}

func Minus_func2(x, y int) int {
	return x - y - 2
}

//init函数在导入import时自动执行 没有参数也没有返回值
func init() {
	fmt.Println("正在执行package calc的init函数2...")
}
