package server

import (
	"blog-backend/middleware"
	model "blog-backend/types"
	"net/http"
	"strconv"

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
	var post model.Post
	err := db.Preload("Comments.User").Preload("User").Where("id = ?", postID).First(&post).Error
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    200,
			"message": "success",
			"data":    []gin.H{},
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	// 转换文章为JSON格式，保持和列表接口一致的字段
	postJSON := gin.H{
		"id":         post.ID,
		"title":      post.Title,
		"content":    post.Content,
		"user_id":    post.UserID,
		"user_name":  post.User.Username,
		"created_at": post.CreatedAt.Format("2006-01-02 15:04:05"),
		"comments":   post.Comments,
	}

	// 返回文章详情
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    postJSON,
	})
}

// 获取所有文章列表
func GetAllPostList(c *gin.Context, db *gorm.DB) {
	var posts []model.Post
	// 使用 Preload 预加载评论和用户信息
	err := db.Preload("Comments.User").Preload("User").Find(&posts).Error
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    200,
			"message": "success",
			"data":    []gin.H{},
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	// 转换文章列表为JSON格式
	var postsData []gin.H
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
		postsData = append(postsData, postJSON)
	}

	// 返回文章列表
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    postsData,
		"total":   len(postsData), // 添加总数
	})
}

// 更新文章
func UpdatePost(c *gin.Context, db *gorm.DB) {
	// 绑定更新请求参数
	var req model.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数错误: " + err.Error(),
		})
		return
	}

	// 从token中获取当前用户信息
	claims, _ := c.Get("claims")
	userClaims, _ := claims.(*middleware.Claims)
	// 查询文章信息
	var post model.Post
	if err := db.Where("id = ?", req.PostID).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	// 权限验证：只有作者本人可以更新文章
	if post.UserID != userClaims.UserID {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "无权限，只有作者本人可以修改文章",
		})
		return
	}

	// 更新文章字段
	if req.Title != "" {
		post.Title = req.Title
	}
	if req.Content != "" {
		post.Content = req.Content
	}

	// 保存更新
	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "更新成功",
	})
}

// 删除文章
func DeletePost(c *gin.Context, db *gorm.DB) {
	// 从请求参数中获取文章ID
	postIDStr := c.Param("postID")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "文章ID格式错误",
		})
		return
	}

	// 从token中获取当前用户信息
	claims, _ := c.Get("claims")
	userClaims, _ := claims.(*middleware.Claims)

	// 查询文章信息
	var post model.Post
	if err := db.Where("id = ?", postID).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	// 权限验证：只有作者本人可以删除文章
	if post.UserID != userClaims.UserID {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "无权限，只有作者本人可以删除文章",
		})
		return
	}

	// 删除文章（GORM 的 Delete 会软删除，因为 gorm.Model 包含 DeletedAt 字段）
	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "删除成功",
	})
}
