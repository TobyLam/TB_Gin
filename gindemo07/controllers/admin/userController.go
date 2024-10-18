package admin

import (
	"fmt"
	"gindemo07/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
)

type UserController struct{

}

func (con UserController) Index(c *gin.Context) {
	c.String(http.StatusOK,"用户列表~~~~")
}

func (con UserController) Add(c *gin.Context) {
	//c.String(http.StatusOK,"新增用户~~~~")
	c.HTML(http.StatusOK,"admin/useradd.html",gin.H{})
}

/**
 1.获取上传的文件
 2.获取后缀名 判断类型是否正确 .jpg .png .gif .jpeg
 3.创建图片保存目录 static/upload/20220202
 4.生成文件名称和文件保存的目录
 5.执行上传
 */
func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	//1.获取上传的文件
	file,err := c.FormFile("face")
	dst := ""
	if err == nil {
		//2.获取后缀名 判断类型是否正确 .jpg .png .gif .jpeg
		extName := path.Ext(file.Filename)

		allowExtMap := map[string]bool{
			".jpg" : true,
			".png" : true,
			".gif" : true,
			".jpeg": true,
		}
		if _,ok := allowExtMap[extName]; !ok {
			c.String(http.StatusOK,"上传的文件类型不合法")
			return
		}

		//3.创建图片保存目录 static/upload/20220202
		day := models.GetDay()
		dir := path.Join("./static/upload",day)

		err = os.MkdirAll(dir,0666)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusOK,"MkdirAll失败")
			return
		}

		//4.生成文件名称和文件保存的目录
		fileName := strconv.FormatInt(models.GetUnix(),10) + extName

		//5.执行上传
		dst = path.Join(dir,fileName) // file.Filename 获取文件名称
		c.SaveUploadedFile(file,dst)
	}
	//c.String(http.StatusOK,"执行上传")
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"username":username,
		"dst":dst,
	})
}

func (con UserController) Edit(c *gin.Context){
	c.HTML(http.StatusOK,"admin/useredit.html",gin.H{})
}

func (con UserController) DoEdit(c *gin.Context){
	username := c.PostForm("username")
	face1,err1 := c.FormFile("face1")
	dst1 := path.Join("./static/upload",face1.Filename)
	if err1 == nil{
		c.SaveUploadedFile(face1,dst1)
	}

	face2,err2 := c.FormFile("face2")
	dst2 := path.Join("./static/upload",face2.Filename)
	if err2 == nil{
		c.SaveUploadedFile(face2,dst2)
	}

	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"username":username,
		"dst1":dst1,
		"dst2":dst2,
	})

}

func (con UserController) Edit2(c *gin.Context){
	c.HTML(http.StatusOK,"admin/useredit2.html",gin.H{})
}

func (con UserController) DoEdit2(c *gin.Context){
	username := c.PostForm("username")

	form,_ := c.MultipartForm()
	files := form.File["face[]"]

	for _,file := range files{
		dst := path.Join("./static/upload",file.Filename)
		//上传文件至指定目录
		c.SaveUploadedFile(file,dst)
	}

	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"username":username,
	})

}
