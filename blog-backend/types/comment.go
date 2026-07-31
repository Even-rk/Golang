package model

import "time"

type Comment struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"column:user_id"`    // 关联用户ID
	PostID    int       `gorm:"column:post_id"`    // 关联文章ID
	Content   string    `gorm:"column:content"`    // 评论内容
	User      User      `gorm:"foreignKey:UserID"` // 关联用户
	Post      Post      `gorm:"foreignKey:PostID"` // 关联文章
	CreatedAt time.Time `gorm:"autoCreateTime"`    // 创建时间
}
