package routers

import (
	"gindemo05/controllers/home"
	"github.com/gin-gonic/gin"
)

func HomeRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", home.IndexController{}.Index)
		defaultRouters.GET("/news", home.IndexController{}.News)
	}
}
