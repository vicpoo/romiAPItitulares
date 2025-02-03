package infrastructure

import (
	"net/http"
	"strconv"

	application "github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/application"
	"github.com/gin-gonic/gin"
)

type DeleteVehiculoController struct {
	deleteUseCase *application.DeleteVehiculoUseCase
}

func NewDeleteVehiculoController(deleteUseCase *application.DeleteVehiculoUseCase) *DeleteVehiculoController {
	return &DeleteVehiculoController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteVehiculoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(id)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar el vehículo",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Vehículo eliminado exitosamente",
	})
}
