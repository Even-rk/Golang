package server

import (
	"blog-backend/middleware"
	model "blog-backend/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 创建评论
func CreateComment(c *gin.Context, db *gorm.DB) {
	// 绑定请求参数到 CreateCommentRequest 模型
	var req model.CreateCommentRequest
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
	// 创建评论
	newComment := model.Comment{
		PostID:  req.PostID,
		Content: req.Content,
		UserID:  userClaims.UserID,
	}
	// 保存评论到数据库
	createErr := db.Create(&newComment).Error
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "评论创建失败" + createErr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "评论创建成功",
	})
}

// 获取文章下评论列表
func GetCommentListByPostID(c *gin.Context, db *gorm.DB) {
	// 从请求参数中获取文章ID
	postID := c.Param("postID")
	var comments []model.Comment
	// 使用 Preload 预加载用户信息
	err := db.Preload("User").Where("post_id = ?", postID).Find(&comments).Error
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

	// 转换评论列表为JSON格式
	var commentsData []gin.H
	for _, comment := range comments {
		commentJSON := gin.H{
			"id":         comment.ID,
			"post_id":    comment.PostID,
			"user_id":    comment.UserID,
			"user_name":  comment.User.Username,
			"content":    comment.Content,
			"created_at": comment.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		commentsData = append(commentsData, commentJSON)
	}

	// 返回评论列表
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    commentsData,
		"total":   len(commentsData),
	})
}

// 删除评论
func DeleteComment(c *gin.Context, db *gorm.DB) {
	// 从请求参数中获取评论ID
	commentID := c.Param("commentID")
	// 从token中获取当前用户信息
	claims, _ := c.Get("claims")
	userClaims, _ := claims.(*middleware.Claims)

	// 查询评论信息
	var comment model.Comment
	if err := db.Where("id = ?", commentID).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	// 权限验证：只有评论作者本人可以删除评论
	if comment.UserID != userClaims.UserID {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "无权限，只有评论作者本人可以删除评论",
		})
		return
	}

	// 删除评论（GORM 的 Delete 会软删除，因为 gorm.Model 包含 DeletedAt 字段）
	if err := db.Delete(&comment).Error; err != nil {
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
