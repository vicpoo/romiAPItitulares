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

	titularCreado, errAdd := ctrl.CreateTitularUseCase.Run(&titular)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar el titular",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "El titular ha sido agregado",
		"empleado": titularCreado,
	})
}
