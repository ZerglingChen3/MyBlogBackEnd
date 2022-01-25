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

对于参与数据库的修改，我们需要引入gorm框架来方便对数据库内容进行修改。

#### 框架下载

> go get -u github.com/jinzhu/gorm

#### 框架引入

```go
import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)
```

#### 数据库相关操作

##### 创建数据库

```go
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
```
在这里db是创建的数据库，数据库存储文件为'test.db'，数据库配置文件这里用的是默认配置。

##### 创建模式

在创建数据库后，我们还需要创建对应的模式（或者说建表），使用下面的语句
```go
err = db.AutoMigrate(&User{})
```

##### 查询语句

查询电话号码是否在数据库中出现过

```go
db.Where("telephone = ?", telephone).First(&user)
```

相当于下面的SQL语句：

```sql
SELECT *
FROM user 
WHERE user.telephone = telephone
ORDER BY user.id
LIMIT 1
```

### jwt

主要用于token的发放和处理

首先简单介绍一下token的使用方法：

在服务器的身份验证时，当客户端使用用户名和密码登录之后，服务器收到请求之后会去验证用户名和密码。

当验证通过之后，服务器会返回一个token用于验证，客户端需要保存这个token。

在之后的每次请求需要携带这个token用于身份验证，服务器在拿到token之后可以直接验证是否正确，若正确则向客户端返回数据。

#### 框架下载

> go get github.com/dgrijalva/jwt-go

#### Token组成

头部(Header): 标记类型和使用的协议

有效载荷(Playload): 存放包括签发者、签发时间、过期时间等需要写入token中的信息

签名(signature): 将Header和Playload拼接起来通过Base64编码和加密算法HS256和秘钥拼接到JWT的后面

#### 中间件

为了对我们的路由进行保护，比如说当我们只能让自己访问自己的个人主页，而不允许其他人访问自己主页时，需要进行权限控制

具体来说，我们的GET请求需要经过一个中间件进行权限(token)审查，通过之后通过全局变量设置我们需要的信息

之后再访问该路由之后，将直接读取全局信息的内容


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
    Name      string `gorm:"type:varchar(20);not null"`
    Telephone string `gorm:"varchar(10);not null;unique"`
    Password  string `gorm:"size:255;not null"`
}
```

可以看到对于结构体的定义，类似于SQL语言的Create Table.

### 占位符

在fmt中需要给特定的格式符合，下面介绍几种

%v: 相应值的默认格式

## 项目结构

为了方便整个项目更加容易维护，因此用文件夹将文件分离出来

common: 数据库的创建和获得

controller: 模型操作的控制

model: 数据库映射的结构体

middleware: 中间件相关组件

dto: 数据过滤

response: HTTP封装

util: 方法类

routes.go: 所有的路由在这里初始化

## 基本功能创建

### 登录功能

这里需要注意一个问题就是在数据库中，用户的密码不能以明文存取。

因此我们需要使用密文存取的方法，具体来说在创建用户的时候，需要使用下面的HASH函数。

```go
encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
```

而在验证用户密码的使用，使用下面的方法对比HASH密码。

```go
err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
```

如果返回err说明输入的密码并不正确。

