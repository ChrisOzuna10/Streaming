package application

import (
	"Streaming/src/peliculas/domain"
	"Streaming/src/peliculas/domain/entities"
)

type FindPeliculaByID struct {
	repo domain.IPelicula
}

func NewFindPeliculaByID(repo domain.IPelicula) *FindPeliculaByID {
	return &FindPeliculaByID{repo: repo}
}

func (uc *FindPeliculaByID) Execute(id uint) (entities.Pelicula, error) {
	return uc.repo.FindByID(id)
}
