package main

import(
	"gindemo07/routers"
	"github.com/gin-gonic/gin"
	"html/template"
	_"net/http"
	_"time"
	"gindemo07/models"
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

	//路由分组
	routers.HomeRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.AdminRoutersInit(r)

	r.Run()
}
