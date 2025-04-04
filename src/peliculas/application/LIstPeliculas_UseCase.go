package application

import (
	"Streaming/src/peliculas/domain"
	"Streaming/src/peliculas/domain/entities"
	"log"
)

type GetAllPeliculas struct {
	repo domain.IPelicula
}

func NewGetAllPeliculas(repo domain.IPelicula) *GetAllPeliculas {
	return &GetAllPeliculas{repo: repo}
}

func (ga *GetAllPeliculas) Execute() ([]entities.Pelicula, error) {
	log.Println("[GetAllPeliculas] Recuperando todas las películas...")

	peliculas, err := ga.repo.FindAll()
	if err != nil {
		log.Printf("[GetAllPeliculas] Error al recuperar películas: %v\n", err)
		return nil, err
	}

	log.Printf("[GetAllPeliculas] %d películas encontradas.\n", len(peliculas))
	return peliculas, nil
}
