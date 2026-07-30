package router

import (
	"blog-backend/middleware"
	"blog-backend/server"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// 配置CORS中间件
	r.Use(middleware.CORS())

	// 配置JWT中间件
	r.Use(middleware.JWT())

	// 配置用户路由
	userRouter := r.Group("/user")
	{
		// 注册用户
		userRouter.POST("/register", server.RegisterUser)
		// 用户登录
		userRouter.POST("/login", server.LoginUser)
	}

	// 配置文章路由
	postRouter := r.Group("/post")
	{
		// 创建文章
		postRouter.POST("/create", server.CreatePost)
		// 获取文章列表
		postRouter.GET("/list", server.GetPostList)
		// 获取文章详情
		postRouter.GET("/detail", server.GetPostDetail)
	}

	// 配置评论路由
	commentRouter := r.Group("/comment")
	{
		// 创建评论
		commentRouter.POST("/create", server.CreateComment)
		// 获取评论列表
		commentRouter.GET("/list", server.GetCommentList)
		// 获取评论详情
		commentRouter.GET("/detail", server.GetCommentDetail)
	}
}
