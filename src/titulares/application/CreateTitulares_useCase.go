// application/CreateTitulares_useCase.go
package application

import (
	"errors"
	"strings"

	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain/entities"
)

// CreateTitularUseCase maneja la creación de titulares
type CreateTitularUseCase struct {
	db domain.ITitulares
}

func NewCreateTitularUseCase(db domain.ITitulares) *CreateTitularUseCase {
	return &CreateTitularUseCase{db: db}
}

func (uc *CreateTitularUseCase) Run(titular *entities.Titular) (*entities.Titular, error) {
	if titular.Email == "" {
		return nil, errors.New("el email es requerido")
	}

	if !strings.Contains(titular.Email, "@") || !strings.Contains(titular.Email, ".") {
		return nil, errors.New("formato de email inválido")
	}

	err := uc.db.Save(*titular)
	return titular, err
}

// SendEmailUseCase maneja el envío de emails entre titulares
type SendEmailUseCase struct {
	titularRepo domain.ITitulares
}

func NewSendEmailUseCase(repo domain.ITitulares) *SendEmailUseCase {
	return &SendEmailUseCase{titularRepo: repo}
}

func (uc *SendEmailUseCase) Run(emailData entities.EmailData) (string, error) {
	fromTitular, err := uc.titularRepo.FindByID(emailData.FromID)
	if err != nil {
		return "", err
	}

	toTitular, err := uc.titularRepo.FindByID(emailData.ToID)
	if err != nil {
		return "", err
	}

	message := "De: " + fromTitular.Nombre + " " + fromTitular.Apellido +
		"\nPara: " + toTitular.Nombre + " " + toTitular.Apellido +
		"\nAsunto: " + emailData.Subject +
		"\nMensaje: " + emailData.Message

	return message, nil
}
