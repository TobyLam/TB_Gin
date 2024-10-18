package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiController struct{}

func (con ApiController) Index (c *gin.Context) {
	c.String(http.StatusOK,"api接口")
}

func (con ApiController) Userlist(c *gin.Context) {
	c.String(http.StatusOK,"api接口---userlist")
}

func (con ApiController) Plist(c *gin.Context) {
	c.String(http.StatusOK,"api接口-----plist")
}

func (con ApiController) Cart(c *gin.Context) {
	c.String(http.StatusOK,"api接口--cart")
}
