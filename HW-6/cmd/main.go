package main

import (
	"os"
	"strconv"
	"time"

	"github.com/LightAlykard/GoBackEnd-2/HW-6/api"
	"github.com/labstack/gommon/log"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	redisHost := os.Getenv("REDISHOST")
	redisPort := os.Getenv("REDISPORT")
	ttl, err := strconv.Atoi(os.Getenv("TTL"))
	if err != nil {
		log.Errorf("can't parse ttl %s", err)
		ttl = 1
	}
	server := api.New()
	server.Serve(host, port, redisHost, redisPort, time.Duration(ttl)*time.Hour)
}
