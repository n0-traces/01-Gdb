package gaa

import (
	"fmt"
	"time"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Username  string    `gorm:"size:50;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null" json:"-"`
	Nickname  string    `gorm:"size:50" json:"nickname"`
	Avatar    string    `gorm:"size:200" json:"avatar"`
	PostCount int       `gorm:"default:0" json:"post_count"` // 文章数量统计字段
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// 关联关系
	Posts []Post `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"posts,omitempty"`
}

// Post 文章模型
type Post struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Title       string    `gorm:"size:200;not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	Summary     string    `gorm:"size:500" json:"summary"`
	Status      string    `gorm:"size:20;default:'published'" json:"status"` // published, draft
	ViewCount   int       `gorm:"default:0" json:"view_count"`
	CommentCount int      `gorm:"default:0" json:"comment_count"` // 评论数量
	CommentStatus string  `gorm:"size:20;default:'有评论'" json:"comment_status"` // 有评论, 无评论
	UserID      uint      `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	
	// 关联关系
	User     User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"comments,omitempty"`
}

// Comment 评论模型
type Comment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	ParentID  *uint     `gorm:"default:null" json:"parent_id"` // 父评论ID，支持回复功能
	Status    string    `gorm:"size:20;default:'approved'" json:"status"` // approved, pending, rejected
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// 关联关系
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Post Post `gorm:"foreignKey:PostID" json:"post,omitempty"`
}

// ==================== 钩子函数 ====================

// BeforeCreate Post创建前的钩子函数
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	fmt.Printf("🔄 正在创建文章: %s\n", p.Title)
	return nil
}

// AfterCreate Post创建后的钩子函数 - 更新用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 更新用户的文章数量
	err := tx.Model(&User{}).Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + 1")).Error
	if err != nil {
		return fmt.Errorf("更新用户文章数量失败: %v", err)
	}
	fmt.Printf("✅ 用户文章数量已更新，文章: %s\n", p.Title)
	return nil
}

// BeforeDelete Post删除前的钩子函数
func (p *Post) BeforeDelete(tx *gorm.DB) error {
	fmt.Printf("🔄 正在删除文章: %s\n", p.Title)
	return nil
}

// AfterDelete Post删除后的钩子函数 - 减少用户的文章数量
func (p *Post) AfterDelete(tx *gorm.DB) error {
	// 减少用户的文章数量
	err := tx.Model(&User{}).Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count - 1")).Error
	if err != nil {
		return fmt.Errorf("减少用户文章数量失败: %v", err)
	}
	fmt.Printf("✅ 用户文章数量已减少，文章: %s\n", p.Title)
	return nil
}

// BeforeCreate Comment创建前的钩子函数
func (c *Comment) BeforeCreate(tx *gorm.DB) error {
	fmt.Printf("🔄 正在创建评论，文章ID: %d\n", c.PostID)
	return nil
}

// AfterCreate Comment创建后的钩子函数 - 更新文章的评论数量
func (c *Comment) AfterCreate(tx *gorm.DB) error {
	// 更新文章的评论数量
	err := tx.Model(&Post{}).Where("id = ?", c.PostID).
		UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error
	if err != nil {
		return fmt.Errorf("更新文章评论数量失败: %v", err)
	}
	
	// 更新文章的评论状态
	err = tx.Model(&Post{}).Where("id = ?", c.PostID).
		Update("comment_status", "有评论").Error
	if err != nil {
		return fmt.Errorf("更新文章评论状态失败: %v", err)
	}
	
	fmt.Printf("✅ 文章评论数量已更新，评论ID: %d\n", c.ID)
	return nil
}

// BeforeDelete Comment删除前的钩子函数
func (c *Comment) BeforeDelete(tx *gorm.DB) error {
	fmt.Printf("🔄 正在删除评论，评论ID: %d\n", c.ID)
	return nil
}

// AfterDelete Comment删除后的钩子函数 - 检查并更新文章的评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	// 减少文章的评论数量
	err := tx.Model(&Post{}).Where("id = ?", c.PostID).
		UpdateColumn("comment_count", gorm.Expr("comment_count - 1")).Error
	if err != nil {
		return fmt.Errorf("减少文章评论数量失败: %v", err)
	}
	
	// 检查文章的评论数量，如果为0则更新状态
	var commentCount int64
	err = tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error
	if err != nil {
		return fmt.Errorf("查询文章评论数量失败: %v", err)
	}
	
	if commentCount == 0 {
		err = tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
		if err != nil {
			return fmt.Errorf("更新文章评论状态失败: %v", err)
		}
		fmt.Printf("✅ 文章评论数量为0，状态已更新为'无评论'，文章ID: %d\n", c.PostID)
	} else {
		fmt.Printf("✅ 文章评论数量已减少，当前剩余: %d，文章ID: %d\n", commentCount, c.PostID)
	}
	
	return nil
}

// ==================== 数据库操作函数 ====================

// CreateBlogTables 创建博客系统相关的数据库表
func CreateBlogTables(db *gorm.DB) error {
	fmt.Println("🏗️ 开始创建博客系统数据库表...")
	
	// 自动迁移创建表
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		return fmt.Errorf("创建数据库表失败: %v", err)
	}
	
	fmt.Println("✅ 博客系统数据库表创建成功!")
	return nil
}

// InsertTestData 插入测试数据
func InsertTestData(db *gorm.DB) error {
	fmt.Println("📝 开始插入测试数据...")
	
	// 创建测试用户
	users := []User{
		{Username: "zhangsan", Email: "zhangsan@example.com", Password: "password123", Nickname: "张三"},
		{Username: "lisi", Email: "lisi@example.com", Password: "password123", Nickname: "李四"},
		{Username: "wangwu", Email: "wangwu@example.com", Password: "password123", Nickname: "王五"},
	}
	
	for i := range users {
		err := db.Create(&users[i]).Error
		if err != nil {
			return fmt.Errorf("创建用户失败: %v", err)
		}
		fmt.Printf("✅ 创建用户: %s (%s)\n", users[i].Nickname, users[i].Username)
	}
	
	// 创建测试文章
	posts := []Post{
		{Title: "Go语言入门指南", Content: "Go语言是一门简洁高效的编程语言...", Summary: "Go语言基础教程", UserID: users[0].ID},
		{Title: "GORM使用详解", Content: "GORM是Go语言中最流行的ORM框架...", Summary: "GORM框架教程", UserID: users[0].ID},
		{Title: "Web开发最佳实践", Content: "在Web开发中，我们需要遵循一些最佳实践...", Summary: "Web开发指南", UserID: users[1].ID},
		{Title: "数据库设计原则", Content: "良好的数据库设计是系统成功的关键...", Summary: "数据库设计教程", UserID: users[1].ID},
		{Title: "微服务架构", Content: "微服务架构是现代应用开发的主流选择...", Summary: "微服务教程", UserID: users[2].ID},
	}
	
	for i := range posts {
		err := db.Create(&posts[i]).Error
		if err != nil {
			return fmt.Errorf("创建文章失败: %v", err)
		}
		fmt.Printf("✅ 创建文章: %s (作者: %s)\n", posts[i].Title, users[posts[i].UserID-1].Nickname)
	}
	
	// 创建测试评论
	comments := []Comment{
		{Content: "这篇文章写得很好，对我帮助很大！", UserID: users[1].ID, PostID: posts[0].ID},
		{Content: "感谢分享，学到了很多知识", UserID: users[2].ID, PostID: posts[0].ID},
		{Content: "GORM确实很好用，感谢推荐", UserID: users[1].ID, PostID: posts[1].ID},
		{Content: "Web开发确实需要遵循这些原则", UserID: users[2].ID, PostID: posts[2].ID},
		{Content: "数据库设计很重要，这篇文章很有价值", UserID: users[0].ID, PostID: posts[3].ID},
		{Content: "微服务架构确实很复杂，需要深入学习", UserID: users[1].ID, PostID: posts[4].ID},
		{Content: "期待更多关于微服务的内容", UserID: users[2].ID, PostID: posts[4].ID},
		{Content: "这篇文章的评论数量很多啊", UserID: users[0].ID, PostID: posts[4].ID},
	}
	
	for i := range comments {
		err := db.Create(&comments[i]).Error
		if err != nil {
			return fmt.Errorf("创建评论失败: %v", err)
		}
		// 安全地截取评论内容，避免数组越界
		commentPreview := comments[i].Content
		if len(commentPreview) > 20 {
			commentPreview = commentPreview[:20] + "..."
		}
		fmt.Printf("✅ 创建评论: %s (用户: %s)\n", commentPreview, users[comments[i].UserID-1].Nickname)
	}
	
	fmt.Println("✅ 测试数据插入完成!")
	return nil
}

// ==================== 关联查询函数 ====================

// QueryUserPostsWithComments 查询某个用户发布的所有文章及其对应的评论信息
func QueryUserPostsWithComments(db *gorm.DB, userID uint) (User, error) {
	var user User
	
	err := db.Preload("Posts.Comments.User").
		Preload("Posts.User").
		Where("id = ?", userID).
		First(&user).Error
	
	if err != nil {
		return User{}, fmt.Errorf("查询用户文章和评论失败: %v", err)
	}
	
	return user, nil
}

// QueryPostWithMostComments 查询评论数量最多的文章信息
func QueryPostWithMostComments(db *gorm.DB) (Post, error) {
	var post Post
	
	err := db.Preload("User").
		Preload("Comments.User").
		Order("comment_count DESC").
		First(&post).Error
	
	if err != nil {
		return Post{}, fmt.Errorf("查询评论最多的文章失败: %v", err)
	}
	
	return post, nil
}

// ==================== 演示函数 ====================

// RunBlogDemo 运行博客系统演示
func RunBlogDemo(db *gorm.DB) {
	fmt.Println("📝 博客系统GORM演示")
	fmt.Println("==================================================")
	
	// 1. 创建数据库表
	fmt.Println("\n1️⃣ 创建博客系统数据库表")
	err := CreateBlogTables(db)
	if err != nil {
		panic(err)
	}
	
	// 2. 插入测试数据
	fmt.Println("\n2️⃣ 插入测试数据")
	err = InsertTestData(db)
	if err != nil {
		panic(err)
	}
	
	// 3. 关联查询演示
	fmt.Println("\n3️⃣ 关联查询演示")
	
	// 查询用户张三的所有文章及评论
	fmt.Println("\n🔍 查询用户'张三'的所有文章及评论:")
	zhangsan, err := QueryUserPostsWithComments(db, 1)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("用户: %s (%s)\n", zhangsan.Nickname, zhangsan.Username)
	fmt.Printf("文章数量: %d\n", len(zhangsan.Posts))
	
	for i, post := range zhangsan.Posts {
		fmt.Printf("\n  %d. 文章: %s\n", i+1, post.Title)
		// 安全地截取内容，避免数组越界
		contentPreview := post.Content
		if len(contentPreview) > 50 {
			contentPreview = contentPreview[:50] + "..."
		}
		fmt.Printf("     内容: %s\n", contentPreview)
		fmt.Printf("     评论数量: %d\n", len(post.Comments))
		
		for j, comment := range post.Comments {
			fmt.Printf("       %d. %s: %s\n", j+1, comment.User.Nickname, comment.Content)
		}
	}
	
	// 查询评论最多的文章
	fmt.Println("\n🔍 查询评论数量最多的文章:")
	mostCommentedPost, err := QueryPostWithMostComments(db)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("文章标题: %s\n", mostCommentedPost.Title)
	fmt.Printf("作者: %s\n", mostCommentedPost.User.Nickname)
	fmt.Printf("评论数量: %d\n", len(mostCommentedPost.Comments))
	fmt.Printf("评论状态: %s\n", mostCommentedPost.CommentStatus)
	
	for i, comment := range mostCommentedPost.Comments {
		fmt.Printf("  %d. %s: %s\n", i+1, comment.User.Nickname, comment.Content)
	}
	
	// 4. 钩子函数演示
	fmt.Println("\n4️⃣ 钩子函数演示")
	
	// 创建新文章（触发AfterCreate钩子）
	fmt.Println("\n📝 创建新文章（测试钩子函数）:")
	newPost := Post{
		Title:   "钩子函数测试文章",
		Content: "这是一篇用于测试钩子函数的文章...",
		Summary: "钩子函数测试",
		UserID:  1,
	}
	
	err = db.Create(&newPost).Error
	if err != nil {
		panic(err)
	}
	
	// 创建新评论（触发AfterCreate钩子）
	fmt.Println("\n💬 创建新评论（测试钩子函数）:")
	newComment := Comment{
		Content: "这是一条测试评论，用于验证钩子函数",
		UserID:  2,
		PostID:  newPost.ID,
	}
	
	err = db.Create(&newComment).Error
	if err != nil {
		panic(err)
	}
	
	// 删除评论（触发AfterDelete钩子）
	fmt.Println("\n🗑️ 删除评论（测试钩子函数）:")
	err = db.Delete(&newComment).Error
	if err != nil {
		panic(err)
	}
	
	// 删除文章（触发AfterDelete钩子）
	fmt.Println("\n🗑️ 删除文章（测试钩子函数）:")
	err = db.Delete(&newPost).Error
	if err != nil {
		panic(err)
	}
	
	// 5. 显示最终统计
	fmt.Println("\n5️⃣ 最终统计")
	
	var userCount, postCount, commentCount int64
	db.Model(&User{}).Count(&userCount)
	db.Model(&Post{}).Count(&postCount)
	db.Model(&Comment{}).Count(&commentCount)
	
	fmt.Printf("用户总数: %d\n", userCount)
	fmt.Printf("文章总数: %d\n", postCount)
	fmt.Printf("评论总数: %d\n", commentCount)
	
	fmt.Println("\n✅ 博客系统GORM演示完成!")
}
