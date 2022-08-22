package app

import "github.com/apex/gateway"

func LocalStart() {
	router := GetRoutes()
	router.Run(":8083")
}

func LambdaStar() {
	gateway.ListenAndServe(":8080", GetRoutes())

}
