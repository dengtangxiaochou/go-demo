package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	//fmt.Println(v)
	//获取原始值
	//v.Kind()//获取种类
	kind := v.Kind()
	switch kind {
	case reflect.Int64:
		fmt.Printf("int类型的原始值%v\n ", v.Int()+10)
	case reflect.Float32:
		fmt.Printf("Float32类型的原始值%v\n ", v.Float()+10.1)
	case reflect.Float64:
		fmt.Printf("Float64类型的原始值%v\n ", v.Float()+10.1)
	case reflect.String:
		fmt.Printf("string类型的原始值%v\n ", v.String())
	default:
		fmt.Printf("还没有判断")
	}
}
func main() {
	var a float32 =3.14
	var b int64 = 100
	var c string = "你好"
	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
}