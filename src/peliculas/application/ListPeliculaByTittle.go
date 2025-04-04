// src/peliculas/application/ListPeliculasByGenre_UseCase.go
package application

import (
	"Streaming/src/peliculas/domain"
	"Streaming/src/peliculas/domain/entities"
)

type ListPeliculasByGenre struct {
	repo domain.IPelicula
}

func NewListPeliculasByGenre(repo domain.IPelicula) *ListPeliculasByGenre {
	return &ListPeliculasByGenre{repo: repo}
}

func (uc *ListPeliculasByGenre) Execute(genre string) ([]entities.Pelicula, error) {
	return uc.repo.FindByGenre(genre)
}
