package db

import "testing"

func init() {
	//dns := "root:root@tcp(192.168.2.32:3306)/blogger?charset=utf8mb4&parseTime=True"
	//err := Init(dns)
	//if err != nil {
	//	panic(err)
	//}
	err := Init("root:root@tcp(192.168.2.32:3306)/blogger?parseTime=true")
	if err != nil {
		return
	}
}

//测试获取单个分类信息
func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		return
	}
	t.Logf("category:%#v", category)
}

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	for _, v := range categoryList {
		t.Logf("id:%d category:%#v\n", v.CategoryId, v)
	}
}

func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil {
		return
	}
	for _, v := range categoryList {
		t.Logf("id:%d category:%#v\n", v.CategoryId, v)
	}
}
