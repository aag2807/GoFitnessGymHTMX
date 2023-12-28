package service

import (
	"github.com/GoGym/src/boundary/persistance"
	entities "github.com/GoGym/src/domain/entities/user"
	lib "github.com/aag2807/triplex-to-go"
)

type UserService struct {
	userRepository *persistance.UserRepository
	state          lib.State
	arguments      lib.Arguments
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: persistance.NewUserRepository(),
		state:          lib.State{},
		arguments:      lib.Arguments{},
	}
}

func (ls *UserService) GetUserByID(id int) (entities.User, error) {
	return ls.userRepository.GetUserByID(id)
}
