package main

import (
	"car-smile-mngr-go/cmd/app"
	"os"
)

func main() {
	if os.Getenv("DEPLOY_LOCAL") == "true" {
		app.LocalStart()
	} else {
		app.LambdaStar()

	}
}
