package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`     // 主键
	Name      string    `gorm:"not null"`       // 姓名
	Email     string    `gorm:"not null"`       // 邮箱
	Password  string    `gorm:"not null"`       // 密码
	CreatedAt time.Time `gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // 更新时间
}
