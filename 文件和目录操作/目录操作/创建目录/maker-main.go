package 创建目录

import (
	"fmt"
	"os"
)

func main()  {
	//创建目录
	err := os.Mkdir("./abc", 0666)
	if err != nil {
		fmt.Println(err)
	}

	err :=  os.MkdirAll("./abcd/adf/sdfa/",0666)//创建多级
	if err != nil {
		fmt.Println(err)
	}
}
