package main

import (
	"fmt"
)

type Sayer interface {
	Say()
}

// MakeHungry 饿肚子了...
func MakeHungry(s Sayer) {
	s.Say()
}

type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog struct{}

func (d Dog) Say() {
	fmt.Println("汪汪汪~")
}

type Sheep struct{}

func (s Sheep) Say() {
	fmt.Println("咩咩咩~")
}

func main() {
	c := Cat{}
	// c.Say() // 和下一行等价
	MakeHungry(c)
	d := Dog{}
	d.Say()
	y := Sheep{}
	y.Say()
}
