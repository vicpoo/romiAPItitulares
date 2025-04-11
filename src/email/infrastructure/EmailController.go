// EmailController.go
package infrastructure

import (
	"net/http"

	"github.com/Romieb26/Arquitectura--hexagonal/src/email/application"
	"github.com/Romieb26/Arquitectura--hexagonal/src/email/domain/entities"
	"github.com/gin-gonic/gin"
)

type EmailController struct {
	sendEmailUseCase *application.SendEmailUseCase
}

func NewEmailController(sendEmailUseCase *application.SendEmailUseCase) *EmailController {
	return &EmailController{
		sendEmailUseCase: sendEmailUseCase,
	}
}

func (ctrl *EmailController) SendEmail(c *gin.Context) {
	var email entities.Email

	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email data",
		})
		return
	}

	if err := ctrl.sendEmailUseCase.Run(email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to send email",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email sent successfully (simulated)",
		"email":   email,
	})
}
