package main

import (
	"fmt"

	"github.com/Romieb26/Arquitectura--hexagonal/src/core"
	titularesInfra "github.com/Romieb26/Arquitectura--hexagonal/src/titulares/infrastructure"
	vehiculosInfra "github.com/Romieb26/Arquitectura--hexagonal/src/vehiculos/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	core.InitDB()

	r := gin.Default()

	titularRouter := titularesInfra.NewRouter(r)
	titularRouter.Run()

	vehiculosRouter := vehiculosInfra.NewRouter(r)
	vehiculosRouter.Run()

	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
