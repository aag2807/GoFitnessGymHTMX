package persistance

import (
	entities "github.com/GoGym/src/domain/entities/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbContext struct {
	Db *gorm.DB
}

func NewDbContext() *DbContext {
	dsn := "root:root@tcp(127.0.0.1:3306)/GoFitnessGym?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	migrateEntities(db)

	if err != nil {
		panic("failed to connect database")
	}

	return &DbContext{
		Db: db,
	}
}

func migrateEntities(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
}
