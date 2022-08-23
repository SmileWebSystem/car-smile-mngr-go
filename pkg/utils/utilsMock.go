package utils

import (
	"car-smile-mngr-go/internal/core/models"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type bodyMock struct {
	Reader io.Reader
}

func GetMockCarInfo(file []byte) *models.CarInfo {

	data := models.CarInfo{}
	_ = json.Unmarshal([]byte(file), &data)
	return &data
}

func GetMockHttpResponse(file []byte, code int, status string, err error) (*http.Response, error) {
	if err != nil {
		return nil, err
	}

	return &http.Response{
		Status:     status,
		StatusCode: code,
		Body: bodyMock{
			Reader: strings.NewReader(string(file)),
		},
	}, nil
}

func (m bodyMock) Read(b []byte) (int, error) {
	return m.Reader.Read(b)
}

func (m bodyMock) Close() error {
	return nil
}
