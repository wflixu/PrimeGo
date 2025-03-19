# PrimeGin

PrimeGin 是一个基于 Go 和 Vue 的全栈项目，后端使用 Gin 框架，前端使用 Vue.js。

## 项目结构

```
PrimeGin/
├── server/         # 后端目录
│   ├── controller/ # 控制器目录
│   │   └── ping.go # 示例控制器
│   ├── model/      # 数据模型目录
│   │   └── model.go # 示例模型
│   ├── router/     # 路由目录
│   │   └── router.go # 路由初始化
│   └── main.go     # 后端入口文件
├── frontend/       # 前端 Vue 项目目录
├── go.mod          # Go 模块配置文件
├── README.md       # 项目说明文件
├── .gitignore      # Git 忽略文件
```

## 后端

后端使用 Gin 框架，数据库使用 PostgreSQL，ORM 使用 GORM。

### 启动后端

```bash
cd server
go run main.go
```

## 前端

前端使用 Vue.js，位于 `frontend/` 目录。

### 启动前端

```bash
cd frontend
npm install
npm run serve
```

## 贡献

欢迎提交 Issue 和 Pull Request！
