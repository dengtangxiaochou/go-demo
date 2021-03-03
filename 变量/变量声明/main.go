package main

import "fmt"



func main()  {
	var name  string
	var age  int
	//声明变量必须使用
	fmt.Print(name)
	fmt.Print(age)

	var (
		a string
		b int
		c bool
		d string
	)
	fmt.Println(a,b,c,d)

	a = "沙河"
	b = 100
	c = true
	d = "100"

	fmt.Println(a,b,c,d)
	//声明变量赋值
	var x string = "老男孩"
	fmt.Println(x)
	fmt.Printf("%s嘿嘿%d\n",x,b)
	//类型推导(编译器根据变量初始值的类型，指定给变量)
	var y = 100
	var z = true
	fmt.Println(y)
	fmt.Println(z)


	//短变量声明（只能在函数内部使用）
	nnz := "嘿嘿"
	fmt.Println(nnz)
}

