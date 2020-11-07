package db

import (
	_ "github.com/go-sql-driver/mysql"
	"go-demo/model"
)

//插入文章
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	//验证
	if article == nil {
		return
	}
	sqlStr := `INSERT INTO 
				article(content,summary,title,username,category_id,view_count,comment_count)values(?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlStr, article.Content, article.Summary, article.Title, article.Username, article.ArticleInfo.CategoryId,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

//获取文章列表，做个分页
func GetAricleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize <= 0 {
		return
	}
	sqlStr := `SELECT 
					id,summary,title,view_count,create_time,comment_count,username,category_id
				from
					article
				where
					status = 1
				order by create_time desc
				limit ?,?`
	err = DB.Select(&articleList, sqlStr, pageNum, pageSize)
	return
}

//根据文章Id，查询单个文章
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		return
	}
	sqlStr := `SELECT 
					id,summary,title,view_count,create_time,comment_count,username,category_id
				from
					article
				where
					id = ?
				and
					status = 1
`
	err = DB.Get(&articleDetail, sqlStr, articleId)
	return
}

//根据分类Id,查询这个一类的文章
func GetArticleListByGategoryId(categoryId,pageNum,pageSize int)(articleList[]*model.ArticleInfo,err error)  {
	if pageNum < 0 || pageSize <= 0 {
		return
	}
	sqlStr := `SELECT 
					id,summary,title,view_count,create_time,comment_count,username,category_id
				from
					article
				where
					status = 1	
				and
					category_id = ?
				order by create_tome desc
				limit ?,?`
	DB.Select(&articleList,sqlStr,categoryId,pageNum,pageSize)
	return
}