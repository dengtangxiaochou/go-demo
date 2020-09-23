package main

import "fmt"

func MPrint(x interface{})  {
	if _, ok := x.(string); ok{
		fmt.Printf("a是一个string 值%v\n",x)
	}else if _, ok := x.(int); ok {
		fmt.Printf("a是一个int 值%v\n",x)
	}else if _, ok := x.(bool); ok {
		fmt.Printf("a是一个bool 值%v\n",x)
	}
}

func MPrint2(x interface{})  {
	switch x.(type) {
	case int:
		fmt.Println("int\n")
	case string:
		fmt.Println("string\n")
	case bool:
		fmt.Println("bool\n")
	default:
		fmt.Println("输入错误\n")
	}
}


//类型断言
func main()  {
	var a  interface{}
	a = "你好"
	_,ok := a.(string)
	if ok {
		fmt.Printf("a是一个string 值%v\n",a)
	}else {
		fmt.Println("断言失败")
	}
	MPrint(12)
	MPrint("哈哈")
	MPrint(true)
	MPrint2(12)
	MPrint2("哈哈")
	MPrint2(true)
}
