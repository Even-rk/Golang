package router

import (
	"blog-backend/middleware"
	"blog-backend/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// 配置CORS中间件
	r.Use(middleware.CORS())
	// 配置日志中间件
	r.Use(middleware.Logger())
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

	// 配置公开文章路由（不需要认证）
	postPublicRouter := r.Group("/post")
	{
		// 获取单个文章详情
		postPublicRouter.GET("/:PostID", func(c *gin.Context) {
			server.GetSinglePost(c, db)
		})
		// 获取所有文章列表
		postPublicRouter.GET("/allList", func(c *gin.Context) {
			server.GetAllPostList(c, db)
		})
	}

	// 配置需要认证的文章路由
	postAuthRouter := r.Group("/post")
	postAuthRouter.Use(middleware.JWT())
	{
		// 创建文章
		postAuthRouter.POST("/create", func(c *gin.Context) {
			server.CreatePost(c, db)
		})
		// 更新文章
		postAuthRouter.PUT("/update", func(c *gin.Context) {
			server.UpdatePost(c, db)
		})
		// 删除文章
		postAuthRouter.DELETE("/delete/:PostID", func(c *gin.Context) {
			server.DeletePost(c, db)
		})
	}

	// 配置公开评论路由（不需要认证）
	commentPublicRouter := r.Group("/comment")
	{
		// 获取评论列表
		commentPublicRouter.GET("/list/:PostID", func(c *gin.Context) {
			server.GetCommentListByPostID(c, db)
		})
	}

	// 配置需要认证的评论路由
	commentAuthRouter := r.Group("/comment")
	commentAuthRouter.Use(middleware.JWT())
	{
		// 创建评论
		commentAuthRouter.POST("/create", func(c *gin.Context) {
			server.CreateComment(c, db)
		})
		// 删除评论
		commentAuthRouter.DELETE("/delete/:CommentID", func(c *gin.Context) {
			server.DeleteComment(c, db)
		})
	}
}
