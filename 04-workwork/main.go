package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func main() {
	// 连接数据库
	var err error
	// GORM数据库连接
	db, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移模型
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 创建Gin路由
	r := gin.Default()

	// 添加中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 设置路由组
	api := r.Group("/api")
	{
		// 用户认证相关路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", Register)
			auth.POST("/login", Login)
		}

		// 文章相关路由
		posts := api.Group("/posts")
		{
			posts.GET("", GetPosts)                            // 获取所有文章
			posts.GET("/:id", GetPost)                         // 获取单个文章
			posts.POST("", AuthMiddleware(), CreatePost)       // 创建文章
			posts.PUT("/:id", AuthMiddleware(), UpdatePost)    // 更新文章
			posts.DELETE("/:id", AuthMiddleware(), DeletePost) // 删除文章
		}

		// 评论相关路由
		comments := api.Group("/comments")
		{
			comments.GET("/post/:postId", GetComments)         // 获取文章评论
			comments.POST("", AuthMiddleware(), CreateComment) // 创建评论
		}
	}

	// 启动服务器
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
