package app

import (
	"car-smile-mngr-go/internal/core/ports"
	"car-smile-mngr-go/internal/core/services/carsrv"
	"car-smile-mngr-go/internal/handlers/carhdl"
	"car-smile-mngr-go/internal/handlers/versionhdl"
	"car-smile-mngr-go/internal/repositories/restClient"
	"car-smile-mngr-go/internal/repositories/soapRepo"
)

type Definition struct {
	CarService           ports.CarService
	ScoreService         ports.AnalysisCarService
	SoapRepository       ports.SoapRepository
	HttpClientRepository ports.HTTPClientRepository
	CarHandler           *carhdl.Handler
	VersionHandler       *versionhdl.Handler
}

func initializeDependencies() *Definition {
	definition := Definition{}

	definition.HttpClientRepository = restClient.NewClientRepository()

	definition.SoapRepository = soapRepo.NewSoapRepository(definition.HttpClientRepository)

	definition.ScoreService = carsrv.NewAnalysisCarService()
	definition.CarService = carsrv.NewCarService(definition.SoapRepository, definition.ScoreService)

	definition.CarHandler = carhdl.New(definition.CarService)

	return &definition
}
