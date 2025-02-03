package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/application"
	"github.com/gin-gonic/gin"
)

type ViewVehiculoController struct {
	useCase *application.ViewVehiculo
}

func NewViewVehiculoController(useCase *application.ViewVehiculo) *ViewVehiculoController {
	return &ViewVehiculoController{useCase: useCase}
}

func (vec *ViewVehiculoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	vehiculo, err := vec.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehículo no encontrado"})
		return
	}

	c.JSON(http.StatusOK, vehiculo)
}
