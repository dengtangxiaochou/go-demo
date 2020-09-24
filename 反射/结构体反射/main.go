package main

import (
	"fmt"
	"reflect"
)

//student结构体
type student struct {
	Name string`json:"name"`
	Age int`json:"age"`
	Score int`json:"int"`
}

func (s student) GetInfo()string  {
	var str = fmt.Sprintf("姓名：%v 年龄：%v  c成绩：%v\n",s.Name,s.Age,s.Score)
	return str
}

func (s *student) SetInfo(name string,age int,score int)  {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s student) Print()  {
	fmt.Println("这是一个打印方法。。。")
}

func PrintStructField(s interface{})  {
	//判断你传过来的值是否是结构体
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() !=reflect.Struct && t.Elem().Kind() != reflect.Struct  {
		fmt.Println("传入参数是不是一个结构体")
		return
	}
	//1、通过类型变量里面的field可以获取结构体的字段
	Field0 := t.Field(0)
	//fmt.Printf("%#v",Field0)
	fmt.Printf("字段名称：%v\n",Field0.Name)
	fmt.Printf("字段类型：%v\n",Field0.Type)
	fmt.Printf("字段Tag：%v\n",Field0.Tag.Get("json"))
	fmt.Printf("字段Tag：%v\n",Field0.Tag.Get("form"))
	fmt.Println("-----------------------------------------")
	//2、通过类型变量FieldByName可以获取结构体的字段
	Field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Printf("字段名称：%v\n",Field1.Name)
		fmt.Printf("字段类型：%v\n",Field1.Type)
		fmt.Printf("字段Tag：%v\n",Field1.Tag.Get("json"))
	}
	fmt.Println(Field1)
	//3、通过类型变量的NumField获取到结构体有几个字段
	var fieldCount = t.NumField()
	fmt.Println("结构体字段",fieldCount,"个")
	fmt.Println("-----------------------")
	//通过值变量获取结构体属性对应的值
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
	for i := 0; i<fieldCount ; i++ {
		//defer func() {
		//	//捕获test抛出的panic
		//	if err := recover();err != nil {
		//		fmt.Println("for  发生错误",err)
		//	}
		//}()
		fmt.Printf("属性名称：%v 属性值：%v 属性类型：%v 属性Tag：%v\n ",t.Field(i).Name,v.Field(i),t.Field(i).Type,t.Field(i).Tag.Get("json"))
		fmt.Println("========================")
	}
}

//打印方法
func PrintDStructFn(s interface{})  {
	//判断你传过来的值是否是结构
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() !=reflect.Struct && t.Elem().Kind() != reflect.Struct  {
		fmt.Println("传入参数是不是一个结构体")
		return
	}
	//1、通过类型变量里面的Method可以获取结构体的方法
	method0 := t.Method(0)
	fmt.Println(method0.Name)
	fmt.Println(method0)
	fmt.Println(method0.Type)
	fmt.Println("-----------------------------")
	//2、通过类型变量获取这个结构体有多少个方法
	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(method1.Name)//Print
		fmt.Println(method1.Type)//func(main.student)
	}
	fmt.Println("----------------------------")

	//3、通过《值变量》执行方法
	//v.Method(1).Call(nil)
	v.MethodByName("Print").Call(nil)
	Info := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(Info)
	fmt.Println("-----------------------------")

	//4、传参
	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(30))
	params = append(params, reflect.ValueOf(99))
	v.MethodByName("SetInfo").Call(params)//执行方法传入参数

	info2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info2)

	//5、获取方法数量
	fmt.Println(t.NumMethod())

}

func main()  {
	stu1 := student{
		Name:  "小木",
		Age:   20,
		Score: 100,
	}
	PrintStructField(stu1)
	PrintDStructFn(&stu1)
}
