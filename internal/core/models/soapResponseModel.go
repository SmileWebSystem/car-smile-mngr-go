package models

import (
	"encoding/xml"
)

//TODO para todos los nombre a ingles

type SoapResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    BodyResponse
}

type BodyResponse struct {
	XMLName xml.Name     `xml:"Body"`
	Sisa    SisaResponse `xml:"ConsultaSisaResponse"`
}

type SisaResponse struct {
	XMLName xml.Name     `xml:"ConsultaSisaResponse"`
	Data    DataResponse `xml:"Data"`
	Header  DataHeader   `xml:"DataHeader"`
}

type DataResponse struct {
	XMLName    xml.Name            `xml:"Data"`
	Historicos []HistoricoResponse `xml:"historicoPolizasSisa"`
	Guia       GuiaResponse        `xml:"guiaSisa"`
	Simi       []SimiResponse      `xml:"SIMITSisa"`
	Wmi        Wmi                 `xml:"codigoWMISisa"`
	Claims     []Claim             `xml:"historicoSiniestrosSisa"`
}

type DataHeader struct {
	XMLName     xml.Name `xml:"DataHeader"`
	CodResponse string   `xml:"codRespuesta"`
	Error       *Errors
}
type HistoricoResponse struct {
	XMLName        xml.Name `xml:"historicoPolizasSisa"`
	LicensePlate   string   `xml:"Placa"`
	NameCompany    string   `xml:"NombreCompania"`
	CodeCompany    string   `xml:"CodigoCompania"`
	PolicyNumber   string   `xml:"NumeroPoliza"`
	Active         string   `xml:"Vigente"`
	DateIni        string   `xml:"FechaVigencia"`
	DateFin        string   `xml:"FechaFinVigencia"`
	InsuredAmount  string   `xml:"ValorAsegurado"`
	Service        string   `xml:"Servicio"`
	Engine         string   `xml:"Motor"`
	Chassis        string   `xml:"Chasis"`
	PolicyHolder   string   `xml:"NombreTomador"`
	PolicyHolderId string   `xml:"NumeroDocumentoTomador"`
	Beneficiary    string   `xml:"NombreBeneficiario"`
}

type GuiaResponse struct {
	XMLName     xml.Name `xml:"guiaSisa"`
	Brand       string   `xml:"Marca"`
	Class       string   `xml:"Clase"`
	Type        string   `xml:"Tipo"`
	Model       string   `xml:"Modelo"`
	ActualValue float64  `xml:"ValorActual"`
}

type SimiResponse struct {
	XMLName xml.Name `xml:"SIMITSisa"`
	Date    string   `xml:"FechaComparendo"`
	Number  string   `xml:"Numero"`
	Name    string   `xml:"Nombre"`
	Status  string   `xml:"Estado"`
	Value   string   `xml:"Valor"`
	Code    string   `xml:"CodigoInfraccion"`
}

type Wmi struct {
	XMLName xml.Name `xml:"codigoWMISisa"`
	Country string   `xml:"Pais"`
	Maker   string   `xml:"NombreFabricante"`
}

type Claim struct {
	XMLName      xml.Name  `xml:"historicoSiniestrosSisa"`
	Company      string    `xml:"NombreCompania"`
	NumberClaim  string    `xml:"NumeroSiniestro"`
	NumberPolicy string    `xml:"NumeroPoliza"`
	Date         string    `xml:"FechaSiniestro"`
	Shelters     []Shelter `xml:"Amparos"`
}

type Shelter struct {
	XMLName         xml.Name `xml:"Amparos"`
	Status          string   `xml:"Estado"`
	Name            string   `xml:"NombreAmparado"`
	ReclaimedAmount float64  `xml:"ValorReclamaAmparo"`
	AmountPaid      string   `xml:"ValorPagadoAmparo"`
	Date            string   `xml:"FechaAviso"`
}

type Errors struct {
	XMLName     xml.Name `xml:"errores"`
	Id          int      `xml:"id"`
	Descripcion string   `xml:"descripcion"`
	Tipo        string   `xml:"tipo"`
}

//
// GetHistoryRemoveDuplicateOwners return history removed duplicates by name
//
func (soapResponse *SoapResponse) GetHistoryRemoveDuplicateOwners() []HistoricoResponse {
	var historyUnique []HistoricoResponse
sampleLoop:
	for _, history := range soapResponse.Body.Sisa.Data.Historicos {
		for i, u := range historyUnique {
			if history.PolicyHolderId == u.PolicyHolderId {
				historyUnique[i] = history
				continue sampleLoop
			}
		}
		historyUnique = append(historyUnique, history)
	}
	return historyUnique
}
