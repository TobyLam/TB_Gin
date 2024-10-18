package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}


func (con IndexController)Index(c *gin.Context) {

	//获取中间件中的值，类型为interface{}
	username,_ := c.Get("username")
	fmt.Println(username)

	//使用类型断言
	v,ok:=username.(string)
	if ok {
		c.String(http.StatusOK,"后台首页"+v)
	}else{
		c.String(http.StatusOK,"后台首页失败")
	}

}