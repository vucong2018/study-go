package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/vucong2018/study-go/pkg/respone"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			respone.ErrorRespone(ctx, respone.ErrInvalidToken, "")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
