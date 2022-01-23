# Go + Vue3 Study

## 环境安装

本地环境：Go 1.17

后端框架：Gin

注意Go在使用Go Module的话需要使用修改Go的代理

首先查看Go相关的环境变量

> go env

修改Go代理

> go env -w Go111MODULE=on
> 
> go env -w GOPROXY=https://goproxy.cn,direct

## 后端框架

### Gin

#### 框架下载

> go get -u github.com/gin-gonic/gin

这里注意代理一定要配置好否则下载会非常慢

#### 框架引入

实例代码如下所示：

```go
package main

import ("github.com/gin-gonic/gin")

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "pong"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
```

首先通过import语句引入相关的包

在main函数中，初始化一个gin类型，对于一个GET请求，返回一个status是200且内容是{"msg": "pong"}的信息

最后开启该项目，默认在localhost:8080启动

使用POSTMAN进行测试

这里简单介绍一下**GET**和**POST**请求的区别：

GET请求是把参数放在URL中的一种TCP/IP协议，在进行传输时只进行一次传输。一般来说服务器会对GET请求的大小进行限制

POST请求是把参数放到Request Body中传输，在进行传输过程中会进行两次传输。一般对于POST的大小限制会弱于GET的大小限制

#### 框架语法

H: Gin中的一个语法，相当于map[string]{interface}

### Gorm

#### 框架下载

> go get -u github.com/jinzhu/gorm

#### 配置文件

## Go语法相关

### 项目入口

与C和C++类似，整个Go项目有一个文件入口main.go，代码如下方所示

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world!")
}
```

编译上面的代码需要下面的编译命令：

> go run main.go

上面的代码成功输出hello world!

### 项目初始化

Go 项目需要先进行初始化，下面是命令

> go mod init <name>

name是项目名称，一般取<person/school/corporation>/<project_name>

例如 **ZerglingChen3/MyBlog**

初始化就会出现下面这个文件：

```mod
module ZerglingChen3/MyBlog

go 1.17
```

上面是模块名，下面是Go的版本

### 结构体

与C++中结构体类似，Go中同样有结构体的定义

```go
type User struct {
    gorm.Model
    Name      string `gorm:"type:varchar(20);not null"`
    Telephone string `gorm:"varchar(10);not null;unique"`
    Password  string `gorm:"size:255;not null"`
}
```

可以看到对于数据库中结构体的定义，类似于SQL语言的Create Table.

