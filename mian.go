package main

import (
	"errors"
	"fmt"
)

//func fn1()  {
//	fmt.Println("fn1")
//}
//
//func fn2()  {
//	defer func() {
//		err := recover()
//		if err != nil {
//			fmt.Println("err:",err)
//		}
//	}()
//	panic("抛出异常")
//}
//
//func main()  {
//	fn1()
//	fn2()
//	fmt.Println("结束")
//}

func readFile(fileName string) error  {
	if fileName == "main.go"{
		return nil
	}else {
		return errors.New("读取失败")
	}
}

func myFn()  {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("给管理员发邮件")
		}
	}()
	err := readFile("xxx.go")
	if err != nil {
		panic(err)
	}
}

func main()  {
	myFn()
	fmt.Println("继续")
}
