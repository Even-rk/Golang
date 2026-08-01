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

	// 配置文章路由
	postRouter := r.Group("/post")
	postRouter.Use(middleware.JWT())
	{
		// 创建文章
		postRouter.POST("/create", func(c *gin.Context) {
			server.CreatePost(c, db)
		})
		// 获取单个文章详情
		postRouter.GET("/:PostID", func(c *gin.Context) {
			server.GetSinglePost(c, db)
		})
		// 获取所有文章列表
		postRouter.GET("/allList", func(c *gin.Context) {
			server.GetAllPostList(c, db)
		})

		// 更新文章
		postRouter.PUT("/update", func(c *gin.Context) {
			server.UpdatePost(c, db)
		})
		// 删除文章
		postRouter.DELETE("/delete/:PostID", func(c *gin.Context) {
			server.DeletePost(c, db)
		})
	}

	// 配置评论路由
	commentRouter := r.Group("/comment")
	commentRouter.Use(middleware.JWT())
	{
		// 创建评论
		commentRouter.POST("/create", func(c *gin.Context) {
			server.CreateComment(c, db)
		})
		// 获取评论列表
		commentRouter.GET("/list/:postID", func(c *gin.Context) {
			server.GetCommentListByPostID(c, db)
		})
		// 删除评论
		commentRouter.DELETE("/delete/:commentID", func(c *gin.Context) {
			server.DeleteComment(c, db)
		})
	}
}
