package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/vucong2018/study-go/global"
	"github.com/vucong2018/study-go/internal/routers"
)

func InitRouter() *gin.Engine {
	// r := gin.Default()
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middlewares
	// r.Use() // logging
	// r.Use() //cross
	// r.Use() // limmiter global
	manageRouter := routers.RouterGroupApp.Manage
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/check-status")
	}
	{
		userRouter.InitRouterUser(MainGroup)
		userRouter.IntProductRouter(MainGroup)
	}
	{
		manageRouter.InitRouterAdmin(MainGroup)
		manageRouter.InitRouterUser(MainGroup)
	}
	return r
}
