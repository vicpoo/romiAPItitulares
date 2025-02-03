package application

import "github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain"

type DeleteVehiculoUseCase struct {
	db domain.IVehiculo
}

func NewDeleteVehiculoUseCase(db domain.IVehiculo) *DeleteVehiculoUseCase {
	return &DeleteVehiculoUseCase{
		db: db,
	}
}

func (uc *DeleteVehiculoUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
