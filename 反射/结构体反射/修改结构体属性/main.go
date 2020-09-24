package main

import (
	"fmt"
	"reflect"
)
//student结构体
type student struct {
	Name string`json:"name"`
	Age int`json:"age"`
	Score int`json:"int"`
}

//反射修改结构体属性
func reflectChangeStruct(s interface{})  {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Ptr {
		fmt.Println("传入的不是结构体指针类型")
		return
	}else if  t.Elem().Kind() != reflect.Struct  {
		fmt.Println("传入的不是结构体指针类型")
		return
	}
	//修改结构体指针的值
	name := v.Elem().FieldByName("Name")
	name.SetString("小黄")

	age := v.Elem().FieldByName("Age")
	age.SetInt(19)
}

func main()  {
	stu1 := student{
		Name:  "小王",
		Age:   20,
		Score: 99,
	}

	reflectChangeStruct(&stu1)
}
