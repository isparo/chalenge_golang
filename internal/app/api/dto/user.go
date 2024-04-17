package dto

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserDTO struct {
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func NewUserDTO(userName string, email string, phoneNumber string, password string) UserDTO {
	return UserDTO{
		UserName:    userName,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    password,
	}
}

func (u UserDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.UserName, validation.Required, validation.Length(3, 20)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.PhoneNumber, validation.Required, validation.Length(10, 10)),
		validation.Field(&u.Password,
			validation.Required,
			validation.Length(6, 12)),
	)
}

// TODO: complete this function
func validatePassword(value interface{}) error {
	password := value.(string)
	// regex := `(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[@$&])^[\S]+$`
	// regex := `^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[@$&\w])[\S]+$`
	regex := `^(?=.*[A-Z])(?=.*[a-z])(?=.*[@$&\w])[\s\S]+$`

	if matched, _ := regexp.MatchString(regex, password); !matched {
		return errors.New("The password must contain at least one uppercase letter, one lowercase letter, one number, and one special character: @, $, or &")
	}

	return nil
}
