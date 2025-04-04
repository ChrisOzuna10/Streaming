package infrastructure

import (
	"Streaming/src/core"
	"Streaming/src/peliculas/domain"
	"Streaming/src/peliculas/domain/entities"
	"fmt"
)

type MySQLRepositoryPeliculas struct{}

func NewMySQLRepositoryPeliculas() domain.IPelicula {
	return &MySQLRepositoryPeliculas{}
}

func (repo *MySQLRepositoryPeliculas) Save(pelicula entities.Pelicula) (entities.Pelicula, error) {
	result := core.BD.Create(&pelicula)
	return pelicula, result.Error
}

func (repo *MySQLRepositoryPeliculas) FindByID(id uint) (entities.Pelicula, error) {
	var pelicula entities.Pelicula
	result := core.BD.First(&pelicula, "id_pelicula = ?", id)
	if result.Error != nil {
		return entities.Pelicula{}, result.Error
	}
	if result.RowsAffected == 0 {
		fmt.Println("No se encontró ninguna película con id", id)
		return entities.Pelicula{}, fmt.Errorf("película no encontrada")
	}
	fmt.Println("Película encontrada correctamente", pelicula)
	return pelicula, nil
}

func (repo *MySQLRepositoryPeliculas) FindAll() ([]entities.Pelicula, error) {
	var peliculas []entities.Pelicula
	result := core.BD.Find(&peliculas)
	return peliculas, result.Error
}

func (repo *MySQLRepositoryPeliculas) FindByGenre(genre string) ([]entities.Pelicula, error) {
	var peliculas []entities.Pelicula
	result := core.BD.Where("genero = ?", genre).Find(&peliculas)
	return peliculas, result.Error
}

func (repo *MySQLRepositoryPeliculas) FindByTitle(title string) ([]entities.Pelicula, error) {
	var peliculas []entities.Pelicula
	result := core.BD.Where("titulo LIKE ?", "%"+title+"%").Find(&peliculas)
	return peliculas, result.Error
}

func (repo *MySQLRepositoryPeliculas) Update(pelicula entities.Pelicula) (entities.Pelicula, error) {
	result := core.BD.Model(&entities.Pelicula{}).
		Where("id_pelicula = ?", pelicula.ID).
		Updates(pelicula)
	return pelicula, result.Error
}

func (repo *MySQLRepositoryPeliculas) Delete(id int) error {
	result := core.BD.Where("id_pelicula = ?", id).Delete(&entities.Pelicula{})
	return result.Error
}
