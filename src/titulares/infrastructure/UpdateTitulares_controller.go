// UpdateTitulares_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/application"
	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateTitularController struct {
	useCase *application.UpdateTitular
}

func NewUpdateTitularController(useCase *application.UpdateTitular) *UpdateTitularController {
	return &UpdateTitularController{useCase: useCase}
}

func (uec *UpdateTitularController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var titular entities.Titular
	if err := c.ShouldBindJSON(&titular); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uec.useCase.Execute(id, titular)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el titular"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "titular actualizado exitosamente"})
}
