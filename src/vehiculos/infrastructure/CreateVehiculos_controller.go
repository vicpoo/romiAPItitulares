package infrastructure

import (
	"net/http"

	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/application"
	"github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateVehiculoController struct {
	CreateVehiculoUseCase *application.CreateVehiculoUseCase
}

func NewCreateVehiculoController(createVehiculoUseCase *application.CreateVehiculoUseCase) *CreateVehiculoController {
	return &CreateVehiculoController{
		CreateVehiculoUseCase: createVehiculoUseCase,
	}
}

func (ctrl *CreateVehiculoController) Run(c *gin.Context) {
	var vehiculo entities.Vehiculo

	if errJSON := c.ShouldBindJSON(&vehiculo); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos del vehículo inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	vehiculoCreado, errAdd := ctrl.CreateVehiculoUseCase.Run(&vehiculo)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar el vehículo",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "El vehículo ha sido agregado",
		"vehiculo": vehiculoCreado,
	})
}
