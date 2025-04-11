// infrastructure/CreateTitulares_controller.go
package infrastructure

import (
	"net/http"

	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/application"
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain/entities"
	"github.com/gin-gonic/gin"
)

// CreateTitularController maneja las solicitudes de creación de titulares
type CreateTitularController struct {
	CreateTitularUseCase *application.CreateTitularUseCase
}

func NewCreateTitularController(useCase *application.CreateTitularUseCase) *CreateTitularController {
	return &CreateTitularController{CreateTitularUseCase: useCase}
}

func (ctrl *CreateTitularController) Run(c *gin.Context) {
	var titular entities.Titular

	if errJSON := c.ShouldBindJSON(&titular); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos del titular inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	nuevoTitular, err := entities.NewTitular(
		0,
		titular.Nombre,
		titular.Apellido,
		titular.Email,
		titular.DNIRaw,
		titular.Telefono,
		titular.Direccion,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al procesar los datos del titular",
			"error":   err.Error(),
		})
		return
	}

	titularCreado, errAdd := ctrl.CreateTitularUseCase.Run(nuevoTitular)
	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar el titular",
			"error":   errAdd.Error(),
		})
		return
	}

	response := map[string]interface{}{
		"message": "El titular ha sido agregado",
		"titular": map[string]interface{}{
			"id":        titularCreado.ID,
			"nombre":    titularCreado.Nombre,
			"apellido":  titularCreado.Apellido,
			"email":     titularCreado.Email,
			"telefono":  titularCreado.Telefono,
			"direccion": titularCreado.Direccion,
		},
	}

	c.JSON(http.StatusOK, response)
}

// SendEmailController maneja las solicitudes de envío de emails
type SendEmailController struct {
	SendEmailUseCase *application.SendEmailUseCase
}

func NewSendEmailController(useCase *application.SendEmailUseCase) *SendEmailController {
	return &SendEmailController{SendEmailUseCase: useCase}
}

func (ctrl *SendEmailController) Run(c *gin.Context) {
	var emailData entities.EmailData

	if err := c.ShouldBindJSON(&emailData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de email inválidos"})
		return
	}

	message, err := ctrl.SendEmailUseCase.Run(emailData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Email enviado exitosamente (simulado)",
		"details": message,
	})
}
