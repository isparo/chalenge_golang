package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserLogIn struct {
	User string `json:"user"`
	Pass string `json:"password"`
}

func (u UserLogIn) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.User, validation.Required),
		validation.Field(&u.Pass, validation.Required),
	)
}
