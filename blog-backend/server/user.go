package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册用户
func RegisterUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Register User",
	})
}

// 用户登录
func LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login User",
	})
}
