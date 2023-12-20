package service

import (
	"github.com/GoGym/src/boundary/persistance"
	lib "github.com/aag2807/triplex-to-go"
)

type LoginService struct {
	userRepository *persistance.UserRepository
	state          lib.State
	arguments      lib.Arguments
}

func NewLoginService() *LoginService {
	return &LoginService{
		userRepository: persistance.NewUserRepository(),
		state:          lib.State{},
		arguments:      lib.Arguments{},
	}
}

func (ls *LoginService) Login(email string, password string) (bool, error) {
	user, err := ls.userRepository.GetUserByEmail(email)
	if err != nil {
		return false, err
	}

	ls.state.IsTrue(user.Password == password, "Invalid Credentials")

	return true, nil
}
