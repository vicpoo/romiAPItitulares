// infrastructure/CreateTitulares_controller.go
package infrastructure

import (
	"net/http"

	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/application"
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateTitularController struct {
	CreateTitularUseCase *application.CreateTitularUseCase
}

func NewCreateTitularController(createTitularUseCase *application.CreateTitularUseCase) *CreateTitularController {
	return &CreateTitularController{
		CreateTitularUseCase: createTitularUseCase,
	}
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

	// Crear nuevo titular con DNI encriptado
	nuevoTitular, err := entities.NewTitular(
		0, // ID se generará en la DB
		titular.Nombre,
		titular.Apellido,
		titular.Email,
		titular.DNIRaw, // Usamos el DNI sin encriptar
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

	// No devolvemos el DNI encriptado en la respuesta
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
