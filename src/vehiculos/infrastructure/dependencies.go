package infrastructure

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/application"
)

func InitVehiculoDependencies() (
	*CreateVehiculoController,
	*ViewVehiculoController,
	*UpdateVehiculoController,
	*DeleteVehiculoController,
	*ViewAllVehiculosController,
) {

	repo := NewMysqlVehiculoRepository()

	createUseCase := application.NewCreateVehiculoUseCase(repo)
	viewUseCase := application.NewViewVehiculo(repo)
	updateUseCase := application.NewUpdateVehiculo(repo)
	deleteUseCase := application.NewDeleteVehiculoUseCase(repo)
	viewAllUseCase := application.NewViewAllvehiculos(repo)

	createController := NewCreateVehiculoController(createUseCase)
	viewController := NewViewVehiculoController(viewUseCase)
	updateController := NewUpdateVehiculoController(updateUseCase)
	deleteController := NewDeleteVehiculoController(deleteUseCase)
	viewAllController := NewViewAllVehiculosController(viewAllUseCase)

	return createController, viewController, updateController, deleteController, viewAllController
}
