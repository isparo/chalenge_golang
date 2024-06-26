package user

type UserRepository interface {
	SaveUser(user User) error
	ValidateUserUniqueInfo(phone string, email string) (bool, error)
	GetUserInfoByLogin(user, password string) (*User, error)
}
