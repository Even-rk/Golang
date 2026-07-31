package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建文章
func CreatePost(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create Post",
	})
}

// 获取单个文章列表
func GetSinglePostList(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Single Post List",
	})
}

// 获取所有文章列表
func GetAllPostList(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get All Post List",
	})
}

// 更新文章
func UpdatePost(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Post",
	})
}

// 删除文章
func DeletePost(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Post",
	})
}
