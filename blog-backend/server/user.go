package server

import (
	"blog-backend/middleware"
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
	// 绑定请求参数到 RegisterUser 模型
	var req model.RegisterUser
	// 验证请求参数
	reqErr := c.ShouldBindJSON(&req)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误: " + reqErr.Error(),
		})
		return
	}
	// 检查用户名是否存在
	var user model.User
	queryUserErr := db.Where("username = ?", req.Username).First(&user).Error
	// 如果用户名不存在，返回错误
	if queryUserErr == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户名已存在",
		})
		return
	}
	// 检查邮箱是否存在
	var userByEmail model.User
	queryEmailErr := db.Where("email = ?", req.Email).First(&userByEmail).Error
	if queryEmailErr == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "邮箱已存在",
		})
		return
	}
	// 密码加密
	encryptedPassword := encryptPassword(req.Password)
	if encryptedPassword == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "密码加密失败",
		})
		return
	}
	// 创建用户
	newUser := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: encryptedPassword,
	}
	// 保存用户到数据库
	createErr := db.Create(&newUser).Error
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "用户注册失败" + createErr.Error(),
		})
		return
	}
	// 注册用户
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "注册成功",
	})
}

// 用户登录
func LoginUser(c *gin.Context, db *gorm.DB) {
	// 绑定请求参数到 LoginUser 模型
	var req model.LoginUser
	// 验证请求参数
	reqErr := c.ShouldBindJSON(&req)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数错误: " + reqErr.Error(),
		})
		return
	}

	// 根据用户名查询用户
	var user model.User
	userErr := db.Where("username = ?", req.Username).First(&user).Error
	if userErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "用户不存在",
		})
		return
	}
	// 验证密码
	passwordValid := verifyPassword(user.Password, req.Password)
	if !passwordValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "密码错误",
		})
		return
	}

	// 登录成功，返回token
	token, _ := middleware.GenerateToken(int(user.ID), user.Username, user.Email)
	c.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"message":     "登录成功",
		"accessToken": token,
	})
}
