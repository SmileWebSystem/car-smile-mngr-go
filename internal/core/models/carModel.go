package models

import "car-smile-mngr-go/pkg/rules"

type CarInfo struct {
	Score        CarScore
	LicensePlate string
	History      []CarHistory
	Owners       []CarOwner
	Guia         CarGuia
	Simis        []CarSimi
	Claims       []CarClaim
}

type CarHistory struct {
	LicensePlate   string
	NameCompany    string
	CodeCompany    string
	PolicyNumber   string
	Active         string
	DateIni        string
	DateFin        string
	InsuredAmount  string
	Service        string
	Engine         string
	Chassis        string
	PolicyHolder   string
	PolicyHolderId string
	Beneficiary    string
}

type CarScore struct {
	ScoreTotal      float64
	ScoreOwners     float64
	ScoreClaims     float64
	ScoreOpenSimis  float64
	ScoreCloseSimis float64
}

type CarOwner struct {
	Name string
	Id   string
}

type CarGuia struct {
	Brand       string
	Class       string
	Type        string
	Model       string
	ActualValue float64
	Country     string
	Maker       string
}

type CarSimi struct {
	Date   string
	Number string
	Name   string
	Status string
	Value  string
	Code   string
}

type CarClaim struct {
	Company      string
	NumberClaim  string
	NumberPolicy string
	Date         string
	Shelters     []CarShelter
}

type CarShelter struct {
	Status          string
	Name            string
	ReclaimedAmount float64
	AmountPaid      string
	Date            string
}

func (carInfo *CarInfo) CalculateAllScores() {
	carInfo.CalculateScoreOwners()
	carInfo.CalculateScoreClaims()
	carInfo.CalculateScoreOpenSimi()
	carInfo.CalculateScoreCloseSimi()
	carInfo.CalculateScoreTotal()
}

func (carInfo *CarInfo) CalculateScoreOwners() {
	carInfo.Score.ScoreOwners = (rules.Owner * float64(len(carInfo.Owners)-1)) * -1
}

func (carInfo *CarInfo) CalculateScoreClaims() {
	var total float64
	for _, claim := range carInfo.Claims {
		for _, shelters := range claim.Shelters {
			total += shelters.ReclaimedAmount
		}
	}

	percentage := (total * 100) / carInfo.Guia.ActualValue
	carInfo.Score.ScoreClaims = rules.ClaimsRulesDefinitions(percentage)

}

func (carInfo *CarInfo) CalculateScoreOpenSimi() {
	var total float64
	for _, item := range carInfo.Simis {
		if item.Status == "Pendiente" {
			total++
		}
	}
	carInfo.Score.ScoreOpenSimis = (total * rules.OpenSimi) * -1
}

func (carInfo *CarInfo) CalculateScoreCloseSimi() {
	var total float64
	for _, item := range carInfo.Simis {
		if item.Status != "Pendiente" {
			total++
		}
	}
	carInfo.Score.ScoreCloseSimis = (total * rules.CloseSimi) * -1
}

func (carInfo *CarInfo) CalculateScoreTotal() {
	var total float64 = 100

	total += carInfo.Score.ScoreOwners + carInfo.Score.ScoreClaims + carInfo.Score.ScoreOpenSimis + carInfo.Score.ScoreCloseSimis
	carInfo.Score.ScoreTotal = total
}
