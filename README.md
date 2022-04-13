# V1

### MODULES
`go mod tidy`

### COVERAGE
`go test -coverprofile=coverage.out`

`go tool cover -func=coverage.out`

`go tool cover -html=coverage.out`


### MOCKS

`go install github.com/golang/mock/mockgen@v1.6.0`

`go get github.com/golang/mock/gomock`

#### Generated mocks

`mockgen -source internal/core/ports/repositoriesPorts.go -destination mocks/mock_repositoriesPorts.go`
