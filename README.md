# go-version

[![GoDoc](https://godoc.org/github.com/wuxiaoxiaoshen/go-version?status.svg)](https://godoc.org/github.com/wuxiaoxiaoshen/go-version)
[![Go Report Card](https://goreportcard.com/badge/github.com/wuxiaoxiaoshen/go-version)](https://goreportcard.com/report/github.com/wuxiaoxiaoshen/go-version)

> Go 构建开源库的流程


- 项目目录
- 版本包管理
- godoc
- gopkg.in
- travisCI
- badges
- markdown

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
### 2. 版本包管理

推荐使用 go mod 

> go1.11 及以上

``` 
go mod init github.com/wuxiaoxiaoshen/go-version
go mod tidy

>> 

生成： go.mod，go.sum，推送到仓库即可

如何下载库？

go mod vendor 即可

```

### 3. godoc 

开源项目文档，要求项目内编写注释，函数或者文件（doc.go）

编写好注释，推送到 github 上， 访问 `godoc.org/github.com/wuxiaoxiaoshen/go-version` 即可

### 4. gopkg.in 

这是什么？

也是一种版本管理方式，和直接 使用 github.com 前缀的地址不同，它将网址进行了缩写，使的项目的导入地址更为清晰、简洁

一般形式为：
``` 
gopkg.in/pkg.v3      → github.com/go-pkg/pkg (branch/tag v3, v3.N, or v3.N.M)
gopkg.in/user/pkg.v3 → github.com/user/pkg (branch/tag v3, v3.N, or v3.N.M)


>>

左侧为简化路径，右侧为真实路径
v3 这种形式，使根据代码仓库的 branch 或者 tag 自动识别的，推荐使用 tag 的形式
```

git tag  的使用

``` 
项目内执行一下命令：

>> git tag v1.0 or git tag -a v1.0 -m "comment" // 创建 tag
>> git push origin v1.0 //  推送到远端仓库

```

这种方式，项目目录为 ：gopkg.in/user/project， 实际项目仓库地址为：github.com/user/project

之后根据 tag 访问：gopkg.in/user/project.tag 即可

> gopkg.in 只是服务器上进行了转发，并不保存真实的代码，同时它带上了软件的版本

则下载项目和导入的使用方式如下：

```
下载
go get gopkg.in/user/project.tag
 
导入
import "gopkg.in/user/project.tag"

```


### 5. travisCI

持续集成，即对代码进行持续集成，查看是否成功，持续集成，一般执行单元测试、功能测试等，检验新合入的代码是否正确，能够推送到生成环境中


### 6. badge

即：徽章，用来对项目的一些"美化"动作，比如是否构建成功，是个测试成功，下载量，发行版本，版本号等， 通常是一个svg(矢量图像)，经 markdown 渲染。

比如：本项目的文档的 badge 即为：[![GoDoc](https://godoc.org/github.com/wuxiaoxiaoshen/go-version?status.svg)](https://godoc.org/github.com/wuxiaoxiaoshen/go-version)

有些第三方服务，比如 travisCI ，构建成功之后，也有badge 

### 7. markdown

标记语言，一般的项目 README.md 即使用的是 markdown 语法编写的，事实上 markdown 已经成为使用 github 的标配



### Reference

- [gopkg.in](http://labix.org/gopkg.in)
- [gopkg.in 是什么](https://www.jianshu.com/p/fae12384cc29)
- [badge 集合](https://shields.io/)
- [travisCI 持续集成](https://travis-ci.org)
- [ci.appveyor 持续集成](https://ci.appveyor.com)
- [coverall 覆盖率](https://coveralls.io)






