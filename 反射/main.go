package main

import (
	"fmt"
	"reflect"
)

type myInt int
type Person struct {
	Naem string
	Age int
}

//通过反射获取空接口的类型
func reflectFn(x interface{})  {
	v := reflect.TypeOf(x)
	//v.Name()//获取类型名称
	//v.Kind() //获取种类
	//fmt.Println(v)
	fmt.Printf("类型：%v 类型名称：%v 类型种类：%v\n",v,v.Name(),v.Kind())
}

func main()  {
	a := 10
	b := 23.4
	c := true
	d := "你好"
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

	var e myInt = 34
	var f  =Person{
		Naem: "张山",
		Age:  23,
	}
	reflectFn(e)
	reflectFn(f)

	var h = 25
	reflectFn(&h)

	var i = [3]int{1,2,3}
	reflectFn(i)

	var j = []int{11,22,33}
	reflectFn(j)
}
