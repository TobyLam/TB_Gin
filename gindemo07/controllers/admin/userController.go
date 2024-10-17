package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{

}

func (con UserController) Index(c *gin.Context) {
	c.String(http.StatusOK,"用户列表~~~~")
}

func (con UserController) Add(c *gin.Context) {
	c.String(http.StatusOK,"新增用户~~~~")
}

func (con UserController) Edit(c *gin.Context) {
	c.String(http.StatusOK,"编辑用户")
}