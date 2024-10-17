package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(c *gin.Context){
	//判断用户是否登录

	fmt.Println(time.Now())

	fmt.Println(c.Request.URL)

	c.Set("username","钟离")

	//定义一个goroutine统计日志
	cCp := c.Copy()
	go func(){
		time.Sleep(2*time.Second)
		//不可以直接使用 context ,需要复制一份
		//err : fmt.Println("Done! in path "+c.Request.URL.Path)

		fmt.Println("Done! in path "+cCp.Request.URL.Path)
	}()
}
