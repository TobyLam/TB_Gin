package main

import(
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

type Article struct{
	Title string  `xml:"title"`
	Content string `xml:"content"`
}

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

//时间戳转换成日期
func UnixToTime(timestamp int) string{
	fmt.Println(timestamp)
	t:=time.Unix(int64(timestamp),0)
	return t.Format("2006-01-02 15:04:05")
}

func Println(str1 string,str2 string) string{
	fmt.Println(str1,str2)
	return str1 +"---"+ str2
}

func main(){
	r := gin.Default()

	//自定义模板函数 注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime":UnixToTime,
		"Println":Println,
	})
	//加载模板 放在配置路由前面
    r.LoadHTMLGlob("templates/**/*")
    //配置静态web目录 第一个参数表示路由，第二个参数表示映射的目录。
	r.Static("/static","./static")

	//前台
	r.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK,"default/index.html",gin.H{
			"title":"首页",
			"msg":"success",
			"score":89,
			"hobby" : []string{"诸葛亮","周瑜","陆逊","荀彧"},
			"newsList" : []interface{}{
				&Article{
					Title:   "11112",
					Content: "1111",
				},
				&Article{
					Title:   "222",
					Content: "222",
				},
			},
			"testSlice":[]string{},
			"news":&Article{
				Title:   "新闻标题",
				Content: "新闻内容",
			},
			"date":1629423555,
		})
	})
	r.GET("/news",func(c *gin.Context){
		news := &Article{
			Title:   "原神",
			Content: "启动",
		}
		c.HTML(http.StatusOK,"default/news.html",gin.H{
			"title":"新闻页面",
			"news" : news,
		})
	})
	//后台
	r.GET("/admin",func(c *gin.Context){
		c.HTML(http.StatusOK,"admin/index.html",gin.H{
			"title":"后台首页",
		})
	})
	r.GET("/admin/news",func(c *gin.Context){
		c.HTML(http.StatusOK,"admin/news.html",gin.H{
			"title":"新闻页面",
		})
	})

	//Get传值
	r.GET("/testget",func(c *gin.Context){
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page","1")
		c.JSON(http.StatusOK,gin.H{
			"username" : username,
			"age" : age,
			"page" : page,
		})
	})

	//Get传值
	r.GET("/testget01",func(c *gin.Context){

		id := c.DefaultQuery("id","1")
		c.JSON(http.StatusOK,gin.H{
			"msg" : "新闻详情",
			"id" : id,
		})
	})

	//post请求
	r.GET("/user",func(c *gin.Context){
		c.HTML(http.StatusOK,"default/user.html",gin.H{})
	})
	//获取表单post过来的数据
	r.POST("/doAddUser1", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		age := context.DefaultPostForm("age","20")

		context.JSON(http.StatusOK,gin.H{
			"username" : username,
			"password" : password,
			"age" : age,
		})
	})

	//获取 GET POST 传递的数据绑定到结构体
	r.GET("/getUser",func(c *gin.Context){
		user := &UserInfo{}
		if err := c.ShouldBind(&user);err == nil {
			fmt.Printf("%#v",user)
			c.JSON(http.StatusOK,user)
		}else{
			c.JSON(http.StatusOK,gin.H{
				"err": err.Error(),
			})
		}
	})

	r.POST("/doAddUser2", func(context *gin.Context) {

		user := &UserInfo{}
		if err := context.ShouldBind(&user);err == nil {
			fmt.Printf("%#v",user)
			context.JSON(http.StatusOK,user)
		}else{
			context.JSON(http.StatusOK,gin.H{
				"err": err.Error(),
			})
		}
	})

	// 获取 Post Xml数据
	r.POST("/xml",func(c *gin.Context){
		xmlSliceData, _ := c.GetRawData() //获取 c.Request.Body 读取请求的数据

		article := &Article{}

		fmt.Println(xmlSliceData)

		if err := xml.Unmarshal(xmlSliceData,&article); err == nil {
			c.JSON(http.StatusOK,article)
		}else{
			c.JSON(http.StatusOK,gin.H{
				"err": err.Error(),
			})
		}
	})

	//动态路由传值
	//list/123 list/456
	r.GET("/list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(http.StatusOK,"%v",cid)
	})

	r.Run()
}

