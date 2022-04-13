package ports

import (
	"car-smile-mngr-go/internal/core/models"
	"net/http"
)

type SoapRepository interface {
	GetSoapInfoCar(licensePlate string) (*models.SoapResponse, error)
}

//HTTPClientRepository HTTPClient interface
type HTTPClientRepository interface {
	DoClient(req *http.Request) (*http.Response, error)
}
