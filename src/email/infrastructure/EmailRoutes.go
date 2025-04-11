// EmailRoutes.go
package infrastructure

import (
	"github.com/Romieb26/Arquitectura--hexagonal/src/email/application"
	"github.com/gin-gonic/gin"
)

type EmailRouter struct {
	engine *gin.Engine
}

func NewEmailRouter(engine *gin.Engine) *EmailRouter {
	return &EmailRouter{
		engine: engine,
	}
}

func (router *EmailRouter) Run() {
	// Inicializar dependencias
	emailService := NewMockEmailService()
	sendEmailUseCase := application.NewSendEmailUseCase(emailService)
	emailController := NewEmailController(sendEmailUseCase)

	emailGroup := router.engine.Group("/email")
	{
		emailGroup.POST("/send", emailController.SendEmail)
	}
}
