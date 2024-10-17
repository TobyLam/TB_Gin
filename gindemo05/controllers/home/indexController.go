package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func (con IndexController) Index (c *gin.Context) {
	c.HTML(http.StatusOK,"home/index.html",gin.H{
		"msg" : "欢迎光临璃月书集",
	})
}

func(con IndexController) News (c *gin.Context) {
	c.String(http.StatusOK,"新闻")
}
