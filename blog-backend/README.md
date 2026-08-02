# 个人博客系统后端

使用 Go 语言结合 Gin 框架和 GORM 库开发的个人博客系统后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能。

## 功能特性

- **用户认证**：用户注册、登录，使用 JWT 进行身份认证
- **文章管理**：文章的创建、读取、更新、删除，只有作者本人可以修改/删除自己的文章
- **评论功能**：对文章发表评论，支持删除评论，只有评论作者本人可以删除
- **错误处理**：统一处理各类错误，返回合适的 HTTP 状态码和错误信息
- **日志记录**：使用 Gin 内置日志中间件记录请求日志，方便调试维护

## 技术栈

- **Web 框架**：[Gin](https://github.com/gin-gonic/gin)
- **ORM**：[GORM](https://gorm.io)
- **数据库**：MySQL
- **认证**：JWT ([golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt))
- **密码加密**：bcrypt
- **配置管理**：godotenv

## 环境要求

- Go 1.19+
- MySQL 5.7+

## 安装步骤

1. 克隆项目

```bash
git clone <repository-url>
cd blog-backend
```

2. 安装依赖

```bash
go mod download
```

3. 配置环境变量

复制环境变量模板文件：

```bash
cp .env.example .env.development
```

编辑 `.env.development`，修改数据库连接信息和 JWT 密钥：

```
MODE=development
SERVER_PORT=8080
DB_USER=root
DB_PASSWORD=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=blog
JWT_SECRET=your_jwt_secret_key
```

4. 创建数据库

在 MySQL 中创建数据库：

```sql
CREATE DATABASE blog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

程序启动时会自动迁移创建表结构。

## 启动运行

### 开发环境

```bash
# 设置环境变量
export MODE=development
# 或 Windows PowerShell
$env:MODE = "development"

# 运行项目
go run main.go
```

使用 air 进行热重载开发（需要先安装 air）：

```bash
air
```

### 生产环境编译

```bash
go build -o blog-backend .
./blog-backend
```

## API 接口

### 用户相关

| 方法 | 路径 | 说明 | 需要认证 |
|------|------|------|----------|
| POST | `/user/register` | 用户注册 | 否 |
| POST | `/user/login` | 用户登录 | 否 |

### 文章相关

| 方法 | 路径 | 说明 | 需要认证 |
|------|------|------|----------|
| GET | `/post/allList` | 获取所有文章列表 | 否 |
| GET | `/post/:PostID` | 获取单个文章详情 | 否 |
| POST | `/post/create` | 创建文章 | 是 |
| PUT | `/post/update` | 更新文章 | 是 |
| DELETE | `/post/delete/:PostID` | 删除文章 | 是 |

### 评论相关

| 方法 | 路径 | 说明 | 需要认证 |
|------|------|------|----------|
| GET | `/comment/list/:PostID` | 获取文章评论列表 | 否 |
| POST | `/comment/create` | 创建评论 | 是 |
| DELETE | `/comment/delete/:CommentID` | 删除评论 | 是 |

## 请求/响应示例

### 用户注册

**请求**：
```json
POST /user/register
{
  "username": "testuser",
  "email": "test@example.com",
  "password": "123456"
}
```

**响应**：
```json
{
  "code": 200,
  "message": "注册成功"
}
```

### 用户登录

**请求**：
```json
POST /user/login
{
  "username": "testuser",
  "password": "123456"
}
```

**响应**：
```json
{
  "code": 200,
  "message": "登录成功",
  "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 创建文章

**请求头**：
```
Authorization: Bearer <accessToken>
Content-Type: application/json
```

**请求体**：
```json
{
  "title": "文章标题",
  "content": "文章内容"
}
```

**响应**：
```json
{
  "code": 200,
  "message": "文章创建成功"
}
```

## 项目结构

```
blog-backend/
├── middleware/       # 中间件
│   ├── CORS.go      # CORS 跨域处理
│   ├── JWT.go       # JWT 认证中间件
│   └── Logger.go    # 日志中间件
├── router/          # 路由
│   └── router.go    # 路由注册
├── server/          # 处理函数
│   ├── user.go      # 用户相关
│   ├── posts.go     # 文章相关
│   └── comment.go   # 评论相关
├── types/           # 数据类型定义
│   ├── user.go
│   ├── post.go
│   └── comment.go
├── main.go          # 入口文件
├── go.mod           # Go 模块定义
└── README.md        # 项目说明
```

## 错误码说明

| HTTP 状态码 | 说明 |
|-------------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 认证失败/未授权 |
| 403 | 无权限操作 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

## 许可证

MIT
