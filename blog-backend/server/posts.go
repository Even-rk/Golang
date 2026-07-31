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

// 获取单个文章详情
func GetSinglePost(c *gin.Context, db *gorm.DB) {
	// 从请求参数中获取文章ID
	postID := c.Param("postID")
	// 从数据库中查询文章
	var post model.Post
	if err := db.Where("id = ?", postID).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}
	// 查询文章评论
	if err := db.Where("post_id = ?", postID).Find(&post.Comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	// 转换文章为JSON格式
	postJSON := gin.H{
		"id":       post.ID,
		"title":    post.Title,
		"content":  post.Content,
		"user_id":  post.UserID,
		"comments": post.Comments,
	}

	// 返回文章详情
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"post":    postJSON,
	})
}

// 获取所有文章列表
func GetAllPostList(c *gin.Context, db *gorm.DB) {
	var posts []model.Post
	// 使用 Preload 预加载评论和用户信息
	err := db.Preload("Comments.User").Preload("User").Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	// 转换文章列表为JSON格式
	var postsJSON []gin.H
	for _, post := range posts {
		postJSON := gin.H{
			"id":         post.ID,
			"title":      post.Title,
			"content":    post.Content,
			"user_id":    post.UserID,
			"user_name":  post.User.Username,
			"created_at": post.CreatedAt.Format("2006-01-02 15:04:05"), // 添加创建时间
			"comments":   post.Comments,
		}
		postsJSON = append(postsJSON, postJSON)
	}

	// 返回文章列表
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"posts":   postsJSON,
		"total":   len(postsJSON), // 添加总数
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
