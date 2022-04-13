package app

import "github.com/gin-gonic/gin"

func Start() {

	router := gin.New()
	SetRoutes(router, initializeDependencies())
	router.Run(":8083")
}
