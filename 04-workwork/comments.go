package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetComments 获取指定文章的所有评论
func GetComments(c *gin.Context) {
	postIDStr := c.Param("postId")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Invalid post ID",
		})
		return
	}
	
	// 检查文章是否存在
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
	
	// 获取评论列表
	var comments []Comment
	query := db.Where("post_id = ?", postID).Preload("User")
	
	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset := (page - 1) * limit
	
	if err := query.Offset(offset).Limit(limit).Order("created_at desc").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to fetch comments",
		})
		return
	}
	
	// 获取评论总数
	var total int64
	db.Model(&Comment{}).Where("post_id = ?", postID).Count(&total)
	
	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Comments retrieved successfully",
		Data: gin.H{
			"comments": comments,
			"pagination": gin.H{
				"page":  page,
				"limit": limit,
				"total": total,
			},
		},
	})
}

// CreateComment 创建新评论
func CreateComment(c *gin.Context) {
	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "Invalid request data: " + err.Error(),
		})
		return
	}
	
	// 检查文章是否存在
	var post Post
	if err := db.First(&post, req.PostID).Error; err != nil {
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
	
	userID := getCurrentUserID(c)
	
	comment := Comment{
		Content: req.Content,
		UserID:  userID,
		PostID:  req.PostID,
	}
	
	if err := db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "Failed to create comment",
		})
		return
	}
	
	// 重新查询以获取用户信息
	db.Preload("User").First(&comment, comment.ID)
	
	c.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Message: "Comment created successfully",
		Data:    comment,
	})
}
