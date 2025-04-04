# PrimeGo

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/yourusername/PrimeGo/blob/main/LICENSE)

PrimeGo 是一个基于 Go 和 Vue 的全栈项目模板，后端使用 Hertz + GORM 框架，前端使用 Vue.js + PrimeVue。

## 功能特性

- 后端使用 Hertz 高性能 HTTP 框架
- 集成 GORM 进行数据库操作
- 支持 PostgreSQL 数据库
- 前端使用 Vue 3 + TypeScript
- 集成 PrimeVue UI 组件库
- 开箱即用的用户认证模块
- 完善的 API 文档生成
- 前后端分离架构
- 支持 Docker 部署

## 项目结构


```
PrimeGin/
├── biz/                      # 后端业务逻辑目录
│   ├── handler/              # HTTP请求处理器
│   │   ├── hello/            # hello示例处理器
│   │   └── home.go           # 首页处理器
│   ├── middleware/           # 中间件
│   │   └── log.go            # 日志中间件
│   ├── model/                # 数据模型定义
│   └── router/               # 路由配置
│       ├── hello/            # hello路由
│       └── routes.go         # 路由注册
├── env.d.ts                  # 环境变量类型定义
├── go.mod                    # Go模块定义
├── go.sum                    # Go模块校验
├── index.html                # 前端入口HTML
├── main.go                   # 后端主入口
├── package.json              # 前端依赖配置
├── public/                   # 公共静态资源
├── script/                   # 脚本目录
│   └── bootstrap.sh          # 项目初始化脚本
├── src/                      # 前端源代码
│   ├── App.vue               # 根组件
│   ├── assets/               # 静态资源
│   ├── components/           # 公共组件
│   ├── main.ts               # 前端主入口
│   ├── router/               # 前端路由
│   │   └── index.ts          # 路由配置
│   ├── stores/               # 状态管理
│   │   └── counter.ts        # 计数器状态
│   └── views/                # 页面视图
│       ├── AboutView.vue     # 关于页面
│       └── HomeView.vue      # 首页
└── vite.config.ts            # Vite构建配置
```

## 目录功能说明

- **biz/**: 后端业务逻辑核心目录，包含处理器、中间件、模型和路由
- **src/**: 前端源代码目录，包含Vue组件、路由和状态管理
- **public/**: 公共静态资源，如favicon等
- **script/**: 项目初始化和管理脚本
- **config文件**: 包含项目各种配置，如TypeScript、Vite等


## 快速开始

### 系统要求

- Go 1.18+
- Node.js 16+
- PostgreSQL 12+ 或 MySQL 8.0+

### 安装

1. 克隆仓库
   ```bash
   git clone https://github.com/yourusername/PrimeGo.git
   cd PrimeGo
   ```

2. 安装后端依赖
   ```bash
   go mod download
   ```

3. 安装前端依赖
   ```bash
   npm install
   ```

### 配置

1. 复制示例配置文件
   ```bash
   cp config.example.yaml config.yaml
   ```

2. 修改 `config.yaml` 中的数据库配置

### 运行

启动后端服务
```bash
go run main.go
```

启动前端开发服务器
```bash
npm run dev
```

### 生产环境部署

使用 Docker 部署
```bash
docker-compose up -d
```

## 贡献指南

我们欢迎任何形式的贡献！在提交 Pull Request 前请确保：

1. Fork 仓库并创建新分支
2. 遵循代码风格规范
3. 添加适当的测试
4. 更新相关文档
5. 提交清晰的 commit 信息

### 代码风格

- 后端代码遵循标准 Go 代码风格
- 前端代码遵循 Vue 风格指南
- 提交信息遵循 Conventional Commits 规范

## 许可证

MIT License © 2023-Present YourName

## 贡献指南

欢迎提交 Issue 和 Pull Request！在贡献代码前，请确保遵循以下步骤：

1. Fork 本仓库。
2. 创建新分支：`git checkout -b feature/your-feature-name`
3. 提交更改：`git commit -m "Add your message"`
4. 推送分支：`git push origin feature/your-feature-name`
5. 创建 Pull Request。

## 许可证

本项目采用 MIT 许可证，详情请参阅 [LICENSE](./LICENSE) 文件。
