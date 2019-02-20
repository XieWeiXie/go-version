# go-version

> Go 构建开源库的流程


- 项目目录
- 版本包管理
- godoc
- gopkg.in
- travisCI
- badges

### 1. 项目目录

**方式一**
``` 
src/github.com/user/projectName
```

example:

``` 
src/github.com/wuxiaoxiaoshen/go-version
```

则下载方式为 ： 

``` 
go get github.com/wuxiaoxiaoshen/go-version
```

**方式二**

``` 
src/gopkg.in/user/project
```

example:

``` 
src/gopkg.in/wuxiaoxiaoshen/go-version
```

则下载方式为 ： 

``` 
go get gopkg.in/wuxiaoxiaoshen/go-version
```
### 版本包管理

推荐使用 go mod 

> go1.11 及以上

``` 
go mod init github.com/wuxiaoxiaoshen/go-version
go mod tidy

>> 

生成： go.mod，go.sum

```

### godoc 

开源项目文档，要求项目内编写注释，函数或者文件（doc.go）

编写好注释，推送到 github 上， 访问 `https://godoc.org/github.com/wuxiaoxiaoshen/go-version` 即可