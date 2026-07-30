package middleware

import "github.com/gin-gonic/gin"

// JWT 中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
