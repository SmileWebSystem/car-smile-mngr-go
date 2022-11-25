package soapRepo

import (
	mockports "car-smile-mngr-go/mocks"
	"car-smile-mngr-go/pkg/errors"
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

func TestGetSoapInfoCarOk(t *testing.T) {

	type want struct {
		totalHist int
	}

	tests := []struct {
		name string
		mock mock
		want want
	}{
		{
			name: "200 OK",
			mock: mock{
				data:   utils.LoadCommonFile("soap/xml-response-200-ok.xml"),
				code:   200,
				status: "200 OK",
				err:    nil,
			},
			want: want{
				totalHist: 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//
			// Set-Up
			//
			mockHTTPClientRepository := mockports.NewMockHTTPClientRepository(gomock.NewController(t))
			mockHTTPClientRepository.EXPECT().DoClient(gomock.Any()).Return(utils.GetMockHttpResponse(tt.mock.data, tt.mock.code, tt.mock.status, tt.mock.err))

			//
			// Execute
			//
			soapRepository := NewSoapRepository(mockHTTPClientRepository)
			res, err := soapRepository.GetSoapInfoCar("222zzz")

			assert.Nil(t, err)
			assert.Equal(t, "0", res.Body.Sisa.Header.CodResponse)
			assert.Equal(t, tt.want.totalHist, len(res.Body.Sisa.Data.Historicos))

		})
	}

}

func TestGetSoapInfoCarErrors(t *testing.T) {

	type want struct {
		err error
	}

	//
	// Test Cases
	//
	tests := []struct {
		name string
		mock mock
		want want
	}{
		{
			name: "Generic error",
			mock: mock{
				data:   nil,
				code:   500,
				status: "500",
				err:    errors.NewError("500", "Generic error"),
			},
			want: want{
				err: errors.NewError("404", "Error call service"),
			},
		},
		{
			name: "404 No existe la placa",
			mock: mock{
				data:   utils.LoadCommonFile("soap/xml-response-200-noExiste.xml"),
				code:   200,
				status: "200 OK",
				err:    nil,
			},
			want: want{
				err: errors.NewError("404", "No existe informacion para la placa consultada"),
			},
		},
		{
			name: "404 Error call service",
			mock: mock{
				data:   utils.LoadCommonFile("soap/xml-response-404.xml"),
				code:   404,
				status: "404",
				err:    nil,
			},
			want: want{
				err: errors.NewError("404", "Error call service"),
			},
		},
		{
			name: "Error Unmarshal response",
			mock: mock{
				data:   utils.LoadCommonFile("soap/xml-response-404.xml"),
				code:   200,
				status: "200 OK",
				err:    nil,
			},
			want: want{
				err: errors.NewError("404", "Error Unmarshal response"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//
			// Set-Up
			//
			mockHTTPClientRepository := mockports.NewMockHTTPClientRepository(gomock.NewController(t))
			mockHTTPClientRepository.EXPECT().DoClient(gomock.Any()).Return(utils.GetMockHttpResponse(tt.mock.data, tt.mock.code, tt.mock.status, tt.mock.err))

			//
			// Execute
			//
			soapRepository := NewSoapRepository(mockHTTPClientRepository)
			res, err := soapRepository.GetSoapInfoCar("222zzz")

			assert.Nil(t, res)
			assert.Equal(t, tt.want.err, err)

		})
	}

}
