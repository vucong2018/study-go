package service

import "github.com/vucong2018/study-go/internal/repo"

type UserService struct{
	userRepo *repo.UserRepo
}

func NewUserService () *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (ur *UserService) GetInfoUser() string {
	return ur.userRepo.GetInfoUser()
}