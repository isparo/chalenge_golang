package service

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/josue/chalenge_golang/internal/app/api/dto"
	"github.com/josue/chalenge_golang/internal/domain/user"
)

// this one could be on a configuration and taken from the EnvVars
const secret string = "jwt_secret"

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

	userInfo, err := us.userService.GetUserInfoByLogin(user, password)
	if err != nil {
		log.Println("On userAPIService.LogIn - error: ", err.Error())
		return "", err
	}

	token, err := createToken(userInfo.ID.String(), userInfo.UserName, userInfo.Email, userInfo.PhoneNumber)
	if err != nil {
		log.Println("On userAPIService.LogIn - error creating token: ", err.Error())
		return "", err
	}

	return *token, nil
}

func createToken(id, userName, email, phoneNumber string) (*string, error) {
	var jwtKey = []byte(secret)

	id = strings.TrimSpace(id)
	userName = strings.TrimSpace(userName)
	email = strings.TrimSpace(email)
	phoneNumber = strings.TrimSpace(phoneNumber)

	claims := jwt.MapClaims{
		"id":           id,
		"username":     userName,
		"email":        email,
		"phone_number": phoneNumber,
		"exp":          time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
