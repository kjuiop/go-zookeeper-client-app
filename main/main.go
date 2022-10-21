package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-zookeeper-client-app/api"
	"log"
	"os"
	"runtime"
)

func main() {
	fmt.Println("hello go-zookeeper-client-app")

	a, err := api.NewHandler()
	if err != nil {
		log.Println("[main] failed NewHandler :", err)
		os.Exit(1)
	}
	defer a.Close()

	runtime.GOMAXPROCS(runtime.NumCPU())

	gMux := gin.Default()
	gMux.GET("/api/health-check", a.HealthCheck)

	apiPort := a.GetApiPort()
	gMux.Run(":" + apiPort)
}
