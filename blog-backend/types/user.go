package model

import (
	"gorm.io/gorm"
)

// 用户模型
type User struct {
	gorm.Model        // 自动添加主键、创建时间、更新时间字段
	Username   string `gorm:"column:username"` // 用户名
	Email      string `gorm:"column:email"`    // 邮箱
	Password   string `gorm:"column:password"` // 密码
}
