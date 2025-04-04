package infrastructure

import (
	"Streaming/src/core"
	"Streaming/src/users/domain"
	"Streaming/src/users/domain/entities"
)

type MySQLUserRepository struct{}

func NewMySQLUserRepository() domain.IUser {
	return &MySQLUserRepository{}
}

func (repo *MySQLUserRepository) Save(user entities.User) (entities.User, error) {
	result := core.BD.Create(&user)
	return user, result.Error
}

func (repo *MySQLUserRepository) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	result := core.BD.Where("email_user = ?", email).First(&user)
	return user, result.Error
}
