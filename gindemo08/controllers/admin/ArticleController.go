package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gindemo08/models"
)

type ArticleController struct{
	BaseController  //继承baseController,可调用success方法
}


func (con ArticleController) Index (c *gin.Context) {
	articleList := []models.Article{}
	models.DB.Find(&articleList)

	c.JSON(http.StatusOK,gin.H{
		"result" : articleList,
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