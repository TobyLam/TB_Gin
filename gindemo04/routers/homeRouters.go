package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK,"home/index.html",gin.H{
				"msg" : "欢迎光临璃月书集",
			})
		})

		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK,"新闻")
		})
	}
}
