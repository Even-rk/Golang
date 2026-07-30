package middleware

import (
	"github.com/gin-gonic/gin"
)

// JWT 中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization 字段
		authHeader := c.GetHeader("Authorization")
		// 检查 Authorization 字段是否存在
		if authHeader == "" {
			// 如果Authorization字段为空，返回401 Unauthorized错误
			c.JSON(401, gin.H{"error": "Authorization header is missing"})
			// 中止请求处理
			c.Abort()
			return
		}
		// 如果Authorization字段存在，继续处理请求，调用下一个中间件或处理函数
		c.Next()
	}
}
