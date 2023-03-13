# 项目介绍

## 项目目录
> http文件夹全部的api测试 
> 
>docs文件夹swagger文档
````
go-gin-video
|-- api
|-- config
|-- docs
|-- domain
|-- http
|-- initialization
|-- logs
|-- middleware
|-- router
|-- service
|   |-- mongoDBService
|   |-- mysqlService
|   |-- redisService
|-- static
|   |-- code
|   |-- resources
|       |-- images
|       |-- video
|-- utils
````
## 项目实用技术:
>Go,Gin,Mysql,Redis,MongoDB

## 项目实现的功能有:
> 
> 1. 用户模块
>    1. 登录,注册,找回,修改
> 2. 视频模块
>    1. 视频上传
>    2. 视频信息校验
>    3. ffmpeg调整视频清晰度和对视频切片
> 3. 评论回复模块
>    1. MongoDB保存用户的评论和回复
>    2. 查询,保存,删除
> 4. 点赞模块
>    1. 使用redis的zset数据结构对对点赞数据进行保存和统计
>    2. 实现功能有:
>       1. 判断是否点赞
>       2. 点赞,取消点赞,统计点赞数
> 5. 关注模块
>    1. 判断是否关注
>    2. 关注,取消关注
>    3. 查询被关注的数量和被关注的用户
>    4. 查询关注的数量和关注的用户
> 6. 收藏模块
>    1. 判断是否收藏
>    2. 收藏,取消收藏
>    3. 查询收藏
> 7. 人机交互模块
>    1. 图形验证码