package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{})  {
	//var mun = 10 + x//10 + x (mismatched types int and interface {})
	//类型断言实现
	b ,_ := x.(int)
	var mun  = 10 + b
	fmt.Println(mun)

	//反射实现获取变量的原始值

	v := reflect.ValueOf(x)
	//fmt.Println(v)
	//var n = v + 13 //v + 13 (mismatched types reflect.Value and int)
	var m = v.Int() +12
	fmt.Println(m)
}

func main()  {
	var a = 13
	reflectValue(a)
}
