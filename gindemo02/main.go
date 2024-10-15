package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

type Hero struct {
	Name string `json:"name"`
    Alias string `json:"alias"`
}

func main(){
	r := gin.Default()

	//配置模板的文件
	r.LoadHTMLGlob("templates/*")

	r.GET("/",func(c *gin.Context){
		c.String(http.StatusOK,"值：%v","首页")
	})
	r.GET("/json1",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"success":true,
			"msg":"hello,gin",
		})
	})
	r.GET("/json2",func(c *gin.Context){
		c.JSON(http.StatusOK,map[string]interface{}{
			"success":true,
			"msg":"hello,world",
		})
	})

	r.GET("/json3",func(c *gin.Context){
		a:=&Hero{
			Name:  "宋江",
			Alias: "及时雨",
		}
		c.JSON(http.StatusOK,a)
	})
	//响应Jsonp请求
	//http://localhost:8082/jsonp?callback=sss
	//sss({"name":"卢俊义","alias":"玉麒麟"});
	r.GET("/jsonp",func(c *gin.Context){
		a:=&Hero{
			Name:  "卢俊义",
			Alias: "玉麒麟",
		}
		c.JSONP(http.StatusOK,a)
	})

	r.GET("/xml",func(c *gin.Context){
		c.XML(http.StatusOK,gin.H{
			"success":true,
			"msg":"hello,xml",
		})
	})
	r.GET("/xml2",func(c *gin.Context){
		//使用结构体返回xml
		type XmlContent struct {
			Success bool `xml:"success"`
			Msg string `xml:"msg"`
		}
		xmlC := &XmlContent{
			Success: true,
			Msg:     "hello,xml2",
		}
		c.XML(http.StatusOK,xmlC)
	})

	r.GET("/news",func(c *gin.Context){
		//r.LoadHTMLGlob("templates/*)
		c.HTML(http.StatusOK,"news.html",gin.H{
			"title":"欲买桂花同载酒",
		})
	})

	r.GET("/goods",func(c *gin.Context){
		//r.LoadHTMLGlob("templates/*)
		c.HTML(http.StatusOK,"goods.html",gin.H{
			"title":"为了纳塔",
			"name":"玛薇卡",
		})
	})


	r.Run(":8080")
}
