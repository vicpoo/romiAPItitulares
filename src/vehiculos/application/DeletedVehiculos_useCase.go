package application

import "github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain"

type DeletedVehiculoUseCase struct {
	db domain.IVehiculo
}

func NewDeletedVehiculoUseCase(db domain.IVehiculo) *DeletedVehiculoUseCase {
	return &DeletedVehiculoUseCase{
		db: db,
	}
}

func (uc *DeleteVehiculoUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
