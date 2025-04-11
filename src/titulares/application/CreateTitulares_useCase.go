// application/CreateTitulares_useCase.go
package application

import (
	"errors"
	"strings"

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
	// Validar que el email no esté vacío
	if titular.Email == "" {
		return nil, errors.New("el email es requerido")
	}

	// Validar formato de email (puedes usar una librería más robusta para esto)
	if !strings.Contains(titular.Email, "@") || !strings.Contains(titular.Email, ".") {
		return nil, errors.New("formato de email inválido")
	}

	err := uc.db.Save(*titular)
	if err != nil {
		return nil, err
	}

	return titular, nil
}
