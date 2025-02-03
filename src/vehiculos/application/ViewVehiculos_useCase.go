package application

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain/entities"
)

type ViewVehiculo struct {
	db domain.IVehiculo
}

func NewViewVehiculo(db domain.IVehiculo) *ViewVehiculo {
	return &ViewVehiculo{db: db}
}

func (uc *ViewVehiculo) Execute(id int) (entities.Vehiculo, error) {
	return uc.db.FindByID(id)
}
