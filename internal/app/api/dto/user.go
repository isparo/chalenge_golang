package dto

import (
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
			validation.Length(6, 12),
			validation.Match(regexp.MustCompile(`[A-Z]`)).Error("must contain at least one uppercase letter"),
			validation.Match(regexp.MustCompile(`[a-z]`)).Error("must contain at least one lowercase letter"),
			validation.Match(regexp.MustCompile(`[0-9]`)).Error("must contain at least one number"),
			validation.Match(regexp.MustCompile(`[@$&]`)).Error("must contain at least one special character (@, $, or &)")),
	)
}
