## 说明

模块说明：

- resource：RESTful API Resource
- service: resource的实现（特别简单的实现可以直接写在resource里面）
- config: 配置
- logger：日志
- docs：swagger文档


## 项目说明

### 搭建 

项目使用`go module`管理，首次搭建项目：

```shell
git clone xxxxx
# 进入项目根目录
go mod tidy
go mod download
go mod vendor
```

如果有新增依赖，无需使用`go get`，直接在代码中引入依赖，然后执行

```
go mod tidy
go mod download
go mod vendor
```

更新模块依赖即可。

### API文档

API文档使用swagger，方法：

1. 在代码中按照[规范](https://swaggo.github.io/swaggo.io/declarative_comments_format/)注释；
2. 下载`swag`工具：`go get -u github.com/swaggo/swag/cmd/swag`；
3. 在项目根目录（包含`main.go`）执行`swag init`；
4. 访问http://localhost:8080/swagger/index.html.

后面每次增加/修改swagger注解都要执行`swag init`。

### 依赖

- API：[Gin Web Framework](https://github.com/gin-gonic/gin)
- API文档：[gin-swagger](https://github.com/swaggo/gin-swagger)
- 日志：[Zab](https://github.com/uber-go/zap)
