package main

import "C"
import "fmt"

type Animal interface {
	SetNnam(string)
}
type Animal2 interface {
	GetNname()string
}



type Dog struct {
	Name string
}

func (d *Dog) SetNnam(name string)  {
	d.Name = name
}

func (d Dog) GetNname() string {
	return d.Name
}

func main()  {

	//DOG实现Animal接口
	d := &Dog{Name:"小哥"}

	var d1 Animal = d //实现 Animal接口
	var d2 Animal2  = d //实现 Animal2接口

	d1.SetNnam("小花")
	fmt.Println(d2.GetNname())
}
