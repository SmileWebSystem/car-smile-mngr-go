package carsrv

import (
	"car-smile-mngr-go/internal/core/models"
	"car-smile-mngr-go/internal/repositories/soapRepo"
	mockports "car-smile-mngr-go/mocks"
	"car-smile-mngr-go/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mock struct {
	data   []byte
	code   int
	status string
	err    error
}

func TestCheckCar(t *testing.T) {

	tests := []struct {
		name string
		mock mock
		want *models.CarInfo
	}{
		{
			name: "analysisCarService-1 OK",
			mock: mock{
				data:   utils.LoadCommonFile("soap/xml-response-200-ok.xml"),
				code:   200,
				status: "200 OK",
				err:    nil,
			},
			want: utils.GetMockCarInfo(utils.LoadJson("resp/json-response-200-ok.json")),
		},
		{
			name: "analysisCarService getOwners-1",
			mock: mock{
				data:   utils.LoadCommonFile("soap/xml-response-200-1.xml"),
				code:   200,
				status: "200 OK",
				err:    nil,
			},
			want: utils.GetMockCarInfo(utils.LoadJson("resp/json-response-200-1.json")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockHTTPClientRepository := mockports.NewMockHTTPClientRepository(gomock.NewController(t))
			mockHTTPClientRepository.EXPECT().DoClient(gomock.Any()).Return(utils.GetMockHttpResponse(tt.mock.data, tt.mock.code, tt.mock.status, tt.mock.err))

			soapRepository := soapRepo.NewSoapRepository(mockHTTPClientRepository)
			analysisCarService := NewAnalysisCarService()

			carService := NewCarService(soapRepository, analysisCarService)

			resp, err := carService.CheckCar("xxx000")

			assert.Nil(t, err)
			assert.Equal(t, tt.want, resp)

		})
	}
}
