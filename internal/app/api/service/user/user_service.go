package service

import (
	"errors"
	"log"

	"github.com/josue/chalenge_golang/internal/app/api/dto"
	"github.com/josue/chalenge_golang/internal/domain/user"
)

type userDomainService interface {
	CreateUser(userData user.User) error
	ValidateUserByPhoneAndEmail(phone string, email string) (bool, error)
	GetUserInfoByLogin(user, password string) (*user.User, error)
}

type userAPIService struct {
	userService userDomainService
}

func NewUserAPIService(userService userDomainService) userAPIService {
	return userAPIService{
		userService: userService,
	}
}

func (us userAPIService) CreateUser(input dto.UserDTO) error {
	log.Println("On userAPIService.CreateUser")

	isValid, err := us.userService.ValidateUserByPhoneAndEmail(input.PhoneNumber, input.Email)
	if err != nil {
		log.Println("On userAPIService.CreateUser - Error validation user email and Phone number: ", err.Error())
		return err
	}

	if !isValid {
		log.Println("On userAPIService.CreateUser -  User already registered")
		return errors.New("user already registered")
	}

	user := user.NewUser(input.UserName, input.Email, input.PhoneNumber, input.Password)
	if err := us.userService.CreateUser(user); err != nil {
		log.Println("On userAPIService.CreateUser - error creating user: ", err.Error())
		return err
	}

	return nil
}

func (us userAPIService) LogIn(user, password string) (string, error) {
	log.Println("On userAPIService.LogIn")

	// generar jwt
	userInfo, err := us.userService.GetUserInfoByLogin(user, password)
	if err != nil {
		log.Println("On userAPIService.LogIn - error: ", err.Error())
		return "", err
	}

	log.Println(userInfo)

	return "", nil
}
