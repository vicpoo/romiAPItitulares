package application

import "github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain"

type DeleteTitularUseCase struct {
	db domain.ITitulares
}

func NewDeleteTitularUseCase(db domain.ITitulares) *DeleteTitularUseCase {
	return &DeleteTitularUseCase{
		db: db,
	}
}

func (uc *DeleteTitularUseCase) Run(id int) error {
	return uc.db.Delete(id)
}
