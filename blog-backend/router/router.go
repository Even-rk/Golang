package router

import (
	"blog-backend/middleware"
	"blog-backend/server"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// 配置CORS中间件
	r.Use(middleware.CORS())

	// 配置JWT中间件
	r.Use(middleware.JWT())

	// 根路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// 配置用户路由
	userRouter := r.Group("/user")
	{
		// 注册用户
		userRouter.POST("/register", func(c *gin.Context) {
			server.RegisterUser(c, db)
		})
		// 用户登录
		userRouter.POST("/login", func(c *gin.Context) {
			server.LoginUser(c, db)
		})
	}

	// 配置文章路由
	postRouter := r.Group("/post")
	{
		// 创建文章
		postRouter.POST("/create", func(c *gin.Context) {
			server.CreatePost(c, db)
		})
		// 获取单个文章列表
		postRouter.GET("/list", func(c *gin.Context) {
			server.GetSinglePostList(c, db)
		})
		// 获取所有文章列表
		postRouter.GET("/all", func(c *gin.Context) {
			server.GetAllPostList(c, db)
		})
		// 更新文章
		postRouter.PUT("/update", func(c *gin.Context) {
			server.UpdatePost(c, db)
		})
		// 删除文章
		postRouter.DELETE("/delete", func(c *gin.Context) {
			server.DeletePost(c, db)
		})
	}

	// 配置评论路由
	commentRouter := r.Group("/comment")
	{
		// 创建评论
		commentRouter.POST("/create", func(c *gin.Context) {
			server.CreateComment(c, db)
		})
		// 获取评论列表
		commentRouter.GET("/list", func(c *gin.Context) {
			server.GetCommentList(c, db)
		})
	}
}
