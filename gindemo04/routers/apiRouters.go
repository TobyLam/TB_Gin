package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK,"api接口")
		})

		apiRouters.GET("/userlist", func(c *gin.Context) {
			c.String(http.StatusOK,"api接口---userlist")
		})

		apiRouters.GET("/plist", func(c *gin.Context) {
			c.String(http.StatusOK,"api接口-----plist")
		})

		apiRouters.GET("/cart", func(c *gin.Context) {
			c.String(http.StatusOK,"api接口--cart")
		})
	}
}
