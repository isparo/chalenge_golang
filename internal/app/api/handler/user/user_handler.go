package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josue/chalenge_golang/internal/app/api/dto"
)

type userAPIService interface {
	CreateUser(input dto.UserDTO) error
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
	log.Println("On userHandler.CreateUser")

	var input dto.UserDTO

	if err := c.BindJSON(&input); err != nil {
		log.Println("On userHandler.CreateUser - error: ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := input.Validate(); err != nil {
		log.Println("On userHandler.CreateUser - error creating user: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := uh.userService.CreateUser(input); err != nil {
		log.Println("On create user handler- error creating user: ", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
