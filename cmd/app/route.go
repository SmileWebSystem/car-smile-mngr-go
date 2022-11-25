package app

import "github.com/gin-gonic/gin"

const baseUrl = "/smile/v2/car"

func GetRoutes() *gin.Engine {

	router := gin.Default()
	setRoutes(router, initializeDependencies())

	return router
}

func setRoutes(router *gin.Engine, definition *Definition) {

	router.GET(baseUrl+"/:licensePlate", definition.CarHandler.CheckCarHandler)
	router.GET(baseUrl+"/version", definition.VersionHandler.VersionHandler)

}
