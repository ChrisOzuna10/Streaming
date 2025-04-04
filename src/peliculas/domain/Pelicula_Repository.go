package domain

import "Streaming/src/peliculas/domain/entities"

type IPelicula interface {
	Save(pelicula entities.Pelicula) (entities.Pelicula, error)
	FindAll() ([]entities.Pelicula, error)
	FindByID(id uint) (entities.Pelicula, error)
	FindByGenre(genre string) ([]entities.Pelicula, error) // 👈 Esto estaba faltando
	FindByTitle(title string) ([]entities.Pelicula, error) // 👈 Si usás BuscarPelicula también
	Update(pelicula entities.Pelicula) (entities.Pelicula, error)
	Delete(id int) error
}
