package controller

import (
	"github.com/gin-gonic/gin"
	"go-demo/service"
	"net/http"
	"strconv"
)

//访问主页的控制器
func IndexHandler(c *gin.Context)  {
	//从service取数据
	//1，加载了文章数据
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
	c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	//2.分类数据
	categoryList, err := service.GetAllCatgoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	//var data map[string]interface{} = make(map[string]interface{},16)
	c.HTML(http.StatusOK,"views/index.html",gin.H{
	"article_list":articleRecordList,
	"category_list":categoryList,
	})
}

//点击分类云进行分类
func CategoryList(c *gin.Context)  {
	catqueryIdStr := c.Query("category_id")
	//转
	catqueryId, err := strconv.ParseInt(catqueryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	//根据分类ID去获取文章列表
	articleRecordList, err := service.GetArticleRecordListById(int(catqueryId), 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	//再次加载所有分类数据，用于分类云显示
	categoryList, err := service.GetAllCatgoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	c.HTML(http.StatusOK,"views/index.html",gin.H{
		"article_list":articleRecordList,
		"category_list":categoryList,
	})
}
