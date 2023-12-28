package service

import (
	"github.com/GoGym/src/boundary/persistance"
	entities "github.com/GoGym/src/domain/entities/user"
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

func (ls *LoginService) GetUserByEmail(email string) (entities.User, error) {
	return ls.userRepository.GetUserByEmail(email)
}

func (ls *LoginService) GetUserByID(ID int) (entities.User, error) {
	return ls.userRepository.GetUserByID(ID)
}
