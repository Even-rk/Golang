package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 可嵌入标准字段的Claims 结构体
type CustomClaims struct {
	UserID               string `json:"user_id"`
	Password             string `json:"password"`
	jwt.RegisteredClaims        // 包含 exp, iat, nbf 等标准字段
}

// JWT 签名密钥（生产环境应放在配置中）
var jwtKey = []byte("my_secret_key")

// 自定义 Claims 结构体，嵌套标准 RegisteredClaims
type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

// GenerateToken 为指定用户名生成 JWT 令牌
func GenerateToken(username string, email string, password string) (string, error) {
	// 设置令牌过期时间为 24 小时后
	expirationTime := time.Now().Add(24 * time.Hour)
	// 构造包含用户信息的 claims
	claims := &Claims{
		Username: username,
		Email:    email,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			// 将过期时间转换为 NumericDate 格式
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// 使用 HS256 算法和 claims 创建新令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥对令牌进行签名并返回字符串
	return token.SignedString(jwtKey)
}

// AuthMiddleware 返回一个用于 JWT 认证的 Gin 中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization 字段
		tokenString := c.GetHeader("Authorization")
		// 检查令牌是否存在且以 Bearer 开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			// 缺失或格式错误，中断请求并返回 401
			c.AbortWithStatusJSON(401, gin.H{"error": "missing token"})
			return
		}
		// 去掉 Bearer 前缀，保留实际令牌部分
		tokenString = tokenString[7:]
		// 创建空 Claims 用于解析
		claims := &Claims{}
		// 解析并验证令牌，同时将声明填充到 claims
		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {
				// 提供解析所需的签名密钥
				return jwtKey, nil
			})
		// 如果解析出错或令牌无效，返回 401
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}
		// 令牌有效，将 claims 存入上下文供后续处理器使用
		c.Set("claims", claims)
		// 继续执行后续处理器
		c.Next()
	}
}
