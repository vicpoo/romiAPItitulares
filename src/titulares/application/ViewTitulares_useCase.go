package application

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain/entities"
)

type ViewTitular struct {
	db domain.ITitulares
}

func NewViewTitular(db domain.ITitulares) *ViewTitular {
	return &ViewTitular{db: db}
}

func (uc *ViewTitular) Execute(id int) (entities.Titular, error) {
	return uc.db.FindByID(id)
}
