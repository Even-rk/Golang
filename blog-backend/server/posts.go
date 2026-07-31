package server

import (
	"blog-backend/middleware"
	model "blog-backend/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建文章
func CreatePost(c *gin.Context, db *gorm.DB) {
	// 绑定请求参数到 CreatePostRequest 模型
	var req model.CreatePostRequest
	reqErr := c.ShouldBindJSON(&req)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误: " + reqErr.Error(),
		})
		return
	}
	// 从token中取用户信息
	var userClaims *middleware.Claims
	claims, _ := c.Get("claims")
	// 类型断言转换为 *middleware.Claims
	userClaims, _ = claims.(*middleware.Claims)
	// 创建文章
	newPost := model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userClaims.UserID,
	}
	// 保存文章到数据库
	createErr := db.Create(&newPost).Error
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "文章创建失败" + createErr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "文章创建成功",
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
