package main

import "fmt"

func fn1(x int)  {
	x = 10
}

func fn2(x *int)  {
	*x = 40
}



func main()  {
	var a  = 10

	var p  = &a
	fmt.Printf("a的值：%v\n a的类型%T\n a的地址%p\n",a,a,&a)
	fmt.Printf("p的值：%v\n p的类型%T\n p的地址%p\n",p,p,&p)
	//取值
	fmt.Println(*p)

	*p  = 30
	fmt.Println(a)

	var c  = 5
	fn1(c)
	fmt.Println(c)
	fn2(&c)
	fmt.Println(c)

}
