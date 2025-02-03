package infrastructure

import (
	"net/http"

	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/application"
	"github.com/gin-gonic/gin"
)

type ViewAllVehiculosController struct {
	useCase *application.ViewAllVehiculos
}

func NewViewAllVehiculosController(useCase *application.ViewAllVehiculos) *ViewAllVehiculosController {
	return &ViewAllVehiculosController{useCase: useCase}
}

func (vec *ViewAllVehiculosController) Execute(c *gin.Context) {
	vehiculos, err := vec.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los vehículos"})
		return
	}

	c.JSON(http.StatusOK, vehiculos)
}
