// EmailService.go
package infrastructure

import (
	"log"

	"github.com/Romieb26/Arquitectura--hexagonal/src/email/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/email/domain/entities"
)

type MockEmailService struct{}

func NewMockEmailService() domain.IEmailService {
	return &MockEmailService{}
}

func (s *MockEmailService) SendEmail(email entities.Email) error {
	log.Printf("Simulando envío de email a: %s\nAsunto: %s\nCuerpo: %s\n",
		email.To, email.Subject, email.Body)

	// Aquí podrías agregar lógica para "fallar" aleatoriamente si quieres probar errores
	return nil
}
