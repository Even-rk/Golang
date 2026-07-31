package server

import (
	model "blog-backend/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 密码加密，使用 bcrypt 哈希算法
func encryptPassword(password string) string {
	// bcrypt 生成哈希密码
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// 如果加密失败，返回空字符串
		return ""
	}
	return string(hash)
}

// 密码验证，使用 bcrypt 哈希算法
func verifyPassword(hashedPassword, plainPassword string) bool {
	// bcrypt 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

// 注册用户
func RegisterUser(c *gin.Context, db *gorm.DB) {
	// 检查用户名是否存在
	username := c.Query("username")
	email := c.Query("email")
	password := c.Query("password")
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username already exists",
		})
		return
	}
	// 密码加密
	password = encryptPassword(password)
	if password == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "注册失败",
		})
		return
	}
	// 创建用户
	newUser := model.User{
		Username: username,
		Email:    email,
		Password: password,
	}
	// 保存用户到数据库
	err := db.Create(&newUser).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "用户注册失败" + err.Error(),
		})
		return
	}
	// 注册用户
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

// 用户登录
func LoginUser(c *gin.Context, db *gorm.DB) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login User",
	})
}
