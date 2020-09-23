package main

import (
	"fmt"
	"reflect"
)

//错误方法
//func reflectValue(x interface{})  {
//	v := reflect.ValueOf(x)
//	if v.Kind() == reflect.Int64 {
//		v.SetInt(120)
//	}
//}

func reflectValue2(x interface{})  {
	//*x = 120 //invalid indirect of x (type interface {})
	//v ,_ = x.(*int)
	//*v =120
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind())//ptr
	fmt.Println(v.Elem().Kind())//int64
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(123)
	}else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("嘿嘿")
	}

}

func main()  {
	var a int64 = 100
	//reflectValue(a)
	reflectValue2(&a)
	fmt.Println(a)

	var b string = "哈哈"
	reflectValue2(&b)
	fmt.Println(b)
}
