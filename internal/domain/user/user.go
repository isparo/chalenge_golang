package user

import uuid "github.com/satori/go.uuid"

type User struct {
	ID          *uuid.UUID
	UserName    string
	Email       string
	PhoneNumber string
	Pass        string
}

func NewUser(userName string, email string, phone string, password string) User {
	return User{
		UserName:    userName,
		Email:       email,
		PhoneNumber: phone,
		Pass:        password,
	}
}
