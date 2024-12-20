package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitRouterUser(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin/user")
	{
		adminRouterPublic.POST("/login")
	}
	// //private router
	adminRouterPrivate := Router.Group("/admin/user")
	{
		adminRouterPrivate.POST("/active-user")
	}
}
