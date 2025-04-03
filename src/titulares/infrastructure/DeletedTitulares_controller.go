// DeletedTitulares_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	application "github.com/Romieb26/Arquitectura--hexagonal/src/titulares/application"
	"github.com/gin-gonic/gin"
)

type DeleteTitularController struct {
	deleteUseCase *application.DeleteTitularUseCase
}

func NewDeleteTitularController(deleteUseCase *application.DeleteTitularUseCase) *DeleteTitularController {
	return &DeleteTitularController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteTitularController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el titular",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Titular eliminado exitosamente",
	})
}
