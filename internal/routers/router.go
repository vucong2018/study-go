package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	c "github.com/vucong2018/study-go/internal/controller"
	"github.com/vucong2018/study-go/internal/middlewares"
)

func AA() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("BEFORE -->> AA")
		ctx.Next()
		fmt.Println("AFTER -->> AA")
	}
}
func BB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("BEFORE -->> BB")
		ctx.Next()
		fmt.Println("AFTER -->> BB")
	}
}

func CC(ctx *gin.Context) {
	fmt.Println("BEFORE -->> CC")
	ctx.Next()
	fmt.Println("AFTER -->> CC")
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthenMiddleware(), BB(), CC)
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUserById)
	}
	return r
}
