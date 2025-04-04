package dependencies

import (
	"Streaming/src/peliculas/application"
	"Streaming/src/peliculas/infrastructure/controllers"
	infrastructure "Streaming/src/peliculas/infrastructure/database"
)

type PeliculaDependencies struct {
	CreatePeliculaController       *controllers.CreatePeliculaController
	UpdatePeliculaController       *controllers.UpdatePeliculaController
	DeletePeliculaController       *controllers.DeletePeliculaController
	ListPeliculasController        *controllers.ListPeliculasController
	ListPeliculaByIDController     *controllers.ListPeliculaByIDController
	ListPeliculasByGenreController *controllers.ListPeliculasByGenreController
}

func InitPeliculaDependencies() *PeliculaDependencies {
	repo := infrastructure.NewMySQLRepositoryPeliculas()

	return &PeliculaDependencies{
		CreatePeliculaController:       controllers.NewCreatePeliculaController(application.NewCreatePelicula(repo)),
		UpdatePeliculaController:       controllers.NewUpdatePeliculaController(application.NewUpdatePelicula(repo)),
		DeletePeliculaController:       controllers.NewDeletePeliculaController(application.NewDeletePelicula(repo)),
		ListPeliculasController:        controllers.NewListPeliculasController(application.NewGetAllPeliculas(repo)),
		ListPeliculaByIDController:     controllers.NewListPeliculaByIDController(application.NewFindPeliculaByID(repo)),
		ListPeliculasByGenreController: controllers.NewListPeliculasByGenreController(application.NewListPeliculasByGenre(repo)),
	}
}
