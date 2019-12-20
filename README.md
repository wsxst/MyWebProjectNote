# MyWebProjectNote
web全栈学习笔记
## 前端
前端使用vue，可以用vue-cli脚手架生成最简单的一个前端项目框架。
```
vue create frontend 
```
安装Element-UI插件。
```
vue add element
```
为了实现前后端的交互，安装axios并进行引入。
```
import axios from "axios";
const instance = axios.create({
  baseURL: "http://127.0.0.1:3000"
});
```
这样就可以实现跨域+隐藏主机地址的效果。
编译+运行：
```
npm run serve
```
## 后端
后端使用Golang基于Gin框架进行实现，基于vendor进行第三方包管理。
为了解决跨域的问题，采用**cors**中间件，默认设置为允许所有跨域请求。
```
r := gin.Default()
r.Use(cors.Default())
```
编译并运行程序：
```
go build
go_learn.exe
```
