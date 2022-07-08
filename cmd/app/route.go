package app

import "github.com/gin-gonic/gin"

func GetRoutes() *gin.Engine {

	router := gin.Default()
	setRoutes(router, initializeDependencies())

	return router
}

func setRoutes(router *gin.Engine, definition *Definition) {

	router.GET("/smile/v2/car/:licensePlate", definition.CarHandler.CheckCarHandler)

}
