# 基于toolbox 的demo 实例。

# Golang工程化标准

- ### 背景

  当越来越多的人参与团队开发中,需要结构化和工程化协同方便，统一管理，规范化。


- ### 实践

### 项目结构

- #### 项目结构

 ```
  shell
  ├── boot
  |	├─boot.go  # 配置初始化文件
  ├─ config      # 配置文件夹
  |   |	   └── config.yaml #配置
  ├─ docs				# swagger文档
  |	├── docs.go
  ├─internal   # 业务逻辑层
  |   ├─ enum 
  |	|	├── consts.go # 常量定义
  |   ├── middleware # 中间件文件
  |	|	├── auth.go # 用户认证
  |   ├── model	# 结构体文件夹定义
  |   ├── route # 路由文件夹
  |   ├── service # 服务层文件家
  |	|   ├── service_name # 服务模块名
  |	|	|  ├── demo_controller.go # 执行控制器
  ├─ .gitignore # git忽略文件
  ├─ main.go # 项目入口
  ├─ README.md
  ├─ go.sum
  └ go.mod
```



  