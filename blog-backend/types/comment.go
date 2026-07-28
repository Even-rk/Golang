package model

import "time"

type Comment struct {
	ID        int       `gorm:"primaryKey"`
	PostID    int       `gorm:"column:post_id"`
	Content   string    `gorm:"column:content"`
	Post      Post      `gorm:"foreignKey:PostID"` // 一对一关联
	CreatedAt time.Time `gorm:"autoCreateTime"`    // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"`    // 更新时间
}
