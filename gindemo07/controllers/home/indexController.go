package home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func (con IndexController) Index (c *gin.Context) {
    //设置cookie
    c.SetCookie("username","张三",3600,"/","localhost",false,true)

    //多个二级域名共享cookie
    c.SetCookie("coo1","李四",3600,"/","a.yixuan.com",false,false) //只有a.yixuan.com可以拿到
    c.SetCookie("coo2","王五",3600,"/",".yixuan.com",false,false) //a.yixuan.com 、b.yixuan.com 都可以拿到

    //过期时间延时
    c.SetCookie("hobby","eat sleep",5, "/","localhost",false, true)

	//fmt.Println(models.UnixToTime(1629788418))
	c.HTML(http.StatusOK,"home/index.html",gin.H{
		"msg" : "欢迎光临璃月书集",
		"t":1629788418,
	})
}

func(con IndexController) News (c *gin.Context) {
	//获取cookie
	username,_ := c.Cookie("username")
	hobby,_ := c.Cookie("hobby")
	coo1,_ := c.Cookie("coo1")
	coo2,_ := c.Cookie("coo2")
	c.String(http.StatusOK,"新闻username=%v---hobby=%v---coo1=%v---coo2=%v",username,hobby,coo1,coo2)
}


func(con IndexController) Shop (c *gin.Context) {
	//获取cookie
	username,_ := c.Cookie("username")
	hobby,_ := c.Cookie("hobby")
	c.String(http.StatusOK,"新闻username=%v---hobby=%v",username,hobby)
}

func(con IndexController) DeleteCookie (c *gin.Context) {
	//删除cookie
	c.SetCookie("username","张三",-1,"/","localhost",false,true)
	c.String(http.StatusOK,"删除成功")
}

