package db

import (
	"fmt"
	"go-demo/model"
	"testing"
	"time"
)

func init() {
	dns := "root:root@tcp(192.168.2.32:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

//插入文章
func TestInsertArticle(t *testing.T) {
	//构建对象
	article := &model.ArticleDetail{
		ArticleInfo: model.ArticleInfo{
			Id:           1,
			CategoryId:   0,
			Summary:      "abc fd",
			Title:        "5qi go",
			ViewCount:    1,
			CreateTime:   time.Now(),
			CommentCount: 0,
			Username:     "sum",
		},
		Content:     "abc",
		Category:    model.Category{
			CategoryId:   0,
			CategoryName: "",
			CategoryNo:   0,
		},
	}
	articleId, err := InsertArticle(article)
	if err != nil {
		panic(err)
	}
	t.Logf("articleId: %d\n",articleId)
}

//查询
func TestGetAricleList(t *testing.T) {
	categoryList, err := GetAricleList(1,15)
	if err != nil {
		return
	}
	t.Logf("article: %v\n",categoryList)
}

func TestGetArticleDetail(t *testing.T) {
	articleDetail, err := GetArticleDetail(1)
	if err != nil {
		return
	}
	fmt.Println(articleDetail)
}