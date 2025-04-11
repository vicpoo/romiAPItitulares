// infrastructure/SendEmailController.go
package infrastructure

import (
	"net/http"

	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/application"
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain/entities"
	"github.com/gin-gonic/gin"
)

type SendEmailController struct {
	useCase *application.SendEmailUseCase
}

func NewSendEmailController(useCase *application.SendEmailUseCase) *SendEmailController {
	return &SendEmailController{useCase: useCase}
}

func (ctrl *SendEmailController) Run(c *gin.Context) {
	var emailData entities.EmailData

	if err := c.ShouldBindJSON(&emailData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de email inválidos"})
		return
	}

	message, err := ctrl.useCase.Run(emailData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email enviado exitosamente (simulado)",
		"details": message,
	})
}
