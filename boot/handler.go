package boot

import "github.com/gin-gonic/gin"

// Handler 封装中间价接口
type Handler interface {
	OnRequest(ctx *gin.Context) error
}
