package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建评论
func CreateComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create Comment",
	})
}

// 获取评论列表
func GetCommentList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Comment List",
	})
}

// 获取评论详情
func GetCommentDetail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Comment Detail",
	})
}
