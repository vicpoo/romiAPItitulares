// application/SendEmailUseCase.go
package application

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain/entities"
)

type SendEmailUseCase struct {
	titularRepo domain.ITitulares
}

func NewSendEmailUseCase(repo domain.ITitulares) *SendEmailUseCase {
	return &SendEmailUseCase{
		titularRepo: repo,
	}
}

func (uc *SendEmailUseCase) Run(emailData entities.EmailData) (string, error) {
	// Obtener información del remitente
	fromTitular, err := uc.titularRepo.FindByID(emailData.FromID)
	if err != nil {
		return "", err
	}

	// Obtener información del destinatario
	toTitular, err := uc.titularRepo.FindByID(emailData.ToID)
	if err != nil {
		return "", err
	}

	// Simular envío de email (en producción aquí iría el código real)
	message := "De: " + fromTitular.Nombre + " " + fromTitular.Apellido +
		"\nPara: " + toTitular.Nombre + " " + toTitular.Apellido +
		"\nAsunto: " + emailData.Subject +
		"\nMensaje: " + emailData.Message

	return message, nil
}
