package main

import (
	"github.com/blondefoxy/study-go/pkg/facade"
	"github.com/blondefoxy/study-go/pkg/health"
	"github.com/blondefoxy/study-go/pkg/recommends"
	"log"
)

func main() {
	healthData := health.NewHealther()
	recommend := recommends.NewRecommender()

	genom := facade.NewGenomService("1qaz", healthData, recommend)

	err := genom.Calculate()
	if err != nil {
		log.Printf("error occured: %v", err)
	}

	log.Println("OK")
}
