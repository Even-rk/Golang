package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建文章
func CreatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create Post",
	})
}

// 获取文章列表
func GetPostList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Post List",
	})
}

// 获取文章详情
func GetPostDetail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Post Detail",
	})
}
