# PrimeGin

PrimeGin 是一个基于 Go 和 Vue 的全栈项目，后端使用 Gin + Gorm框架，前端使用 Vue.js + primevue。

## 项目简介

PrimeGin 是一个全栈开发模板，旨在快速构建现代 Web 应用程序。后端采用 Gin 框架，结合 PostgreSQL 和 GORM 实现高效的数据处理；前端基于 Vue.js，提供响应式用户界面。

## 项目结构


```
# 项目目录结构
PrimeGin/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── controllers/
│   │   ├── user_controller.go
│   │   └── auth_controller.go
│   ├── models/
│   │   ├── user.go
│   │   ├── role.go
│   │   └── permission.go
│   ├── routes/
│   │   └── routes.go
│   ├── services/
│   │   ├── user_service.go
│   │   └── auth_service.go
│   └── utils/
│       └── database.go
├── pkg/
├── go.mod
└── go.sum
```

## 快速开始

### 后端

后端使用 Gin 框架，数据库使用 PostgreSQL，ORM 使用 GORM。

#### 启动后端

1. 安装依赖：
   ```bash
   go mod tidy
   ```

2. 配置数据库连接：
   修改 `internal/config/config.go` 文件中的数据库配置。

3. 启动服务：
   ```bash
   go run cmd/server/main.go
   ```

#### 后端目录说明

- `controllers/`：存放业务逻辑控制器。
- `models/`：定义数据库模型。
- `route/`：路由初始化。
- `config/`：存放配置文件。
- `middleware/`：定义中间件。

### 前端

前端使用 Vue.js，位于 `src/` 目录。

#### 启动前端

1. 安装依赖：
   ```bash
   cd frontend
   npm install
   ```

2. 启动开发服务器：
   ```bash
   npm run serve
   ```

#### 前端目录说明

- `src/`：存放前端源码，包括组件、路由等。
- `public/`：存放静态资源。

## 贡献指南

欢迎提交 Issue 和 Pull Request！在贡献代码前，请确保遵循以下步骤：

1. Fork 本仓库。
2. 创建新分支：`git checkout -b feature/your-feature-name`
3. 提交更改：`git commit -m "Add your message"`
4. 推送分支：`git push origin feature/your-feature-name`
5. 创建 Pull Request。

## 许可证

本项目采用 MIT 许可证，详情请参阅 [LICENSE](./LICENSE) 文件。
