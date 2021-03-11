package controller

import (
	"crypto/md5"
	"fmt"
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
		c.JSON(200, gin.H{"code": 101, "err": "username exist"})
	}

	// 密码md5处理下
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))

	// 新账号写入DB
	newAccount := model.UserAccount{Username: username, Password: password}
	model.NewAccount(newAccount)

	c.JSON(200, gin.H{"code": 0, "data": gin.H{}})
}
