package infrastructure

import (
	"Streaming/src/core"
	"Streaming/src/core/cache" // Asegúrate de importar el paquete de caché
	"Streaming/src/users/domain"
	"Streaming/src/users/domain/entities"
	"time" // Asegúrate de importar para trabajar con tiempo
)

type MySQLUserRepository struct{}

func NewMySQLUserRepository() domain.IUser {
	return &MySQLUserRepository{}
}

func (repo *MySQLUserRepository) Save(user entities.User) (entities.User, error) {
	// Guarda el usuario en la base de datos
	result := core.BD.Create(&user)

	// Limpiar el caché si se ha creado o actualizado un usuario
	cache.Cache.Delete(user.Email)
	return user, result.Error
}

func (repo *MySQLUserRepository) FindByEmail(email string) (entities.User, error) {
	// Verifica si el usuario está en el caché
	if cachedUser, found := cache.Cache.Get(email); found {
		return cachedUser.(entities.User), nil
	}

	// Si el usuario no está en caché, lo buscamos en la base de datos
	var user entities.User
	result := core.BD.Where("email_user = ?", email).First(&user)
	if result.Error != nil {
		return entities.User{}, result.Error
	}

	// Almacenar el usuario en el caché para futuras consultas
	// Puedes ajustar el tiempo de expiración como desees. Ejemplo: 1 hora de expiración.
	cache.Cache.Set(email, user, time.Hour) // o cache.NoExpiration para no expirar nunca

	return user, nil
}
