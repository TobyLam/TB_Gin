package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

// 模板函数
func UnixToTime(timeStamp int) string{
	t := time.Unix(int64(timeStamp),0)
	return t.Format("2006-01-02 15:04:05")
}

func initMiddleware(c *gin.Context){
	start := time.Now().UnixNano()
	fmt.Println("1-我是一个中间件")

	//调用该请求的剩余处理程序
	c.Next()

	//终止该请求剩余处理程序
	//直接不执行请求后面的程序，而是执行中间件中的代码
	//c.Abort()

	fmt.Println("2-我也是一个中间件")
	end := time.Now().UnixNano()

	fmt.Println(end-start)
}

func initMiddlewareOne(c *gin.Context){
	fmt.Println("1-我是一个中间件--initMiddlewareOne")

	//调用该请求的剩余处理程序
	c.Next()

	//终止该请求剩余处理程序
	//直接不执行请求后面的程序，而是执行中间件中的代码
	//c.Abort()

	fmt.Println("2-我也是一个中间件--initMiddlewareOne")
}

func initMiddlewareTwo(c *gin.Context){
	fmt.Println("1-我是一个中间件--initMiddlewareTwo")

	//调用该请求的剩余处理程序
	c.Next()

	//终止该请求剩余处理程序
	//直接不执行请求后面的程序，而是执行中间件中的代码
	//c.Abort()

	fmt.Println("2-我也是一个中间件--initMiddlewareTwo")
}

//全局中间件
func initMiddlewareGlobOne(c *gin.Context){
	fmt.Println("1-我是一个中间件--initMiddlewareGlobOne")
	//调用该请求的剩余处理程序
	c.Next()

	fmt.Println("2-我也是一个中间件--initMiddlewareGlobOne")
}
func initMiddlewareGlobTwo(c *gin.Context){
	fmt.Println("1-我是一个中间件--initMiddlewareGlobTwo")
	//调用该请求的剩余处理程序
	c.Next()

	fmt.Println("2-我也是一个中间件--initMiddlewareGlobTwo")
}

func main(){
	//默认引擎
	r := gin.Default()
	//自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime":UnixToTime,
	})
	//加载模板
	r.LoadHTMLGlob("templates/**/*")
	//配置静态web目录
	r.Static("/static","./static")


	/*
	//使用中间件
	//以下例子：使用中间件统计程序执行时间
	r.GET("/",initMiddleware, func(c *gin.Context){
		fmt.Println("我是一个首页")
		//time.Sleep(time.Second)
		c.String(200,"gin首页")
	})
	*/

	/*
	//同时使用多个中间件
	r.GET("/",initMiddlewareOne,initMiddlewareTwo, func(c *gin.Context){
		fmt.Println("我是一个首页")
		//time.Sleep(time.Second)
		c.String(200,"gin首页")
	})
	*/
	//使用全局中间件
	r.Use(initMiddlewareGlobOne,initMiddlewareGlobTwo)

	r.GET("/", func(c *gin.Context){
		fmt.Println("我是一个首页")
		//time.Sleep(time.Second)
		c.String(200,"gin首页")
	})


	r.GET("/news",func(c *gin.Context){
		c.String(200,"新闻页面")
	})

	r.GET("/login",func(c *gin.Context){
		c.String(200,"login")
	})


	r.Run()
}
