package main

import (
	"github.com/LightAlykard/GoBackEnd-2/hw8/api"
)

func main() {
	// addr := os.Getenv("ADDRESS")
	// port := os.Getenv("PORT")
	addr := "localhost"
	port := "8080"
	api.Start(addr, port)
}
