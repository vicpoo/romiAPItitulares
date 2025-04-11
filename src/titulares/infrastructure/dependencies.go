// infrastructure/dependencies.go
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
	*SendEmailController,
) {

	repo := NewMysqlTitularRepository()

	createUseCase := application.NewCreateTitularUseCase(repo)
	viewUseCase := application.NewViewTitular(repo)
	updateUseCase := application.NewUpdateTitular(repo)
	deleteUseCase := application.NewDeleteTitularUseCase(repo)
	viewAllUseCase := application.NewViewAlltitulares(repo)
	sendEmailUseCase := application.NewSendEmailUseCase(repo)

	createController := NewCreateTitularController(createUseCase)
	viewController := NewViewTitularController(viewUseCase)
	updateController := NewUpdateTitularController(updateUseCase)
	deleteController := NewDeleteTitularController(deleteUseCase)
	viewAllController := NewViewAllTitularesController(viewAllUseCase)
	sendEmailController := NewSendEmailController(sendEmailUseCase)

	return createController, viewController, updateController, deleteController, viewAllController, sendEmailController
}
