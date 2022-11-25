package soapRepo

type RequestParam struct {
	LicensePlate string
}

func populateRequest(licensePlate string) *RequestParam {
	req := RequestParam{}
	req.LicensePlate = licensePlate
	return &req
}
