# Persimmon Blog

![macbook](https://raw.githubusercontent.com/cong5/myPersimmon/master/screen.jpg)


## 运行环境

- Nginx 1.8+ (Nginx reverse proxy)
- Golang 1.10+
- MySQL 5.5+
- Redis 3.0+

## 首先设置 Golang 的运行环境

```
export GOPATH=$HOME/go
export PATH=$HOME/bin:$GOPATH/bin:$PATH
```

## 安装 Golang 依赖管理工具 

Glide: Golang 依赖管理工具  [https://glide.sh](https://github.com/Masterminds/glide)

安装

```
curl https://glide.sh/get | sh
```

## 基本安装

```
go get github.com/cong5/persimmon
```

## 安装依赖

```
cd $GOPATH/src/github.com/cong5/persimmon
glide install
```

## 编辑配置文件

把配置文件 `conf/app.conf.example` 复制一份，修改文件名为 `conf/app.conf`

修改 MySQL 数据库配置信息、七牛配置信息、百度翻译、邮件发送等的配置信息。

## 运行应用

```
cd $GOPATH && revel run github.com/cong5/persimmon
```

## 编译二进制文件

把 persimmon 编译成二进制可执行文件

```
cd $GOPATH/src/github.com/cong5/persimmon
go build -o ./bin/persimmon github.com/cong5/persimmon/app/tmp
```

## 运行 persimmon

必选参数: `-importPath`

```
./bin/persimmon -importPath=github.com/cong5/persimmon -runMode=prod -port=9100
```

更多的信息请查看： `https://revel.github.io/manual/tool.html`

生产环境部署，直接跨平台编译

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 revel package github.com/cong5/persimmon
```
编译后得到 `persimmon.tar.gz` 的压缩包文件，把这个压缩包上传到服务器，解压出来就可以运行了。


## 后端 

后台地址

http://example.com/backend

默认用户名和密码.

用户名: persimmon@cong5.net

密码: Persimmon2018



### 前端开发

> 后台模块使用 ES6 +  VueJS 框架 + IViewUI 框架
> 前台模块使用 ES5 + VueJS 框架

1). 安装 node.js

到 [https://nodejs.org/en/](https://nodejs.org/en/) 下载最新发行的版本.

2). npm install

```shell
npm install
```

如果在国内，可以使用淘宝的镜像 [Taobao NPM mirror:http://npm.taobao.org/](http://npm.taobao.org/)

3). 开发模式

首先 `cd` 到项目路径.

```
cd $GOPATH/src/github.com/cong5/persimmon
```

后台模块
```
# run dev
npm run backend-watch
```

前台模块

```
# run dev
npm run home-watch
```

4). 生产模式

后台模块

```
# run prod
npm run backend-production
```

前台模块

```
# run prod
npm run home-production
```

## 使用到的第三方服务

文件和图片云储存，[七牛](https://www.qiniu.com/) .

文章标题转英文，使用 [百度翻译](https://api.fanyi.baidu.com/api/trans/product/index) .

垃圾评论检测 [Akismet](https://akismet.com) .  
