package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ar *AdminRouter) InitRouterAdmin(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// //private router
	adminRouterPrivate := Router.Group("/admin/admin")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active-user")
	}
}
