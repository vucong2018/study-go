// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/vucong2018/study-go/internal/controller"
	"github.com/vucong2018/study-go/internal/repo"
	"github.com/vucong2018/study-go/internal/service"
)

// Injectors from user.router.go:

func InitUserRouterHandler() (*controller.UserController, error) {
	iUserRepository := repo.NewUserRepository()
	iUserService := service.NewUserService(iUserRepository)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
