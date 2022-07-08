package app

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func LocalStart() {
	router := GetRoutes()
	router.Run(":8083")
}

func LambdaStar() {
	lambda.Start(Handler)

}
