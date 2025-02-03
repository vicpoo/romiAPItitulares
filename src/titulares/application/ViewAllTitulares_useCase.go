package application

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain/entities"
)

type ViewAllTitular struct {
	db domain.ITitulares
}

func NewViewAlltitulares(db domain.ITitulares) *ViewAllTitular {
	return &ViewAllTitular{db: db}
}

func (uc *ViewAllTitular) Execute() ([]entities.Titular, error) {
	return uc.db.GetAll()
}
