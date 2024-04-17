package api

import (
	"log"

	"github.com/gin-gonic/gin"
	handler "github.com/josue/chalenge_golang/internal/app/api/handler"
	service "github.com/josue/chalenge_golang/internal/app/api/service"
	"github.com/josue/chalenge_golang/internal/domain/user"
	"github.com/josue/chalenge_golang/internal/infrastructure/database"
	"github.com/josue/chalenge_golang/internal/infrastructure/persistency"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
	LogIn(c *gin.Context)
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
	log.Println("loading API v1")

	dbConfig, err := database.NewConfig()
	if err != nil {
		log.Fatal("error loading configuration: ", err.Error())
	}

	dbConnection, err := database.NewDatabseConnection(dbConfig)
	if err != nil {
		log.Fatal("error connecting to database: ", err.Error())
	}

	userRepository := persistency.NewUserMySQLRepository(dbConnection)

	userDomainService := user.NewUserService(userRepository)

	userService := service.NewUserAPIService(userDomainService)
	userHandler := handler.NewUserHandler(userService)

	api := newApiV1(userHandler)
	api.loadRoutes()
}
