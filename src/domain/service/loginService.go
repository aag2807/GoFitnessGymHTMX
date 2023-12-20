package service

import "github.com/GoGym/src/boundary/persistance"

type LoginService struct {
	userRepository *persistance.UserRepository
}

func NewLoginService() *LoginService {
	return &LoginService{
		userRepository: persistance.NewUserRepository(),
	}
}

func (ls *LoginService) Login(email string, password string) (bool, error) {
	user, err := ls.userRepository.GetUserByEmail(email)
	if err != nil {
		return false, err
	}

	if user.Password != password {
		panic("Invalid Credentials")
	}

	return true, nil
}
