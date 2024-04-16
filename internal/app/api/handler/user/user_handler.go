package user

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// Interface to mock user API service
type userAPIService interface {
	CreateUser() error
}

type userHandler struct {
	userService userAPIService
}

func NewUserHandler(userService userAPIService) userHandler {
	return userHandler{
		userService: userService,
	}
}

func (uh userHandler) CreateUser(c *gin.Context) {
	glog.Info("On userHamdler.CreateUser")

	uh.userService.CreateUser()
}
