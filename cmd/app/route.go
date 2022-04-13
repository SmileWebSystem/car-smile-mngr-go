package app

import "github.com/gin-gonic/gin"

func SetRoutes(router *gin.Engine, definition *Definition) {

	router.GET("/smile/v2/car/:licensePlate", definition.CarHandler.CheckCarHandler)

}
