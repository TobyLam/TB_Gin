package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
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

func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	file,err := c.FormFile("face")

	// file.Filename 获取文件名称
	dst := path.Join("./static/upload",file.Filename)
	if err == nil {
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
