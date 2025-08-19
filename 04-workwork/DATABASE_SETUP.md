# 数据库配置说明

## MySQL数据库设置

### 1. 安装MySQL

#### Windows
1. 下载MySQL安装包：https://dev.mysql.com/downloads/installer/
2. 运行安装程序，选择"Server only"或"Custom"安装
3. 设置root密码为"root"（或修改代码中的密码）
4. 完成安装

#### macOS
```bash
# 使用Homebrew安装
brew install mysql

# 启动MySQL服务
brew services start mysql

# 设置root密码
mysql_secure_installation
```

#### Linux (Ubuntu/Debian)
```bash
# 安装MySQL
sudo apt update
sudo apt install mysql-server

# 启动MySQL服务
sudo systemctl start mysql
sudo systemctl enable mysql

# 设置root密码
sudo mysql_secure_installation
```

### 2. 创建数据库

连接到MySQL并创建数据库：

```bash
# 连接到MySQL
mysql -u root -p

# 在MySQL中执行以下命令
CREATE DATABASE gorm CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

或者使用以下命令直接创建：

```bash
mysql -u root -p -e "CREATE DATABASE gorm CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
```

### 3. 验证数据库连接

```bash
# 测试连接
mysql -u root -p -e "USE gorm; SHOW TABLES;"
```

### 4. 修改数据库配置（可选）

如果需要修改数据库连接配置，请编辑 `main.go` 文件中的第16行：

```go
// 当前配置
db, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))

// 修改示例（如果密码不是root）
db, err = gorm.Open(mysql.Open("root:your_password@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))

// 修改示例（如果数据库名不是gorm）
db, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/your_database?charset=utf8mb4&parseTime=True&loc=Local"))
```

### 5. 数据库连接参数说明

连接字符串格式：`username:password@tcp(host:port)/database?parameters`

参数说明：
- `charset=utf8mb4`: 使用UTF-8字符集，支持emoji等特殊字符
- `parseTime=True`: 自动解析时间类型
- `loc=Local`: 使用本地时区

### 6. 常见问题

#### 问题1: 连接被拒绝
```
Error 1045: Access denied for user 'root'@'localhost'
```
**解决方案**: 检查用户名和密码是否正确

#### 问题2: 数据库不存在
```
Error 1049: Unknown database 'gorm'
```
**解决方案**: 确保已创建gorm数据库

#### 问题3: 字符集问题
```
Error 1366: Incorrect string value
```
**解决方案**: 确保数据库使用utf8mb4字符集

### 7. 生产环境建议

1. **创建专用用户**: 不要使用root用户
```sql
CREATE USER 'blog_user'@'localhost' IDENTIFIED BY 'strong_password';
GRANT ALL PRIVILEGES ON gorm.* TO 'blog_user'@'localhost';
FLUSH PRIVILEGES;
```

2. **配置连接池**: 在代码中添加连接池配置
```go
db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
})
if err != nil {
    log.Fatal("Failed to connect to database:", err)
}

sqlDB, err := db.DB()
if err != nil {
    log.Fatal("Failed to get database instance:", err)
}

// 设置连接池参数
sqlDB.SetMaxIdleConns(10)
sqlDB.SetMaxOpenConns(100)
sqlDB.SetConnMaxLifetime(time.Hour)
```

3. **使用环境变量**: 将数据库配置放在环境变量中
```go
dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_HOST"),
    os.Getenv("DB_PORT"),
    os.Getenv("DB_NAME"),
)
```
