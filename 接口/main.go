package main

import "fmt"

type A interface {} //空接口，没有任何约束，任意的类型都可以实现空接口

func show(a interface{}) {
	fmt.Printf("值：%v 类型%T\n",a,a)
}


func main()  {
	var a  A
	var str  = "你好"
	a = str
	fmt.Printf("值：%v 类型%T\n",a,a)

	var num  = 20
	a = num//表示让int类型实现A这个接口
	fmt.Printf("值：%v 类型%T\n",a,a)

	var ns  = false
	a = ns
	fmt.Printf("值：%v 类型%T\n",a,a)

	//空接口可以直接当类型来使用。可以表示任意类型
	var b  interface{}
	b = 20
	fmt.Printf("值：%v 类型%T\n",b,b)
	b = "哈哈"
	fmt.Printf("值：%v 类型%T\n",b,b)
	b = false
	fmt.Printf("值：%v 类型%T\n",b,b)

	show(20)
	show("你好")
	slice := []int{1,2,3}
	show(slice)

	//map
	var m1  = make(map[string]interface{})
	m1["name"] = "张三"
	m1["age"] = 20
	m1["married"] = false
	fmt.Println(m1)

	//切片
	var s1  = []interface{}{1,2,3,"你好",false}
	fmt.Println(s1)


}
