package application

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain/entities"
)

type UpdateTitular struct {
	db domain.ITitulares
}

func NewUpdateTitular(db domain.ITitulares) *UpdateTitular {
	return &UpdateTitular{db: db}
}

func (uc *UpdateTitular) Execute(id int, titular entities.Titular) error {
	return uc.db.Update(id, titular)
}
