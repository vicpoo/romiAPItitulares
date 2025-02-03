package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/application"
	"github.com/gin-gonic/gin"
)

type ViewTitularController struct {
	useCase *application.ViewTitular
}

func NewViewTitularController(useCase *application.ViewTitular) *ViewTitularController {
	return &ViewTitularController{useCase: useCase}
}

func (vec *ViewTitularController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	titular, err := vec.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "titular no encontrado"})
		return
	}

	c.JSON(http.StatusOK, titular)
}
