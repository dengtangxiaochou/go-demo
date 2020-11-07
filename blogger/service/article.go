package service

import (
	"go-demo/dao/db"
	"go-demo/model"
)

//获取文章和对应的分类
func GetArticleRecordList(pageNum, pagesSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//1.获取文章的列表
	articleInfoList, err := db.GetAricleList(pageNum, pagesSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}

	//2.获取文章对应的分类（多个）
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	//聚合
	for _, article := range articleInfoList {
		//根据当前文章，生产结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		//从这个文章中取出分类ID
		categoryId := article.CategoryId
		//遍历分类列表
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break



			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

//根据多个文章的id，获取多个分类的id的集合
func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
	LABEL:
	//遍历文章，得到每个文章
	for _, article := range articleInfoList {
		//从当前文章取出分类id
		categoryId := article.CategoryId
		//去重
		for _, id := range ids {
			//看当前Id是否存在
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids,categoryId)
	}
	return
}

//根据分类ID，获取该文章和他们对应的分类信息
func GetArticleRecordListById(categoryId, pageNum, pagesSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//1.获取文章的列表
	articleInfoList, err := db.GetArticleListByGategoryId(categoryId, pageNum, pagesSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}

	//2.获取文章对应的分类（多个）
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	//聚合
	for _, article := range articleInfoList {
		//根据当前文章，生产结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		//从这个文章中取出分类ID
		categoryId := article.CategoryId
		//遍历分类列表
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}
