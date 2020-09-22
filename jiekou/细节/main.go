package main

import "fmt"

type Address struct {
	Name string
	Phone int
}


func main()  {
	var userinfo  = make(map[string]interface{})

	userinfo["name"]  = "张三"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"睡觉","吃饭"}

	fmt.Println(userinfo["age"])
	fmt.Println(userinfo["hobby"])

	var address  = Address{
		Name:  "李四",
		Phone: 12134452435627,
	}
	fmt.Println(address.Name)

	userinfo["address"] = address

	fmt.Println(userinfo["address"])

	hobby2,_ := userinfo["hobby"].([]string)
	fmt.Println(hobby2[1])

	address2, _ := userinfo["address"].(Address)
	fmt.Println(address2.Name,address2.Phone)


}

