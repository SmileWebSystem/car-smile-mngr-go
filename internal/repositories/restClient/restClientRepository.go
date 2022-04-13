package restClient

import (
	"car-smile-mngr-go/pkg/utils"
	"net/http"
)

type restClientRepository struct {
}

func NewClientRepository() *restClientRepository {
	return &restClientRepository{}

}

func (restClient *restClientRepository) DoClient(req *http.Request) (*http.Response, error) {
	//client := &http.Client{
	//	Timeout: 20 * time.Second,
	//}
	//response, err := client.Do(req)

	//TODO refactorizar mock en variable de entorno
	data := utils.LoadCommonFile("soap/xml-response-200-1.xml")
	//data := utils.LoadCommonFile("soap/xml-response-200-ok.xml")
	response, err := utils.GetMockHttpResponse(data, 200, "200 OK", nil)

	return response, err
}
