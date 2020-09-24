package main

import (
	"fmt"
	"io"
	"os"
)


func CopyFile(srcFileName string,dstFileName string) error  {
	sFile, err  := os.Open(srcFileName)
	defer sFile.Close()
	dFile,err2 := os.OpenFile(dstFileName,os.O_CREATE|os.O_WRONLY,0666)
	defer sFile.Close()
	if err != nil {
		return err
	}
	if err2 != nil {
		return err
	}
	var tempSlice = make([]byte,128)
	for  {
		//读取数据
		_, err = sFile.Read(tempSlice)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}//

		//写入数据
		_, e2 := dFile.Write(tempSlice[:1])
		if e2 != nil {
			return e2
		}
	}
	return nil
}
func main()  {
	src := "F:/test.txt"
	dst := "D:/test.txt"
	err := CopyFile(src,dst)
	if err == nil {
		fmt.Println("拷贝完成\n")
	}else {
		fmt.Println("失败\n",err)
	}

}

