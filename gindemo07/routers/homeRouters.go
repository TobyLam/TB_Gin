package routers

import (
	"gindemo07/controllers/home"
	"github.com/gin-gonic/gin"
)

func HomeRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", home.IndexController{}.Index)
		defaultRouters.GET("/news", home.IndexController{}.News)
		defaultRouters.GET("/shop", home.IndexController{}.Shop)
		defaultRouters.GET("/deleteCookie", home.IndexController{}.DeleteCookie)
	}
}
