package restClient

import (
	"car-smile-mngr-go/pkg/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type restClientRepository struct {
}

func NewClientRepository() *restClientRepository {
	return &restClientRepository{}

}

func (restClient *restClientRepository) DoClient(req *http.Request) (*http.Response, error) {

	//use mock
	if os.Getenv("MOCK") == "true" {
		return doClientMock()
	}

	log.Info("Use Mode REAL")
	client := &http.Client{
		Timeout: 120 * time.Second,
	}
	response, err := client.Do(req)

	return response, err
}

func doClientMock() (*http.Response, error) {
	log.Info("Use Mode MOCK")

	data := utils.LoadCommonFile("soap/xml-response-200-1.xml")
	//data := utils.LoadCommonFile("soap/xml-response-200-ok.xml")
	response, err := utils.GetMockHttpResponse(data, 200, "200 OK", nil)

	return response, err

}
