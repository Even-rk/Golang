package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // 自动添加主键、创建时间、更新时间字段
	Username   string `gorm:"not null"` // 用户名
	Email      string `gorm:"not null"` // 邮箱
	Password   string `gorm:"not null"` // 密码
}
