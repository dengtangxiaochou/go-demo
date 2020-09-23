package main

import "fmt"

type A interface {
	SetName(string)
}

type B interface {
	Getname()string
}

type Animaler interface {//接口的嵌套
	A
	B
}

type Dog struct {
	Name string
}

func (d Dog) SetName(name string)  {
	d.Name = name
}

func (d Dog) Getname()string  {
	return d.Name
}

func main()  {
	c := &Dog{Name:"小米"}
	var d1 Animaler = c //实现 Animal接口

	d1.SetName("小号")
	fmt.Println(d1.Getname())
}

