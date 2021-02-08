package controller

import (
	"github.com/gin-gonic/gin"
	"helloGin/model"
	"net/http"
)

func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

// 注册逻辑
func RegisterPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 检查用户名是否已注册过
	if model.QueryAccount(username) != (model.UserAccount{}) {
		c.JSON(200, gin.H{"code": 101, "err": "username used"})
	}

}
