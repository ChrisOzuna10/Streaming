package application

import (
	"Streaming/src/peliculas/domain"
	"Streaming/src/peliculas/domain/entities"
	"fmt"
	"log"
)

type UpdatePelicula struct {
	repo domain.IPelicula
}

func NewUpdatePelicula(repo domain.IPelicula) *UpdatePelicula {
	return &UpdatePelicula{repo: repo}
}

func (up *UpdatePelicula) Execute(pelicula entities.Pelicula) (entities.Pelicula, error) {
	log.Printf("[UpdatePelicula] Iniciando actualización de la película con ID %d...\n", pelicula.ID)

	existing, err := up.repo.FindByID(pelicula.ID)
	if err != nil {
		log.Printf("[UpdatePelicula] Error al buscar la película con ID %d: %v\n", pelicula.ID, err)
		return entities.Pelicula{}, err
	}

	if existing.ID == 0 {
		log.Printf("[UpdatePelicula] Película con ID %d no encontrada.\n", pelicula.ID)
		return entities.Pelicula{}, fmt.Errorf("película no encontrada")
	}

	updated, err := up.repo.Update(pelicula)
	if err != nil {
		log.Printf("[UpdatePelicula] Error al actualizar la película con ID %d: %v\n", pelicula.ID, err)
		return entities.Pelicula{}, err
	}

	log.Printf("[UpdatePelicula] Película con ID %d actualizada correctamente.\n", pelicula.ID)
	return updated, nil
}
