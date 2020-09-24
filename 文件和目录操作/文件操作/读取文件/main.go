package 读取文件

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	//读取文件（1）
	//1、只读方式打开文件file,err := os.Open()
	//打开文件
	file, err  := os.Open("F:/go-demo/time/main.go")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)
	//读取文件
	var strSlice []byte
	var tempSice = make([]byte,124)
	for {
		n , err := file.Read(tempSice)
		if err == io.EOF{ //err==io.EOF表示读取完毕
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		fmt.Printf("读取到了%v个字节\n",n)
		strSlice = append(strSlice,tempSice[:n]...)
	}
	fmt.Println(string(strSlice))
}


