package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] | %s | %s | %s | %d | %s | %s | %s\n", // 按指定格式生成日志字符串
			param.ClientIP,                       // 客户端 IP 地址
			param.TimeStamp.Format(time.RFC1123), // 时间戳，格式化为 RFC1123 标准
			param.Method,                         // HTTP 请求方法
			param.Path,                           // 请求路径
			param.Request.Proto,                  // HTTP 协议版本
			param.StatusCode,                     // 响应状态码
			param.Latency,                        // 请求处理延迟
			param.Request.UserAgent(),            // 请求头中的 User-Agent
			param.ErrorMessage,                   // 错误信息（如果有）
		)
	})
}
