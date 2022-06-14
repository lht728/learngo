package main

import "fmt"

type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

type Car struct {
	Brand string
}

func (d *Dog) Move() {
	fmt.Printf("Dog %s is running!\n", d.Name)
}

func (c *Car) Move() {
	fmt.Printf("Car %s is running, too!\n", c.Brand)
}

func main() {
	var m Mover

	fmt.Printf("type:%T value:%v\n", m, m)
	m = &Dog{Name: "旺旺"}
	fmt.Printf("type:%T value:%v\n", m, m)
	m.Move()

	m = &Car{Brand: "宝马"}
	fmt.Printf("type:%T value:%v\n", m, m)
	m.Move()

	var n Mover = &Dog{Name: "旺财"}
	v, ok := n.(*Dog)
	if ok {
		fmt.Println("类型断言成功")
		v.Name = "富贵"
	} else {
		fmt.Println("类型断言失败")
	}

}
