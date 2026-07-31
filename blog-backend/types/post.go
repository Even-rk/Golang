package model

import "time"

// 创建post模型 ：文章模型，包含文章ID、标题、内容、用户ID，评论状态字段
type Post struct {
	ID        int       `gorm:"primaryKey"`
	Title     string    `gorm:"column:title"`      // 文章标题
	Content   string    `gorm:"column:content"`    // 文章内容
	UserID    int       `gorm:"column:user_id"`    // 关联用户ID
	User      User      `gorm:"foreignKey:UserID"` // 一对一关联
	Comments  []Comment `gorm:"foreignKey:PostID"` // 一对多关联
	CreatedAt time.Time `gorm:"autoCreateTime"`    // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"`    // 更新时间
}
