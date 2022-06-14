package main

import (
	"fmt"
)

type Sayer interface {
	Say()
}

type Mover interface {
	Move()
}

// MakeHungry 饿肚子了...
func MakeHungry(s Sayer) {
	s.Say()
}

// 只要一个类型实现了接口中规定的所有方法，那么它就实现了这个接口。
// 类型Cat dog sheep都实现了接口Sayer
type Cat struct{}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog struct {
	Name string
}

func (d Dog) Say() {
	fmt.Printf("%s汪汪汪~\n", d.Name)
}

func (d Dog) Move() {
	fmt.Printf("%s会动~\n", d.Name)
}

type Sheep struct{}

func (s Sheep) Say() {
	fmt.Println("咩咩咩~")
}

func main() {
	// c := Cat{}
	// MakeHungry(c)
	// d := Dog{}
	// d.Say()
	// y := Sheep{}
	// y.Say()
	var d = Dog{Name: "旺财"}
	var s Sayer = d
	var m Mover = d
	s.Say()
	m.Move()
}
