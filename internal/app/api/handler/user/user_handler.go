package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josue/chalenge_golang/internal/app/api/dto"
)

type userAPIService interface {
	CreateUser(input dto.UserDTO) error
	LogIn(user, password string) (string, error)
}

type userHandler struct {
	userService userAPIService
}

func NewUserHandler(userService userAPIService) userHandler {
	return userHandler{
		userService: userService,
	}
}

// Create godoc
//
//	@Summary        Create new users
//	@Description    Add a new User
//	@Tags           User
//	@Accept         json
//	@Param          request body    dto.UserDTO  true    "user info"
//	@Success        201
//	@Failure        400
//	@Failure        500
//	@Router         /user [post]
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

// Create godoc
//
//	@Summary        Login
//	@Description    Loging and generate token
//	@Tags           User
//	@Accept         json
//	@Param          request body    dto.UserLogIn  true    "user login info"
//	@Success        202
//	@Failure        400
//	@Failure        500
//	@Router         /user/login [post]
func (uh userHandler) LogIn(c *gin.Context) {
	log.Println("On userHandler.LogIn")

	var input dto.UserLogIn
	if err := c.BindJSON(&input); err != nil {
		log.Println("On userHandler.LogIn - error: ", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err := input.Validate(); err != nil {
		log.Println("On userHandler.LogIn - Invalid params: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	token, err := uh.userService.LogIn(input.User, input.Pass)
	if err != nil {
		log.Println("On userHandler.LogIn - error: ", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}
