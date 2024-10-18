package routers

import (
	"gindemo08/controllers/admin"
	"gindemo08/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//middlewares.InitMiddleware中间件
	adminRouters := r.Group("/admin",middlewares.InitMiddleware)
	{
		adminRouters.GET("/", admin.IndexController{}.Index)

		//方法不能加括号，表示注册方法，加了括号表示执行
		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.POST("/user/doUpload", admin.UserController{}.DoUpload)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.POST("/user/doEdit", admin.UserController{}.DoEdit)
		adminRouters.GET("/user/edit2", admin.UserController{}.Edit2)
		adminRouters.POST("/user/doEdit2", admin.UserController{}.DoEdit2)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)
	}
}
