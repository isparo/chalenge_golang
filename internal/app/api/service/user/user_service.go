package service

import "github.com/golang/glog"

type userAPIService struct {
	//userService userDomainService
}

func NewUserAPIService() userAPIService {
	return userAPIService{
		// apiConfig:   apiConfig,
		// userService: userService,
	}
}

func (us userAPIService) CreateUser() error {
	glog.Info("On userAPIService.CreateUser")

	return nil
}
