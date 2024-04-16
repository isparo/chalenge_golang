package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	user_handler "github.com/josue/chalenge_golang/internal/app/api/handler/user"
	user_service "github.com/josue/chalenge_golang/internal/app/api/service/user"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
}

type apiV1 struct {
	userHandler UserHandler
}

func newApiV1(userHandler UserHandler) apiV1 {
	return apiV1{
		userHandler: userHandler,
	}
}

func LoadApiV1() {
	glog.Info("loading API v1")

	userService := user_service.NewUserAPIService()
	userHandler := user_handler.NewUserHandler(userService)

	api := newApiV1(userHandler)
	api.loadRoutes()
}
