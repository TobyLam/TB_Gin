package home

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func (con IndexController) Index (c *gin.Context) {
	//设置sessions
	session := sessions.Default(c)
	//配置session的过期时间
	session.Options(sessions.Options{
		MaxAge:3600*6,//6hrs 单位；秒
	})
	session.Set("username","张良111")
	session.Save() //设置session的使用必须调用

	//fmt.Println(models.UnixToTime(1629788418))
	c.HTML(http.StatusOK,"home/index.html",gin.H{
		"msg" : "欢迎光临璃月书集",
		"t":1629788418,
	})
}

func(con IndexController) News (c *gin.Context) {
	//获取sessions
	session := sessions.Default(c)

	username := session.Get("username")
	c.String(http.StatusOK,"新闻username=%v",username)
}


