package infrastructure

import (
	"net/http"

	"github.com/Romieb26/Arquitectura--hexagonal/src/titulares/application"
	"github.com/gin-gonic/gin"
)

type ViewAllTitularesController struct {
	useCase *application.ViewAllTitular
}

func NewViewAllTitularesController(useCase *application.ViewAllTitular) *ViewAllTitularesController {
	return &ViewAllTitularesController{useCase: useCase}
}

func (vec *ViewAllTitularesController) Execute(c *gin.Context) {
	titulares, err := vec.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los titulares"})
		return
	}

	c.JSON(http.StatusOK, titulares)
}
