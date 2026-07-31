package model

import "gorm.io/gorm"

// 文章模型
type Post struct {
	gorm.Model
	Title    string    `gorm:"column:title"`      // 文章标题
	Content  string    `gorm:"column:content"`    // 文章内容
	UserID   int       `gorm:"column:user_id"`    // 关联用户ID
	User     User      `gorm:"foreignKey:UserID"` // 一对一关联
	Comments []Comment `gorm:"foreignKey:PostID"` // 一对多关联
}

// 创建文章请求结构
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 更新文章请求结构
type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
