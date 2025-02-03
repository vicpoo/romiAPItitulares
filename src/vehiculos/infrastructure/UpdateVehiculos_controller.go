package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/application"
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain/entities"
	"github.com/gin-gonic/gin"
)

type UpdateVehiculoController struct {
	useCase *application.UpdateVehiculo
}

func NewUpdateVehiculoController(useCase *application.UpdateVehiculo) *UpdateVehiculoController {
	return &UpdateVehiculoController{useCase: useCase}
}

func (uec *UpdateVehiculoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var vehiculo entities.Vehiculo
	if err := c.ShouldBindJSON(&vehiculo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uec.useCase.Execute(id, vehiculo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el vehículo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vehículo actualizado exitosamente"})
}
