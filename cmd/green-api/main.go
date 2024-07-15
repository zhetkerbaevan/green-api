package main

import (
	"log"

	"github.com/zhetkerbaevan/green-api/cmd/api"
)

func main() {
	server := api.NewAPIServer(":9000")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}