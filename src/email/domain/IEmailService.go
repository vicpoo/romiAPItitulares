package domain

import "github.com/Romieb26/Arquitectura--hexagonal/src/email/domain/entities"

type IEmailService interface {
	SendEmail(email entities.Email) error
}
