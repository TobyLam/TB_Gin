package admin

import (
	"fmt"
	"gindemo08/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
)

type UserController struct{

}

func (con UserController) Index(c *gin.Context) {

	//查询数据库
	/*userList := []models.User{}
	models.DB.Find(&userList)
	c.JSON(http.StatusOK,gin.H{
		"result" :userList,
	})*/

	//查询age大于20的用户
	userList := []models.User{}
	models.DB.Where("age>20").Find(&userList)
	c.JSON(http.StatusOK,gin.H{
		"result":userList,
	})

}

func (con UserController) Add(c *gin.Context) {

	user := models.User{
		Username: "符玄",
		Age:      50,
		Email:    "taibusi@luofu.com",
		AddTime:  int(models.GetUnix()),
	}
	models.DB.Create(&user)
	fmt.Println(user)

	c.String(http.StatusOK,"增加数据成功")

	//c.HTML(http.StatusOK,"admin/useradd.html",gin.H{})
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
	//保存所有字段

	////查询id等于6的数据
	//user := models.User{Id:3}
	//models.DB.Find(&user)
	////更新数据
	//user.Username = "太卜：符玄"
	//user.Email = "TaiBu@luofu.com"
	//user.AddTime = int(models.GetUnix())
	//models.DB.Save(&user)

	//更新单个列
	/*user := models.User{}
	models.DB.Model(&user).Where("id = ?",3).Update("age","30")*/

	user := models.User{}
	models.DB.Where("id = ?",2).Find(&user)
	user.Username = "彦卿"
	user.Age = 18
	user.AddTime = int(models.GetUnix())

	models.DB.Save(&user)

	//c.HTML(http.StatusOK,"admin/useredit.html",gin.H{})
}

func (con UserController) Delete(c *gin.Context){

	//user := models.User{Id:4}
	//models.DB.Delete(&user)

	//删除数据
	user := models.User{}
	models.DB.Where("username = ?","李素裳").Delete(&user)
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
