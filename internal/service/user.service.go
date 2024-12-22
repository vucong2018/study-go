package service

import (
	"github.com/vucong2018/study-go/internal/repo"
	"github.com/vucong2018/study-go/pkg/respone"
)

// import "github.com/vucong2018/study-go/internal/repo"

// type UserService struct{
// 	userRepo *repo.UserRepo
// }

// func NewUserService () *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRepo(),
// 	}
// }

// func (ur *UserService) GetInfoUser() string {
// 	return ur.userRepo.GetInfoUser()
// }

// INTERFACE_VERSION

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	if us.userRepo.GetUserByEmail(email) {
		return respone.ErrCodeUserHasExist
	}
	return respone.ErrCodeSuccess
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}
