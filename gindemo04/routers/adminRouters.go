package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK,"后台首页1")
		})

		adminRouters.GET("/user", func(c *gin.Context) {
			c.String(http.StatusOK,"用户列表")
		})

		adminRouters.GET("/user/add", func(c *gin.Context) {
			c.String(http.StatusOK,"新增用户")
		})

		adminRouters.GET("/user/edit", func(c *gin.Context) {
			c.String(http.StatusOK,"编辑用户")
		})

		adminRouters.GET("/aritcle", func(c *gin.Context) {
			c.String(http.StatusOK,"首页1")
		})
	}
}
