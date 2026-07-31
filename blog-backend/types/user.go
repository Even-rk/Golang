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

// 注册用户
type RegisterUser struct {
	Username string `json:"username" binding:"required"`              // 用户名不能为空
	Email    string `json:"email" binding:"required,email"`           // 邮箱不能为空且格式正确
	Password string `json:"password" binding:"required,min=6,max=20"` // 密码不能为空且长度在6到20之间
}

// 登录用户
type LoginUser struct {
	Username string `json:"username" binding:"required"` // 用户名不能为空
	Email    string `json:"email"`                       // 邮箱
	Password string `json:"password" binding:"required"` // 密码不能为空
}
