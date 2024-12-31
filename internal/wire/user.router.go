//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/vucong2018/study-go/internal/controller"
	"github.com/vucong2018/study-go/internal/repo"
	"github.com/vucong2018/study-go/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}
