package app

func Start() {
	router := GetRoutes()
	router.Run(":8083")
}
