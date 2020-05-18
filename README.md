### 个人博客系统

基于Go语言的一个简易markdown博客系统 

示例：[codwiki.cn](https://codwiki.cn) 

源码：[github.com/MilkyMoon/goblog](https://github.com/MilkyMoon/goblog) 

#### 项目目的

通过实现一个简单的博客系统来学习Go语言的开发，在个人学习的同时也给其他正在学习Golang的同学一个参考。本项目将不依赖数据库存储，主要通过md文件的形式存储，也方便编辑。

#### 项目需求

1. 实现一个简单的博客网站
2. 实现一个简单的markdown文档阅读器
3. 实现一个简单的缓存管理
4. 通过文件夹实现阅读器的目录结构
5. 通过github的Webhooks实现自动更新

#### 项目依赖

- [github.com/kataras/iris](https://github.com/kataras/iris) Web框架
- [github.com/pelletier/go-toml](https://github.com/pelletier/go-toml) toml配置文件读取

#### 项目运行

下载源码

```shell script
mkdir $GOPATH/src/codwiki.com/goblog
cd $GOPATH/src/codwiki.com/goblog
git clone https://github.com/MilkyMoon/goblog.git
```

下载依赖

```shell
go get github.com/kataras/iris
go get github.com/pelletier/go-toml
```

编译项目

```shell script
go build -o goblog main.go
```

运行项目

```shell script
./goblog
#注意执行文件要与静态资源文件路径一致，项目根目录为当前运行目录
```

浏览器访问
[http://localhost:8181](