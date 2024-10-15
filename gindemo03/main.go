package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct{
	Title string
	Content string
}

func main(){
	r := gin.Default()
	//加载模板 放在配置路由前面
    r.LoadHTMLGlob("templates/**/*")

	//前台
	r.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK,"default/index.html",gin.H{
			"title":"首页",
			"score":89,
			"hobby" : []string{"诸葛亮","周瑜","陆逊","荀彧"},
			"newsList" : []interface{}{
				&Article{
					Title:   "1111",
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


	r.Run()
}

