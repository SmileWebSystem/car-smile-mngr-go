package carhdl

import (
	"car-smile-mngr-go/internal/core/ports"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Handler struct {
	carService ports.CarService
}

func New(carService ports.CarService) *Handler {
	return &Handler{
		carService: carService,
	}
}

//
//
//
func (handler *Handler) CheckCarHandler(c *gin.Context) {
	log.Info("carhdl.CheckCarHandler")

	licensePlate := c.Param("licensePlate")

	//TODO validar refactor manejo de errores
	result, err := handler.carService.CheckCar(licensePlate)
	if err != nil {
		c.JSON(404, err.Error())
	}

	c.JSON(200, result)
}
