package main

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"` // 密码不返回给前端
	Email    string `gorm:"unique;not null" json:"email"`
	Posts    []Post `json:"posts,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

// Post 文章模型
type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user,omitempty"`
	PostID  uint   `json:"post_id"`
	Post    Post   `json:"post,omitempty"`
}

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest 注册请求结构
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// CreatePostRequest 创建文章请求结构
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// UpdatePostRequest 更新文章请求结构
type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// CreateCommentRequest 创建评论请求结构
type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
}

// JWTClaims JWT声明结构
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
}

// APIResponse API响应结构
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}
