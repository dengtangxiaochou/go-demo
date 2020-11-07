package db

import (
	"github.com/jmoiron/sqlx"
	"go-demo/model"
)

//添加分类
func InstertCategory(category *model.Category) (categoryId int64, err error) {
	sqlStr := "insert into category(category_name,category_no) value(?,?)"
	result, err := DB.Exec(sqlStr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	categoryId, err = result.LastInsertId()
	return
}

//获取单个文章分类
func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	//sqlStr := `select id,category_name,category_no from category where id=?`
	//DB.Get(category, sqlStr, id)
	//return
	DB.Get(category, "SELECT id,category_name,category_no FROM category WHERE id = ?", id)
	return
}

//获取多个分类
func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	sqlStr, args, err := sqlx.In("SELECT id,category_name,category_no FROM category WHERE id IN(?)", categoryIds)
	//err = DB.Select(&categoryList, sqlStr, args...)
	if err != nil {
		return
	}
	//查询
	err = DB.Select(&categoryList, sqlStr, args...)
	return
}

//获取所有分类
func GetAllCategoryList() (categoryList []*model.Category, err error) {
	sqlStr := "select id,category_name,category_no from category order by category_no asc "
	err = DB.Get(&categoryList, sqlStr)
	return
}
