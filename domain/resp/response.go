package resp

import "github.com/gin-gonic/gin"

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// 成功响应
func OkResponse(c *gin.Context, data interface{}) {
	c.JSON(200, &Response{Data: data})
}

// 错误响应
func ErrorResponse(c *gin.Context, error string) {
	c.JSON(200, &Response{Error: error})
}
