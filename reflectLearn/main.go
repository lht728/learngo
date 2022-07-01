package main

import (
	"fmt"
	"reflect"
)

type myInt int64

//空接口作为函数参数 可以接收任意类型变量
func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%-20v kind:%v\n", t.Name(), t.Kind())
}

func main() {
	var a *float32
	var b myInt
	var c rune
	reflectType(a)	//Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
	reflectType(b)
	reflectType(c)

	type person struct {
		name string
		age  int
	}
	type book struct {
		title    string
		category string
	}

	var d = person{name: "月月", age: 18}
	var e = book{title: "另一个，同一个", category: "诗歌"}
	reflectType(d)
	reflectType(e)
}
