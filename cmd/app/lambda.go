package app

type MyInput struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

type MyResponse struct {
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

func Handler(event MyInput) (MyResponse, error) {
	return MyResponse{StatusCode: "201", StatusMessage: event.Action}, nil
}
