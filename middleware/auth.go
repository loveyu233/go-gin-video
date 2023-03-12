package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin-video/domain/resp"
	"go-gin-video/initialization"
	"go-gin-video/utils/jwt"
	"strconv"
)

// 认证
func Auth(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		resp.ErrorResponse(c, "无权访问")
		c.Abort()
		return
	}
	authorization = authorization[7:]

	userid, roleid, err := jwt.ParseToken(authorization)
	if err != nil {
		resp.ErrorResponse(c, "token解析失败")
		c.Abort()
		return
	}
	sub := initialization.Role(int(roleid))
	obj := c.Request.URL.Path
	act := c.Request.Method
	if !initialization.Check(sub, obj, act) {
		resp.ErrorResponse(c, "无权访问")
		c.Abort()
		return
	}
	c.Set("userid", strconv.Itoa(int(userid)))
}
