package main

import (
	"os"
	"strconv"
	"time"

	"github.com/LightAlykard/GoBackEnd-2/HW-6/api"
	"github.com/labstack/gommon/log"
)

func main() {
	host := "localhost"
	port := "8080"
	redisHost := "localhost"
	redisPort := "6379"
	ttl := 1
	server := api.New()
	server.Serve(host, port, redisHost, redisPort, time.Duration(ttl)*time.Hour)
}
