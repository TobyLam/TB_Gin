package main

// 导入gin
import(
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"log"
	"net/http"
)

// 自定义go 中间件 （拦截器）
func myHandler()(gin.HandlerFunc){
	return func(context *gin.Context) {
		// 通过自定义的中间件,设置的值，在后续处理只要调用了这个中间件的都可以拿到这里的参数
		context.Set("usersession","userid-1")


		context.Next() //放行


		//context.Abort() //阻止
	}
}

func main(){
	// 创建一个服务
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./favicon.ico"))
	// 注册中间件
	ginServer.Use(myHandler())

	// 加载静态页面
	ginServer.LoadHTMLGlob("templates/*")
	// 加载资源文件
	ginServer.Static("/static","./static")

	// 连接数据库的代码


    // 访问地址，处理请求  Request Response
    /*ginServer.GET("/hello",func(context *gin.Context) {
    	context.JSON(200,gin.H{
    		"msg":"hello,world",
		})
	})*/

	// 响应一个页面给前端
	ginServer.GET("/index",func(c *gin.Context){
		//c.JSON() json数据
		c.HTML(http.StatusOK,"index.html",gin.H{
			"msg":"王离月干年",
		})
	})

	// 接收前端传过来的参数
	// usl?userid=xxx&username=toby
	// 如果不指定中间件，myHandler()，则中间件为全局使用，否则只在/user/info使用
	ginServer.GET("/user/info",myHandler(),func(c *gin.Context){

		// 取出中间件中的值 key
		usersession := c.MustGet("usersession").(string)
		log.Println("=====>",usersession)

		userid := c.Query("userid")
		username := c.Query("username")
		c.JSON(http.StatusOK,gin.H{
			"userid":userid,
			"username":username,
		})
	})

	// /user/info/1/toby
	ginServer.GET("user/info/:userid/:username",func(c *gin.Context){
		userid := c.Param("userid")
		username := c.Param("username")
		c.JSON(http.StatusOK,gin.H{
			"userid":userid,
			"username":username,
		})
	})

	// 前端给后端传json
	ginServer.POST("/json",func(c *gin.Context){
		// request.body
		data,_ := c.GetRawData()

		var m map[string]interface{}
		// 序列化包装为json []byte
		_ = json.Unmarshal(data,&m)

		c.JSON(http.StatusOK,m)
	})

	// 支持函数式编程 =>
	ginServer.POST("/user/add",func(context *gin.Context){
		username := context.PostForm("username")
		password := context.PostForm("password")

		context.JSON(http.StatusOK,gin.H{
			"msg":"ok",
			"username":username,
			"password":password,
		})
	})

	// 路由
	ginServer.GET("/test",func(context *gin.Context){
		//重定向
		context.Redirect(http.StatusMovedPermanently,"https://www.baidu.com")
	})

	// 404 NoRoute
	ginServer.NoRoute(func(context *gin.Context){
		context.HTML(http.StatusNotFound,"404.html",nil)
	})

	// 路由组
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.POST("/login")
		userGroup.POST("/logout")
	}

	orderGroup := ginServer.Group("/order")
	{
		orderGroup.GET("/add")
		orderGroup.DELETE("/delete")
	}


	// Gin RestFul
	/*
	ginServer.POST("/user",func(c *gin.Context){
		c.JSON(200,gin.H{"msg":"post.User"})
	})
	ginServer.PUT("/user")
	ginServer.DELETE("/user")
	*/

	// 服务器端口
	ginServer.Run(":8082")
}
