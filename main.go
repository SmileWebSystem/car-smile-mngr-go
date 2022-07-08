package main

import (
	"car-smile-mngr-go/cmd/app"
	"os"
)

func main() {
	if os.Getenv("LOCAL") == "" {
		app.LambdaStar()
	} else {
		app.LocalStart()
	}
}
