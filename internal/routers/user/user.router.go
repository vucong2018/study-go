package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vucong2018/study-go/internal/controller"
	"github.com/vucong2018/study-go/internal/repo"
	"github.com/vucong2018/study-go/internal/service"
)

type UserRouter struct{}

func (pr *UserRouter) InitRouterUser(Router *gin.RouterGroup) {
	// public router

	// Non dependendcy
	ur := repo.NewUserRepository()
	us := service.NewUserService(ur)
	userHandlerNondependency := controller.NewUserController(us)
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userHandlerNondependency.Register)
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
