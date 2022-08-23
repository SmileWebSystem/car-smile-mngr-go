package carsrv

import (
	"car-smile-mngr-go/internal/core/models"
	"car-smile-mngr-go/internal/core/ports"
)

type analysisCarService struct {
}

func NewAnalysisCarService() ports.AnalysisCarService {
	return &analysisCarService{}
}

func (sc *analysisCarService) CarInformationAnalysis(soapResponse models.SoapResponse) (*models.CarInfo, error) {

	var carInfo models.CarInfo

	history, err := getHistory(soapResponse)
	if err != nil {
		return nil, err
	}

	owners, err := getOwners(soapResponse)
	if err != nil {
		return nil, err
	}

	guia, err := getDetailCar(soapResponse)
	if err != nil {
		return nil, err
	}

	simis, err := getSimis(soapResponse)
	if err != nil {
		return nil, err
	}

	claims, err := getClaims(soapResponse)
	if err != nil {
		return nil, err
	}

	carInfo.History = history
	carInfo.Owners = owners
	carInfo.Guia = guia
	carInfo.Simis = simis
	carInfo.Claims = claims
	carInfo.LicensePlate = soapResponse.Body.Sisa.Data.Historicos[0].LicensePlate

	return &carInfo, nil
}

func getHistory(soapResponse models.SoapResponse) ([]models.CarHistory, error) {
	var history []models.CarHistory

	for _, item := range soapResponse.Body.Sisa.Data.Historicos {
		history = append(history, models.CarHistory{
			LicensePlate:   item.LicensePlate,
			NameCompany:    item.NameCompany,
			CodeCompany:    item.CodeCompany,
			PolicyNumber:   item.PolicyNumber,
			Active:         item.Active,
			DateFin:        item.DateFin,
			DateIni:        item.DateIni,
			InsuredAmount:  item.InsuredAmount,
			Service:        item.Service,
			Engine:         item.Engine,
			Chassis:        item.Chassis,
			PolicyHolder:   item.PolicyHolder,
			PolicyHolderId: item.PolicyHolderId,
			Beneficiary:    item.Beneficiary,
		})
	}

	return history, nil
}

//
//
//
func getOwners(soapResponse models.SoapResponse) ([]models.CarOwner, error) {
	var owners []models.CarOwner

	for _, item := range soapResponse.GetHistoryRemoveDuplicateOwners() {
		owners = append(owners, models.CarOwner{Name: item.PolicyHolder, Id: item.PolicyHolderId})
	}
	return owners, nil
}

func getDetailCar(soapResponse models.SoapResponse) (models.CarGuia, error) {

	return models.CarGuia{
		Brand:       soapResponse.Body.Sisa.Data.Guia.Brand,
		Model:       soapResponse.Body.Sisa.Data.Guia.Model,
		Class:       soapResponse.Body.Sisa.Data.Guia.Class,
		Type:        soapResponse.Body.Sisa.Data.Guia.Type,
		ActualValue: soapResponse.Body.Sisa.Data.Guia.ActualValue,
		Country:     soapResponse.Body.Sisa.Data.Wmi.Country,
		Maker:       soapResponse.Body.Sisa.Data.Wmi.Maker,
	}, nil

}

func getSimis(soapResponse models.SoapResponse) ([]models.CarSimi, error) {
	var simis []models.CarSimi
	for _, item := range soapResponse.Body.Sisa.Data.Simi {
		simis = append(simis, models.CarSimi{
			Date:   item.Date,
			Name:   item.Name,
			Number: item.Number,
			Status: item.Status,
			Value:  item.Value,
			Code:   item.Code,
		})
	}
	return simis, nil

}

func getClaims(soapResponse models.SoapResponse) ([]models.CarClaim, error) {
	var claims []models.CarClaim

	for _, item := range soapResponse.Body.Sisa.Data.Claims {

		var shelters []models.CarShelter
		for _, item := range item.Shelters {
			shelters = append(shelters, models.CarShelter{
				Status:          item.Status,
				Name:            item.Name,
				ReclaimedAmount: item.ReclaimedAmount,
				AmountPaid:      item.AmountPaid,
				Date:            item.Date,
			})
		}

		claims = append(claims, models.CarClaim{
			Company:      item.Company,
			NumberClaim:  item.NumberClaim,
			NumberPolicy: item.NumberPolicy,
			Date:         item.Date,
			Shelters:     shelters,
		})
	}

	return claims, nil
}
