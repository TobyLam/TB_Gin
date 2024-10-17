package home

import (
	"fmt"
	"gindemo07/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func (con IndexController) Index (c *gin.Context) {

	fmt.Println(models.UnixToTime(1629788418))
	c.HTML(http.StatusOK,"home/index.html",gin.H{
		"msg" : "欢迎光临璃月书集",
		"t":1629788418,
	})
}

func(con IndexController) News (c *gin.Context) {
	c.String(http.StatusOK,"新闻")
}
