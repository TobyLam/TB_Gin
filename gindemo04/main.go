package main

import(
	"gindemo04/routers"
	"github.com/gin-gonic/gin"
	"html/template"
	_"net/http"
	"time"
)

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp),0)
	return t.Format("2006-01-02 15:04:05")
}

func main(){
	//创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime" : UnixToTime,
	})
	//加载模板
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录
	r.Static("/static","./static")

	//路由分组
	routers.HomeRoutersInit(r)
	//defaultRouters := r.Group("/")
	//{
	//	defaultRouters.GET("/", func(c *gin.Context) {
	//		c.String(http.StatusOK,"首页1")
	//	})
	//
	//	defaultRouters.GET("/news", func(c *gin.Context) {
	//		c.String(http.StatusOK,"新闻")
	//	})
	//}

	routers.ApiRoutersInit(r)
	//apiRouters := r.Group("/api")
	//{
	//	apiRouters.GET("/", func(c *gin.Context) {
	//		c.String(http.StatusOK,"api接口")
	//	})
	//
	//	apiRouters.GET("/userlist", func(c *gin.Context) {
	//		c.String(http.StatusOK,"api接口---userlist")
	//	})
	//
	//	apiRouters.GET("/plist", func(c *gin.Context) {
	//		c.String(http.StatusOK,"api接口-----plist")
	//	})
	//
	//	apiRouters.GET("/cart", func(c *gin.Context) {
	//		c.String(http.StatusOK,"api接口--cart")
	//	})
	//}

	routers.AdminRoutersInit(r)
	//adminRouters := r.Group("/admin")
	//{
	//	adminRouters.GET("/", func(c *gin.Context) {
	//		c.String(http.StatusOK,"后台首页1")
	//	})
	//
	//	adminRouters.GET("/user", func(c *gin.Context) {
	//		c.String(http.StatusOK,"用户列表")
	//	})
	//
	//	adminRouters.GET("/user/add", func(c *gin.Context) {
	//		c.String(http.StatusOK,"新增用户")
	//	})
	//
	//	adminRouters.GET("/user/edit", func(c *gin.Context) {
	//		c.String(http.StatusOK,"编辑用户")
	//	})
	//
	//	adminRouters.GET("/aritcle", func(c *gin.Context) {
	//		c.String(http.StatusOK,"首页1")
	//	})
	//}


	r.Run()
}
