## swag安装
### 安装命令
go get -u github.com/swaggo/swag/cmd/swag
参考 https://github.com/swaggo/gin-swagger#readme

### swagger 文档生成命令
- (可选) 使用fmt格式化 SWAG 注释
```
swag fmt
```

- 生成doc文档
```
swag init --pd --parseInternal -p pascalcase
```

### 文档代码参考：
https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html#param-type