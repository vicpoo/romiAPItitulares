package application

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain/entities"
)

type ViewAllVehiculos struct {
	db domain.IVehiculo
}

func NewViewAllvehiculos(db domain.IVehiculo) *ViewAllVehiculos {
	return &ViewAllVehiculos{db: db}
}

func (uc *ViewAllVehiculos) Execute() ([]entities.Vehiculo, error) {
	return uc.db.GetAll()
}
