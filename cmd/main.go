package main

import (
	"log"

	"github.com/josue/chalenge_golang/internal/app/api"
)

func main() {
	log.Println("Starting service")
	api.LoadApiV1()
}
