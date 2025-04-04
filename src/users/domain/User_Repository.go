// src/users/domain/user_repository.go
package domain

import "Streaming/src/users/domain/entities"

type IUser interface {
	Save(user entities.User) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
}
