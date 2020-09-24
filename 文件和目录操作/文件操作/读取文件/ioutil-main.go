package 读取文件

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	byteStr, err := ioutil.ReadFile("F:/go-demo/time/main.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteStr))
}
