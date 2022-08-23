package ports

import (
	"car-smile-mngr-go/internal/core/models"
)

type CarService interface {
	CheckCar(licensePlate string) (*models.CarInfo, error)
}

type AnalysisCarService interface {
	CarInformationAnalysis(soapResponse models.SoapResponse) (*models.CarInfo, error)
}
