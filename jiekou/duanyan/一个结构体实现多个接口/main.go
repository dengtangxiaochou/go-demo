package main

import "C"
import "fmt"

type Animal interface {
	SetNnam(string)
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

type Gat struct {
	Name string
}

func (g *Gat) SetNnam(name string)  {
	g.Name = name
}

func (g Gat) GetNname()string  {
	return g.Name
}

func main()  {

	//DOG实现Animal接口
	d := &Dog{Name:"小哥"}

	var d1  Animal = d
	fmt.Println(d1.GetNname())
	d1.SetNnam("小弟")
	fmt.Println(d1.GetNname())

	//GAT实现Animal接口
	c := &Gat{Name:"小花"}
	var c1  Animal = c
	fmt.Println(c1.GetNname())

}
