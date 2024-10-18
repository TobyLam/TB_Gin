package main

import(
	"gindemo08/routers"
	"github.com/gin-gonic/gin"
	"html/template"
	"gindemo08/models"

	"github.com/gin-contrib/sessions"
	_"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
)



func main(){
	//创建一个默认的路由引擎
	r := gin.Default()
	//自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime" : models.UnixToTime,
	})
	//加载模板
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录
	r.Static("/static","./static")

	//配置session中间件

	// 创建基于cookie的存储引擎，secret 参数是用于加密的密钥【可修改】
	//store := cookie.NewStore([]byte("secret"))
	//store是前面创建的存储引擎，可以替换成其他的存储引擎
	//mysession是session名称
	//r.Use(sessions.Sessions("mysession",store))

	// 创建基于redis的存储引擎
	store,_ := redis.NewStore(10,"tcp","192.168.1.18:6379","123456",[]byte("secret111"))
	r.Use(sessions.Sessions("mysession",store))

	//路由分组
	routers.HomeRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.AdminRoutersInit(r)

	r.Run()
}
