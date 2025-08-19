#!/bin/bash

# API测试脚本
BASE_URL="http://localhost:8080/api"

echo "=== 博客系统API测试 ==="

# 1. 注册用户
echo "1. 注册用户..."
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123",
    "email": "test@example.com"
  }')
echo "注册响应: $REGISTER_RESPONSE"

# 2. 登录获取token
echo -e "\n2. 用户登录..."
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }')
echo "登录响应: $LOGIN_RESPONSE"

# 提取token
TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*"' | cut -d'"' -f4)
echo "获取到的Token: $TOKEN"

# 3. 创建文章
echo -e "\n3. 创建文章..."
CREATE_POST_RESPONSE=$(curl -s -X POST "$BASE_URL/posts" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "测试文章",
    "content": "这是一篇测试文章的内容"
  }')
echo "创建文章响应: $CREATE_POST_RESPONSE"

# 4. 获取文章列表
echo -e "\n4. 获取文章列表..."
GET_POSTS_RESPONSE=$(curl -s -X GET "$BASE_URL/posts")
echo "文章列表响应: $GET_POSTS_RESPONSE"

# 5. 创建评论
echo -e "\n5. 创建评论..."
CREATE_COMMENT_RESPONSE=$(curl -s -X POST "$BASE_URL/comments" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "content": "这是一条测试评论",
    "post_id": 1
  }')
echo "创建评论响应: $CREATE_COMMENT_RESPONSE"

# 6. 获取评论列表
echo -e "\n6. 获取评论列表..."
GET_COMMENTS_RESPONSE=$(curl -s -X GET "$BASE_URL/comments/post/1")
echo "评论列表响应: $GET_COMMENTS_RESPONSE"

echo -e "\n=== 测试完成 ==="
