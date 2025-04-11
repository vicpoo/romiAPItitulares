package application

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/email/domain"
	"github.com/Romieb26/Arquitectura--hexagonal/src/email/domain/entities"
)

type SendEmailUseCase struct {
	emailService domain.IEmailService
}

func NewSendEmailUseCase(emailService domain.IEmailService) *SendEmailUseCase {
	return &SendEmailUseCase{
		emailService: emailService,
	}
}

func (uc *SendEmailUseCase) Run(email entities.Email) error {
	return uc.emailService.SendEmail(email)
}
