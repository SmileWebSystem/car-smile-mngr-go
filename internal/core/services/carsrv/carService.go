package carsrv

import (
	"car-smile-mngr-go/internal/core/models"
	"car-smile-mngr-go/internal/core/ports"
	log "github.com/sirupsen/logrus"
)

type carService struct {
	carSoapRepository ports.SoapRepository
	analysisCar       ports.AnalysisCarService
}

func NewCarService(soapRepository ports.SoapRepository, analysisCarService ports.AnalysisCarService) ports.CarService {
	return &carService{
		carSoapRepository: soapRepository,
		analysisCar:       analysisCarService,
	}
}

//
// CheckCar Service
//
func (carSrv *carService) CheckCar(licensePlate string) (*models.CarInfo, error) {

	//TODO revisar refactorizacion para manejo de errores

	//get soap info car
	soapResponse, err := carSrv.carSoapRepository.GetSoapInfoCar(licensePlate)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// mapper info in json struct
	carInfo, err := carSrv.analysisCar.CarInformationAnalysis(*soapResponse)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	//calculate all scores about car info
	carInfo.CalculateAllScores()

	return carInfo, nil
}
