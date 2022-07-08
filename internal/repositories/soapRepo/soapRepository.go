package soapRepo

import (
	"bytes"
	"car-smile-mngr-go/internal/core/models"
	"car-smile-mngr-go/internal/core/ports"
	"car-smile-mngr-go/pkg/errors"
	"car-smile-mngr-go/pkg/utils"
	"encoding/xml"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"text/template"
)

const SoapEndpoint = "aHR0cHM6Ly9hbWJpZW50ZXBydWViYXMuc2VndXJvc2JvbGl2YXIuY29tL1NpbW9uV1MvRmFzZWNvbGRhU2VydmljZQ=="

type soapRepository struct {
	httpClientRepository ports.HTTPClientRepository
}

func NewSoapRepository(httpClientRepository ports.HTTPClientRepository) *soapRepository {
	return &soapRepository{
		httpClientRepository: httpClientRepository,
	}

}

//
//
//
func (repo *soapRepository) GetSoapInfoCar(licensePlate string) (*models.SoapResponse, error) {
	log.Info("soapRepo.GetSoapInfoCar: ", licensePlate)

	req := populateRequest(licensePlate)
	httpReq, err := generateSOAPRequest(req)
	if err != nil {
		log.Error("Some problem occurred in request generation:", err)
		return nil, err
	}

	response, err := soapCall(repo, httpReq)
	//log.Info("SOAP Response: ", response)
	if err != nil {
		log.Error("soapRepository.GetSoapInfoCar: ", err)
		return nil, err
	}

	if response.Body.Sisa.Header.CodResponse != "0" {
		log.Error("No existe informacion para la placa consultada", response)
		return nil, errors.NewError("404", "No existe informacion para la placa consultada")
	}

	return response, nil
}

//
//
//
func generateSOAPRequest(req *RequestParam) (*http.Request, error) {
	template, err := template.New("InputRequest").Parse(templateRequest)
	if err != nil {
		log.Error("Error while marshling object: ", err.Error())
		return nil, errors.NewError("404", "Error while marshling object")
	}

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	err = template.Execute(doc, req)
	if err != nil {
		log.Error("template.Execute error: ", err.Error())
		return nil, errors.NewError("404", "Error generate request")
	}

	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		log.Error("encoder.Encode error: ", err.Error())
		return nil, errors.NewError("404", "Error generate request encode")
	}

	r, err := http.NewRequest(http.MethodPost, utils.DecodeString(SoapEndpoint), bytes.NewBuffer([]byte(doc.String())))
	if err != nil {
		log.Error("Error making a request: ", err.Error())
		return nil, errors.NewError("404", "Error making a request")
	}

	return r, nil
}

//
//
//
func soapCall(repo *soapRepository, req *http.Request) (*models.SoapResponse, error) {

	resp, err := repo.httpClientRepository.DoClient(req)

	if err != nil {
		log.Error("Error call service:", err.Error())
		return nil, errors.NewError("404", "Error call service")
	}
	log.Info("Response SOAP: ", resp.Status)

	if resp.StatusCode != 200 {
		log.Error("Error call service:", resp.StatusCode)
		return nil, errors.NewError(resp.Status, "Error call service")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error read response:", err.Error())
		return nil, errors.NewError("404", "Error read response")
	}
	defer resp.Body.Close()

	r := &models.SoapResponse{}
	err = xml.Unmarshal(body, &r)
	if err != nil {
		log.Error("Error Unmarshal response:", err.Error())
		return nil, errors.NewError("404", "Error Unmarshal response")
	}

	log.Info("Response Code: ", r.Body.Sisa.Header.CodResponse)

	return r, nil
}
