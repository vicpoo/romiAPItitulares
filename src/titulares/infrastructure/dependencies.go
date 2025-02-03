package infrastructure

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/application"
)

func InitTitularDependencies() (
	*CreateTitularController,
	*ViewTitularController,
	*UpdateTitularController,
	*DeleteTitularController,
	*ViewAllTitularesController,
) {

	repo := NewMysqlTitularRepository()

	createUseCase := application.NewCreateTitularUseCase(repo)
	viewUseCase := application.NewViewTitular(repo)
	updateUseCase := application.NewUpdateTitular(repo)
	deleteUseCase := application.NewDeleteTitularUseCase(repo)
	viewAllUseCase := application.NewViewAlltitulares(repo)

	createController := NewCreateTitularController(createUseCase)
	viewController := NewViewTitularController(viewUseCase)
	updateController := NewUpdateTitularController(updateUseCase)
	deleteController := NewDeleteTitularController(deleteUseCase)
	viewAllController := NewViewAllTitularesController(viewAllUseCase)

	return createController, viewController, updateController, deleteController, viewAllController
}
