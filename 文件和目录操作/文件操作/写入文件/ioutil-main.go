package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	str := "halou"
	err := ioutil.WriteFile("F:/test.txt",[]byte(str),0666)
	if err != nil{
		fmt.Println("写入失败",err)
		return
	}
}
