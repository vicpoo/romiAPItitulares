package application

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain/entities"
)

type CreateTitularUseCase struct {
	db domain.ITitulares
}

func NewCreateTitularUseCase(db domain.ITitulares) *CreateTitularUseCase {
	return &CreateTitularUseCase{
		db: db,
	}
}

func (uc *CreateTitularUseCase) Run(titular *entities.Titular) (*entities.Titular, error) {
	err := uc.db.Save(*titular)
	return titular, err
}
