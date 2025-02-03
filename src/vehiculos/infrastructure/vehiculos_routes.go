package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		engine: engine,
	}
}

func (router *Router) Run() {
	createController, viewController, updateController, deleteController, viewAllController := InitVehiculoDependencies()

	titularGroup := router.engine.Group("/vehiculos")
	{
		titularGroup.POST("/", createController.Run)
		titularGroup.GET("/:id", viewController.Execute)
		titularGroup.PUT("/:id", updateController.Execute)
		titularGroup.DELETE("/:id", deleteController.Run)
		titularGroup.GET("/", viewAllController.Execute)
	}
}
