package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vucong2018/study-go/internal/wire"
)

type UserRouter struct{}

func (pr *UserRouter) InitRouterUser(Router *gin.RouterGroup) {
	// public router

	// Non dependendcy
	// ur := repo.NewUserRepository()
	// us := service.NewUserService(ur)
	// userHandlerNondependency := controller.NewUserController(us)

	userController, _ := wire.InitUserRouterHandler()
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}
	//private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())

	{
		userRouterPrivate.GET("/info")
	}
}
