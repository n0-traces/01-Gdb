package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPosts 获取所有文章列表
func GetPosts(c *gin.Context) {
	var posts []Post
	
	// 预加载用户信息和评论数量
	query := db.Preload("User").Preload("Comments")
	
	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit
	
	// 执行查询
	if err := query.Offset(offset).Limit(limit).Order("created_at desc").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to fetch posts",
		})
		return
	}
	
	// 获取总数
	var total int64
	db.Model(&Post{}).Count(&total)
	
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Posts retrieved successfully",
		Data: gin.H{
			"posts": posts,
			"pagination": gin.H{
				"page":  page,
				"limit": limit,
				"total": total,
			},
		},
	})
}

// GetPost 获取单个文章详情
func GetPost(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Invalid post ID",
		})
		return
	}
	
	var post Post
	if err := db.Preload("User").Preload("Comments.User").First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, APIResponse{
				Success: false,
				Error:   "Post not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, APIResponse{
				Success: false,
				Error:   "Failed to fetch post",
			})
		}
		return
	}
	
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Post retrieved successfully",
		Data:    post,
	})
}

// CreatePost 创建新文章
func CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Invalid request data: " + err.Error(),
		})
		return
	}
	
	userID := getCurrentUserID(c)
	
	post := Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}
	
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to create post",
		})
		return
	}
	
	// 重新查询以获取用户信息
	db.Preload("User").First(&post, post.ID)
	
	c.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Message: "Post created successfully",
		Data:    post,
	})
}

// UpdatePost 更新文章
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Invalid post ID",
		})
		return
	}
	
	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Invalid request data: " + err.Error(),
		})
		return
	}
	
	// 查找文章
	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, APIResponse{
				Success: false,
				Error:   "Post not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, APIResponse{
				Success: false,
				Error:   "Failed to fetch post",
			})
		}
		return
	}
	
	// 检查权限：只有作者才能更新文章
	userID := getCurrentUserID(c)
	if post.UserID != userID {
		c.JSON(http.StatusForbidden, APIResponse{
			Success: false,
			Error:   "You can only update your own posts",
		})
		return
	}
	
	// 更新文章
	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}
	
	if err := db.Model(&post).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to update post",
		})
		return
	}
	
	// 重新查询以获取完整信息
	db.Preload("User").First(&post, post.ID)
	
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Post updated successfully",
		Data:    post,
	})
}

// DeletePost 删除文章
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Invalid post ID",
		})
		return
	}
	
	// 查找文章
	var post Post
	if err := db.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, APIResponse{
				Success: false,
				Error:   "Post not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, APIResponse{
				Success: false,
				Error:   "Failed to fetch post",
			})
		}
		return
	}
	
	// 检查权限：只有作者才能删除文章
	userID := getCurrentUserID(c)
	if post.UserID != userID {
		c.JSON(http.StatusForbidden, APIResponse{
			Success: false,
			Error:   "You can only delete your own posts",
		})
		return
	}
	
	// 删除文章（会级联删除相关评论）
	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to delete post",
		})
		return
	}
	
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Post deleted successfully",
	})
}
