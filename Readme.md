# go-project-frame
<p >
  <a href="https://golang.google.cn/">
    <img src="https://img.shields.io/badge/Golang-1.19-green.svg" alt="golang">
  </a>
  <a href="https://gin-gonic.com/">
    <img src="https://img.shields.io/badge/Gin-1.8.1-red.svg" alt="gin">
  </a>
  <a href="https://gorm.io/">
    <img src="https://img.shields.io/badge/Gorm-1.25.5-orange.svg" alt="gorm">
  </a>
</p>

> go-project-frame是一个快速搭建golang项目的三层架构后端框架，使用了目前最主流的gin+gorm进行搭建。
> - 对于初学golang开发的同学来说，可以作为练手项目，快速体验golang项目的搭建过程，也可以基于此项目快速开发业务功能
> - 对于企业来说，可以作为二次开发的模板，快速搭建项目

技术栈选型当下最主流框架，后端使用Gin+GORM，前后端分离开发模式（前端自行开发），其中内置了kuberenetes配置，使用client-go与K8S交互，可以用于开发云原生项目

## 开始部署
### 初始化数据库
需要手动创建数据库，数据表与数据会通过`InitDBService`自动初始化

```sql
# 以实际使用到的数据库为准，这里用kubemanage为例
CREATE DATABASE kubemanage;
```
### 运行工程
后端
- 修改配置文件config.yaml

- 请确保用户名/./kube  文件夹下存在k8s的kubeconfig文件

  - 开始前请设置配置文件环境变量`KUBEMANAGE-CONFIG`，或通过命令行参数`configFile`指定，配置文件优先级: 默认配置 < 环境变量< 命令行

    ``` shell
    git clone https://github.com/graham924/go-project-frame.git
    
    cd go-project-frame
    
    go mod tidy
    
    go run cmd/main.go
    ```
- 然后可以使用 postman 或 apifox，发送127.0.0.1:6180/api/user/login，测试项目是否运行

## Issue 规范
- issue 仅用于提交 Bug 或 Feature 以及设计相关的内容，其它内容可能会被直接关闭。

- 在提交 issue 之前，请搜索相关内容是否已被提出。

## Pull Request 规范
- 请先 fork 一份到自己的项目下，在自己项目下新建分支。

- commit 信息要以`feat(model): 描述信息` 的形式填写，例如 `fix(user): fix xxx bug / feat(user): add xxx`。

- 如果是修复 bug，请在 PR 中给出描述信息。

- 合并代码需要两名维护人员参与：一人进行 review 后 approve，另一人再次 review，通过后即可合并。
