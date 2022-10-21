package main

import (
	"fmt"
	"go-zookeeper-client-app/api"
	"log"
	"os"
)

func main() {
	fmt.Println("hello go-zookeeper-client-app")

	a, err := api.NewHandler()
	if err != nil {
		log.Println("[main] failed NewHandler :", err)
		os.Exit(1)
	}
	defer a.Close()

}
