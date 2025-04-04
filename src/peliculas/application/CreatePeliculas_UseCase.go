package application

import (
	"Streaming/src/peliculas/domain"
	"Streaming/src/peliculas/domain/entities"
	"log"
)

type CreatePelicula struct {
	repo domain.IPelicula
}

func NewCreatePelicula(repo domain.IPelicula) *CreatePelicula {
	return &CreatePelicula{repo: repo}
}

func (cp *CreatePelicula) Execute(pelicula entities.Pelicula) (entities.Pelicula, error) {
	log.Println("[CreatePelicula] Iniciando proceso de guardado...")

	created, err := cp.repo.Save(pelicula)
	if err != nil {
		log.Printf("[CreatePelicula] Error al guardar la película: %v\n", err)
		return entities.Pelicula{}, err
	}

	log.Println("[CreatePelicula] Película guardada exitosamente.")
	return created, nil
}
