package main

import (
	"log"
	"os"

	"github.com/zhetkerbaevan/green-api/cmd/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" //Locally we use port 9000
	}
	server := api.NewAPIServer(":" + port)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}