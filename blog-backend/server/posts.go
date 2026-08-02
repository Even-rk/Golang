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
	userClaims, _ = claims.(*middleware.Claims)
	if userClaims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "用户认证失败",
		})
		return
	}
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
	postID := c.Param("PostID")
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
	bindErr := c.ShouldBindJSON(&req)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数错误: " + bindErr.Error(),
		})
		return
	}

	// 从token中获取当前用户信息
	claims, _ := c.Get("claims")
	userClaims, _ := claims.(*middleware.Claims)
	if userClaims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "用户认证失败",
		})
		return
	}
	// 查询文章信息
	var post model.Post
	queryPostErr := db.Where("id = ?", req.PostID).First(&post).Error
	if queryPostErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": queryPostErr.Error(),
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
	savePostErr := db.Save(&post).Error
	if savePostErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": savePostErr.Error(),
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
	postID := c.Param("PostID")
	// 从token中获取当前用户信息
	claims, _ := c.Get("claims")
	userClaims, _ := claims.(*middleware.Claims)
	if userClaims == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "用户认证失败",
		})
		return
	}
	// 查询文章信息
	var post model.Post
	queryPostErr := db.Where("id = ?", postID).First(&post).Error
	if queryPostErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": queryPostErr.Error(),
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

	// 使用事务保证删除操作的原子性：先删除评论，再删除文章
	txErr := db.Transaction(func(tx *gorm.DB) error {
		// 删除文章下的所有评论（批量删除）
		delCommentsErr := tx.Where("post_id = ?", postID).Delete(&model.Comment{}).Error
		if delCommentsErr != nil {
			// 返回错误会自动回滚
			return delCommentsErr
		}
		// 删除文章
		delPostErr := tx.Delete(&post).Error
		if delPostErr != nil {
			// 返回错误会自动回滚
			return delPostErr
		}
		// 返回nil会自动提交
		return nil
	})

	if txErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": txErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "删除成功",
	})
}
