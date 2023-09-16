package versionhdl

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (handler *Handler) VersionHandler(c *gin.Context) {
	c.JSON(200, "1.0.1")
}
