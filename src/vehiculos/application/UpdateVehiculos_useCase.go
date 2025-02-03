package application

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain/entities"
)

type UpdateVehiculo struct {
	db domain.IVehiculo
}

func NewUpdateVehiculo(db domain.IVehiculo) *UpdateVehiculo {
	return &UpdateVehiculo{db: db}
}

func (uc *UpdateVehiculo) Execute(id int, vehiculo entities.Vehiculo) error {
	return uc.db.Update(id, vehiculo)
}
