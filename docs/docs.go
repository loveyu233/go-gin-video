// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "loveyu233"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/captcha/get": {
            "get": {
                "tags": [
                    "captcha"
                ],
                "summary": "获取滑块验证码",
                "responses": {}
            }
        },
        "/captcha/ver": {
            "get": {
                "tags": [
                    "captcha"
                ],
                "summary": "验证滑块验证码",
                "responses": {}
            }
        },
        "/collec/add": {
            "get": {
                "tags": [
                    "collec"
                ],
                "summary": "添加收藏该视频",
                "responses": {}
            }
        },
        "/collec/count": {
            "get": {
                "tags": [
                    "collec"
                ],
                "summary": "统计该视频的收藏量",
                "responses": {}
            }
        },
        "/collec/del": {
            "get": {
                "tags": [
                    "collec"
                ],
                "summary": "删除收藏该视频",
                "responses": {}
            }
        },
        "/collec/is": {
            "get": {
                "tags": [
                    "collec"
                ],
                "summary": "判断用户是否收藏该视频",
                "responses": {}
            }
        },
        "/collec/usercollec": {
            "get": {
                "tags": [
                    "collec"
                ],
                "summary": "获取用户收藏的视频",
                "responses": {}
            }
        },
        "/comment/byid": {
            "get": {
                "tags": [
                    "comment"
                ],
                "summary": "通过评论id获取评论",
                "responses": {}
            }
        },
        "/comment/delete": {
            "get": {
                "tags": [
                    "comment"
                ],
                "summary": "删除评论",
                "responses": {}
            }
        },
        "/comment/get": {
            "get": {
                "tags": [
                    "comment"
                ],
                "summary": "获取评论信息",
                "responses": {}
            }
        },
        "/comment/insert": {
            "post": {
                "tags": [
                    "comment"
                ],
                "summary": "添加评论",
                "responses": {}
            }
        },
        "/follow/add": {
            "get": {
                "tags": [
                    "follow"
                ],
                "summary": "添加关注",
                "responses": {}
            }
        },
        "/follow/del": {
            "get": {
                "tags": [
                    "follow"
                ],
                "summary": "取消关注",
                "responses": {}
            }
        },
        "/follow/followmecount": {
            "get": {
                "tags": [
                    "follow"
                ],
                "summary": "查询被关注数量",
                "responses": {}
            }
        },
        "/follow/followmeuser": {
            "get": {
                "tags": [
                    "follow"
                ],
                "summary": "查询被关注的用户",
                "responses": {}
            }
        },
        "/follow/is": {
            "get": {
                "tags": [
                    "follow"
                ],
                "summary": "查询该是否关注了该用户",
                "responses": {}
            }
        },
        "/follow/mefollowcount": {
            "get": {
                "tags": [
                    "follow"
                ],
                "summary": "查询关注数量",
                "responses": {}
            }
        },
        "/follow/mefollowuser": {
            "get": {
                "tags": [
                    "follow"
                ],
                "summary": "查询被关注用户",
                "responses": {}
            }
        },
        "/like/add": {
            "get": {
                "tags": [
                    "like"
                ],
                "summary": "添加点赞",
                "responses": {}
            }
        },
        "/like/count": {
            "get": {
                "tags": [
                    "like"
                ],
                "summary": "获取视频点赞个数",
                "responses": {}
            }
        },
        "/like/del": {
            "get": {
                "tags": [
                    "like"
                ],
                "summary": "删除点赞",
                "responses": {}
            }
        },
        "/like/is": {
            "get": {
                "tags": [
                    "like"
                ],
                "summary": "判断用户是否点赞该视频",
                "responses": {}
            }
        },
        "/reply/add": {
            "post": {
                "tags": [
                    "comment"
                ],
                "summary": "添加回复",
                "responses": {}
            }
        },
        "/reply/delete": {
            "get": {
                "tags": [
                    "comment"
                ],
                "summary": "删除回复信息",
                "responses": {}
            }
        },
        "/root/video/check": {
            "get": {
                "tags": [
                    "video"
                ],
                "summary": "修改视频状态",
                "responses": {}
            }
        },
        "/root/video/status": {
            "get": {
                "tags": [
                    "video"
                ],
                "summary": "查询用户所有未过审视频,用于root用户查询用来审核视频使用",
                "responses": {}
            }
        },
        "/upload/video": {
            "post": {
                "tags": [
                    "upload"
                ],
                "summary": "上传视频",
                "responses": {}
            }
        },
        "/user/authcode": {
            "get": {
                "tags": [
                    "user"
                ],
                "summary": "用户邮箱验证码发送",
                "responses": {}
            }
        },
        "/user/byid": {
            "get": {
                "tags": [
                    "user"
                ],
                "summary": "通过userid获取用户信息",
                "responses": {}
            }
        },
        "/user/login": {
            "post": {
                "tags": [
                    "user"
                ],
                "summary": "用户登录",
                "responses": {}
            }
        },
        "/user/register": {
            "post": {
                "tags": [
                    "user"
                ],
                "summary": "用户注册",
                "responses": {}
            }
        },
        "/user/resetpassword": {
            "post": {
                "tags": [
                    "user"
                ],
                "summary": "重置密码",
                "responses": {}
            }
        },
        "/user/update": {
            "post": {
                "tags": [
                    "user"
                ],
                "summary": "修改用户信息",
                "responses": {}
            }
        },
        "/video/checkfalse": {
            "get": {
                "tags": [
                    "video"
                ],
                "summary": "查询用户所有未过审视频,用于用于自己查询自己的视频",
                "responses": {}
            }
        },
        "/video/get": {
            "get": {
                "tags": [
                    "video"
                ],
                "summary": "根据视频id获取视频信息",
                "responses": {}
            }
        },
        "/video/list": {
            "get": {
                "tags": [
                    "video"
                ],
                "summary": "分页查询视频",
                "responses": {}
            }
        },
        "/video/search": {
            "get": {
                "tags": [
                    "video"
                ],
                "summary": "关键字查询视频",
                "responses": {}
            }
        },
        "/video/user": {
            "get": {
                "tags": [
                    "video"
                ],
                "summary": "通过用户id查询视频",
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://127.0.0.1:9898",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "go-gin-videoSwaggerApi",
	Description:      "go-gin-videoSwaggerApi",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
