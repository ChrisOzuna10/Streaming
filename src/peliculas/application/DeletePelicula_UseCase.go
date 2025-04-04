package application

import (
	"Streaming/src/peliculas/domain"
	"log"
)

type DeletePelicula struct {
	repo domain.IPelicula
}

func NewDeletePelicula(repo domain.IPelicula) *DeletePelicula {
	return &DeletePelicula{repo: repo}
}

func (dp *DeletePelicula) Execute(peliculaID int) error {
	log.Println("[DeletePelicula] Iniciando eliminación de película...")

	err := dp.repo.Delete(peliculaID)
	if err != nil {
		log.Printf("[DeletePelicula] Error al eliminar película con ID %d: %v\n", peliculaID, err)
		return err
	}

	log.Println("[DeletePelicula] Película eliminada exitosamente.")
	return nil
}
