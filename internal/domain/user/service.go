package user

import "log"

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) userService {
	return userService{
		userRepo: userRepo,
	}
}

func (us userService) CreateUser(userData User) error {
	log.Println("On domain.service.CreateUser service")

	if err := us.userRepo.SaveUser(userData); err != nil {
		log.Println("On domain.service.CreateUser service - error creating user")
		return err
	}

	return nil
}

func (us userService) ValidateUserByPhoneAndEmail(phone string, email string) (bool, error) {
	log.Println("On domain.service.ValidateUserByPhoneAndEmail service")

	isValid, err := us.userRepo.ValidateUserUniqueInfo(phone, email)
	if err != nil {
		log.Println("On domain.service.ValidateUserByPhoneAndEmail service. Error: ", err.Error())
		return false, err
	}

	return isValid, nil
}
