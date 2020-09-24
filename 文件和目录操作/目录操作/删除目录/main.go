package main

import (
	"fmt"
	"os"
)

func main()  {
	//删除文件
	//err := os.Remove("aaa.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//删除目录
	err := os.RemoveAll("./abcd")
	if err != nil {
		fmt.Println(err)
	}
}
