package service

import (
	"go-demo/dao/db"
	"go-demo/model"
)

//获取所以分类
func GetAllCatgoryList()(categoryList []*model.Category,err error)  {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		return
	}
	return
}


