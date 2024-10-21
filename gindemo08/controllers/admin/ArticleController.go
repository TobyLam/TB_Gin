package admin

import (
	"gindemo08/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct{
	BaseController  //继承baseController,可调用success方法
}


func (con ArticleController) Index (c *gin.Context) {
	//获取所有的文章
	//articleList := []models.Article{}
	//models.DB.Find(&articleList)
	//
	//c.JSON(http.StatusOK,gin.H{
	//	"result" : articleList,
	//})

	//查询文章，获取文章对应的分类
	//articleList := []models.Article{}
	//models.DB.Preload("ArticleCate").Find(&articleList)
	//
	//c.JSON(200,gin.H{
	//	"result":articleList,
	//})

	//获取所有的文章分类
	articleCateList := []models.ArticleCate{}
	models.DB.Preload("Article").Find(&articleCateList)
	c.JSON(200,gin.H{
		"result" : articleCateList,
	})

	//c.String(http.StatusOK,"文章")
	//con.success(c)
}

func (con ArticleController) Add (c *gin.Context) {
	c.String(http.StatusOK,"新增文章")
}

func (con ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK,"编辑文章")
}