# 个人博客系统后端

这是一个使用 Go 语言、Gin 框架和 GORM 开发的个人博客系统后端，实现了完整的博客文章管理功能，包括用户认证、文章CRUD操作和评论功能。

## 功能特性

- ✅ 用户注册和登录（JWT认证）
- ✅ 文章的创建、读取、更新、删除（CRUD）
- ✅ 评论功能
- ✅ 权限控制（只有作者可以修改自己的文章）
- ✅ 分页查询
- ✅ 统一的错误处理和响应格式
- ✅ MySQL数据库（生产环境推荐）

## 技术栈

- **语言**: Go 1.24+
- **Web框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL
- **认证**: JWT
- **密码加密**: bcrypt

## 项目结构

```
blog-system/
├── main.go          # 主程序入口
├── models.go        # 数据模型定义
├── auth.go          # 用户认证相关功能
├── posts.go         # 文章管理功能
├── comments.go      # 评论管理功能
├── go.mod           # Go模块依赖
├── go.sum           # 依赖校验文件
├── blog.db          # MySQL数据库（需要预先创建）
└── README.md        # 项目说明文档
```

## 安装和运行

### 1. 环境要求

- Go 1.24 或更高版本
- MySQL 5.7+ 或 MySQL 8.0+
- Git

### 2. 克隆项目

```bash
git clone <repository-url>
cd blog-system
```

### 3. 配置数据库

确保MySQL服务正在运行，并创建数据库：

```sql
CREATE DATABASE gorm CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 4. 安装依赖

```bash
go mod tidy
```

### 5. 运行项目

```bash
go run .
```

服务器将在 `http://localhost:8080` 启动。

## API 文档

### 基础信息

- **Base URL**: `http://localhost:8080/api`
- **认证方式**: Bearer Token (JWT)
- **响应格式**: JSON

### 认证相关

#### 用户注册

```http
POST /api/auth/register
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123",
  "email": "test@example.com"
}
```

**响应示例**:
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "user_id": 1,
    "username": "testuser",
    "email": "test@example.com"
  }
}
```

#### 用户登录

```http
POST /api/auth/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

**响应示例**:
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user_id": 1,
    "username": "testuser",
    "email": "test@example.com"
  }
}
```

### 文章管理

#### 获取文章列表

```http
GET /api/posts?page=1&limit=10
```

#### 获取单个文章

```http
GET /api/posts/{id}
```

#### 创建文章（需要认证）

```http
POST /api/posts
Authorization: Bearer <your-jwt-token>
Content-Type: application/json

{
  "title": "我的第一篇博客",
  "content": "这是博客的内容..."
}
```

#### 更新文章（需要认证）

```http
PUT /api/posts/{id}
Authorization: Bearer <your-jwt-token>
Content-Type: application/json

{
  "title": "更新后的标题",
  "content": "更新后的内容..."
}
```

#### 删除文章（需要认证）

```http
DELETE /api/posts/{id}
Authorization: Bearer <your-jwt-token>
```

### 评论管理

#### 获取文章评论

```http
GET /api/comments/post/{postId}?page=1&limit=20
```

#### 创建评论（需要认证）

```http
POST /api/comments
Authorization: Bearer <your-jwt-token>
Content-Type: application/json

{
  "content": "这是一条评论",
  "post_id": 1
}
```

## 测试用例

### 使用 Postman 测试

1. **注册用户**
   - Method: POST
   - URL: `http://localhost:8080/api/auth/register`
   - Body: 
     ```json
     {
       "username": "testuser",
       "password": "password123",
       "email": "test@example.com"
     }
     ```

2. **登录获取Token**
   - Method: POST
   - URL: `http://localhost:8080/api/auth/login`
   - Body:
     ```json
     {
       "username": "testuser",
       "password": "password123"
     }
     ```

3. **创建文章**
   - Method: POST
   - URL: `http://localhost:8080/api/posts`
   - Headers: `Authorization: Bearer <token-from-login>`
   - Body:
     ```json
     {
       "title": "测试文章",
       "content": "这是测试文章的内容"
     }
     ```

4. **获取文章列表**
   - Method: GET
   - URL: `http://localhost:8080/api/posts`

5. **创建评论**
   - Method: POST
   - URL: `http://localhost:8080/api/comments`
   - Headers: `Authorization: Bearer <token-from-login>`
   - Body:
     ```json
     {
       "content": "这是一条测试评论",
       "post_id": 1
     }
     ```

## 错误处理

系统使用统一的错误响应格式：

```json
{
  "success": false,
  "error": "错误描述信息"
}
```

常见HTTP状态码：
- `200`: 成功
- `201`: 创建成功
- `400`: 请求参数错误
- `401`: 未认证
- `403`: 权限不足
- `404`: 资源不存在
- `409`: 资源冲突（如用户名已存在）
- `500`: 服务器内部错误

## 数据库

项目使用MySQL数据库，需要预先创建数据库。数据库包含以下表：

- `users`: 用户表
- `posts`: 文章表
- `comments`: 评论表

### 数据库配置

默认数据库连接配置：
- 主机: 127.0.0.1:3306
- 用户名: root
- 密码: root
- 数据库名: gorm
- 字符集: utf8mb4

如需修改配置，请编辑 `main.go` 中的数据库连接字符串。

## 安全特性

- 密码使用bcrypt加密存储
- JWT token认证
- 权限控制（用户只能操作自己的资源）
- 输入验证和错误处理

## 开发说明

### 添加新功能

1. 在相应的文件中添加新的处理函数
2. 在 `main.go` 中添加路由
3. 更新模型定义（如需要）

### 配置修改

- JWT密钥: 修改 `auth.go` 中的 `JWTSecret` 常量
- 数据库: 修改 `main.go` 中的数据库连接配置
- 端口: 修改 `main.go` 中的 `r.Run(":8080")`

## 许可证

MIT License
