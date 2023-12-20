package persistance

import (
	entities "github.com/GoGym/src/domain/entities/user"
	lib "github.com/aag2807/triplex-to-go"
)

type UserRepository struct {
	Context   *DbContext
	State     lib.State
	Arguments lib.Arguments
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Context: NewDbContext(),
	}
}

// Get's a user by email
func (ur *UserRepository) GetUserByEmail(email string) (entities.User, error) {
	ur.Arguments.NotWhiteSpace(email, "Email cannot be empty")
	user := entities.User{}
	result := ur.Context.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return entities.User{}, result.Error
	}

	ur.Arguments.NotNil(user, "User not found")

	return user, nil
}
