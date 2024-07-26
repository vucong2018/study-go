package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/vucong2018/study-go/internal/controller"
)


func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{	
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUserById)
	}
	return r
}

