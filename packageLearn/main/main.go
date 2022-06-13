package main

import (
	"fmt"
	ccc "learngo/packageLearn/calc" //import calc包时会自动执行calc的init函数
)

func main() {
	fmt.Println("开始执行main函数~")
	ans3 := ccc.Minus_func2(1, 2)
	ans1 := ccc.Add1(10, 20)
	ans2 := ccc.Add2(10, 20)
	fmt.Println(ans1)
	fmt.Println(ans2)
	fmt.Println(ans3)

}
