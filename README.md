# 项目描述 #
PineCMS是一个GO语言开发的内容管理系统, 让您可以在短时间内以制作模板的方式搭建出来一个网站, 非开发者也能快速愉悦的使用系统. 
简单使用情况下无需关注代码逻辑. 

![](./resources/assets/backend/static/images/1.png)

# 版本 #
项目目前处于打磨阶段, 每个可能会有小部分修改

# 下载部署 #
## 下载编译包 ##

git仓库: [pinecms](https://github.com/xiusin/pinecms/releases)

gitee仓库: [pinecms](https://gitee.net/xiusin/pinecms/releases)

解压下载包内容, 直接在目录下执行`pinecms(.exe)`

## 自行编译 ##
 ```
 git clone https://github.com/xiusin/pinecms.git
 cd pinecms
 go build -o pinecms
```

## 配置 ##
1. 数据库配置
    > 导入数据库结构`resources/pinecms.sql`
    >
    >修改`resources/configs/database.yml.dist`为`resources/configs/database.yml`
    

2. 安装依赖
    > `go build`

3. 运行项目
    > `./pinecms serve start` 

4. 开发期间自动构建
    > `go run main.go serve dev`

6. 访问后端登陆页面
    > 访问 `http://localhost:2019/b/login/index` 默认账号密码 `用户名: admin 密码: 123456`

# 自动静态化路由 #
完全自动静态文件和动态路由. 更友好的SEO方式

# 支持多主体一键切换 #
系统支持多种主题, 可自由切换不同风格的模板. 

# 内置标签系统 #
支持类似织梦的标签系统, 可以让您在不写任何代码的情况下完成网站建设. 

# 内置多种功能 #
- 个人信息
- 系统设置
- 内容管理
- 分类管理
- 模型管理
- 数据库管理
- 资源管理
- 管理员设置
- 广告管理
- 日志管理
- 友链管理


# 文档 #
[pinecms.xyz](http://pinecms.xyz/)

