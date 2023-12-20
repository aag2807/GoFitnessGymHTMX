package persistance

import (
	entities "github.com/GoGym/src/domain/entities/user"
	lib "github.com/aag2807/triplex-to-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB        *gorm.DB
	State     lib.State
	Arguments lib.Arguments
}

func NewUserRepository() *UserRepository {
	dsn := "root:root@tcp(127.0.0.1:3306)/GoFitnessGym?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&entities.User{})
	if err != nil {
		panic("failed to connect database")
	}

	return &UserRepository{
		DB: db,
	}
}

// Get's a user by email
func (ur *UserRepository) GetUserByEmail(email string) (entities.User, error) {
	ur.Arguments.NotWhiteSpace(email, "Email cannot be empty")
	user := entities.User{}

	result := ur.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return entities.User{}, result.Error
	}

	ur.Arguments.NotNil(user, "User not found")

	return user, nil
}
