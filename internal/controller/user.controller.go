package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vucong2018/study-go/internal/service"
	"github.com/vucong2018/study-go/pkg/respone"
)

// type UserController struct {
// 	userService *service.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// func (uc *UserController) GetUserById(c *gin.Context) {
// 	respone.SuccessRespone(c, 20001, []string{"tiphs", "vucong","haha"})
// }

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	respone.SuccessRespone(c, result, nil)
}
